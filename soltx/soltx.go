package soltx

import (
	"context"
	"encoding/json"
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/go-hclog"
	"github.com/shopspring/decimal"
	"github.com/sol-tx/blocksubscribe"
	"github.com/sol-tx/config"
	"github.com/sol-tx/db"
	"github.com/sol-tx/log"
	raydium_amm "github.com/sol-tx/raydiumamm/generated"
	"github.com/sol-tx/types"
	"os"
)

type Handler struct {
	ctx            context.Context
	log            hclog.Logger
	blockSubscribe *blocksubscribe.BlockSubscribe
	updatedBlocks  chan *rpc.GetParsedBlockResult
	dao            *db.Dao
}

func newBlockSubscribe(ctx context.Context, cfg config.BlockSubscribe, cb blocksubscribe.Callback) *blocksubscribe.BlockSubscribe {
	rpcUrls := make([]string, 0)
	wsUrls := make([]string, 0)
	for _, node := range cfg.Nodes {
		if node.Enable {
			rpcUrls = append(rpcUrls, node.Rpc)
			wsUrls = append(wsUrls, node.Ws)
		}
	}
	return blocksubscribe.New(ctx, rpcUrls, wsUrls, cb)
}

func newDao(ctx context.Context, cfg config.Dao) *db.Dao {
	c := &db.Config{
		User:     cfg.User,
		Password: cfg.Password,
		Url:      cfg.Url,
		Scheme:   cfg.Scheme,
		Port:     cfg.Port,
		Debug:    cfg.Debug,
	}
	return db.New(c)
}

func New(ctx context.Context, dir string) *Handler {
	configFilePath := fmt.Sprintf("%s/config.json", dir)
	configJson, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	var cfg config.Config
	err = json.Unmarshal(configJson, &cfg)
	if err != nil {
		panic(err)
	}
	//
	t := &Handler{
		ctx:           ctx,
		log:           log.NewLog("sol-tx"),
		updatedBlocks: make(chan *rpc.GetParsedBlockResult, 1024),
	}
	//
	t.blockSubscribe = newBlockSubscribe(ctx, cfg.BlockSubscribe, t)
	t.dao = newDao(ctx, cfg.Dao)
	return t
}

func (h *Handler) Start() error {
	h.blockSubscribe.Start()
	go h.process()
	return nil
}

func (h *Handler) Service() {
	h.Start()
	<-h.ctx.Done()
	h.Stop()
}

func (h *Handler) Stop() error {
	return nil
}

func (h *Handler) OnBlock(block *rpc.GetParsedBlockResult) error {
	h.updatedBlocks <- block
	return nil
}

func (h *Handler) process() {
	defer func() {
		h.log.Info("blockHandler process exit")
	}()
	for {
		select {
		case block := <-h.updatedBlocks:
			h.processBlock(block)
		case <-h.ctx.Done():
			return
		}
	}
}

func subInstructions(ins []*rpc.ParsedInstruction, inTree *types.Transaction) {
	for i, in := range ins {
		in.StackHeight = 1
		current := &types.InstructionNode{
			Seq:         i,
			Instruction: in,
			Children:    nil,
		}
		inTree.Instructions = append(inTree.Instructions, current)
	}
}

func subInnerInstructions(ins []*rpc.ParsedInstruction, depth int, inTree *types.InstructionNode) {
	var current *types.InstructionNode
	for i, in := range ins {
		if in.StackHeight == depth+1 {
			subInnerInstructions(ins[i:], in.StackHeight, current)
		} else if in.StackHeight == depth-1 {
			return
		} else {
			current = &types.InstructionNode{
				Seq:         i,
				Instruction: in,
				Children:    nil,
			}
			inTree.Children = append(inTree.Children, current)
		}
	}
}

func (h *Handler) parseBlock(block *rpc.GetParsedBlockResult) []*types.Transaction {
	h.log.Info("blockHandler", "hash", block.Blockhash, "time", block.BlockTime.Time().UTC().Format("2006-01-02 15:04:05"))
	// must with meta - json parsed
	txTrees := make([]*types.Transaction, 0)
	for _, transaction := range block.Transactions {
		meta := transaction.Meta
		if meta == nil {
			h.log.Error("transaction meta is missing")
			continue
		}
		if meta.Err != nil {
			h.log.Warn("transaction failed, ignore this one")
			continue
		}
		parsedTransaction := transaction.Transaction
		message := parsedTransaction.Message
		instructions := message.Instructions
		if len(instructions) == 0 {
			h.log.Warn("no instruction")
			continue
		}
		if instructions[0].ProgramId == types.Vote {
			continue
		}
		//
		accounts := make([]solana.PublicKey, 0)
		for _, item := range message.AccountKeys {
			accounts = append(accounts, item.PublicKey)
		}
		inTree := &types.Transaction{
			Hash:         transaction.Transaction.Signatures[0],
			Instructions: nil,
		}
		subInstructions(instructions, inTree)
		innerInstructions := meta.InnerInstructions
		for _, innerInstruction := range innerInstructions {
			pinn := inTree.Instructions[innerInstruction.Index]
			subInnerInstructions(innerInstruction.Instructions, 2, pinn)
		}
		txTrees = append(txTrees, inTree)
	}
	return txTrees
}

func (h *Handler) processBlock(block *rpc.GetParsedBlockResult) {
	tree := h.parseBlock(block)
	b := &db.Block{
		Height: *block.BlockHeight,
		Time:   uint64(*block.BlockTime),
		Hash:   block.Blockhash.String(),
		Slot:   0,
	}
	// save hash2Transactions
	hash2Transactions := make(map[string]*db.Transaction, 0)
	id2Trades := make(map[string]*db.Trade, 0)
	id2Transfers := make(map[string]*db.Transfer, 0)
	//tokens := make(map[string]*db.Token, 0)
	hash2Pools := make(map[string]*db.Pool, 0)
	for _, item := range tree {
		hash2Transactions[item.Hash.String()] = &db.Transaction{
			Hash:        item.Hash.String(),
			BlockHeight: *block.BlockHeight,
			Time:        uint64(*block.BlockTime),
		}
		for _, item1 := range item.Instructions {
			h.processInstruction(item1, id2Trades, id2Transfers, hash2Pools)
		}
	}
	// save all
	transactions := make([]*db.Transaction, 0)
	for _, v := range hash2Transactions {
		transactions = append(transactions, v)
	}
	trades := make([]*db.Trade, 0)
	for _, v := range id2Trades {
		trades = append(trades, v)
	}
	transfers := make([]*db.Transfer, 0)
	for _, v := range id2Transfers {
		transfers = append(transfers, v)
	}
	pools := make([]*db.Pool, 0)
	for _, v := range hash2Pools {
		pools = append(pools, v)
	}
	h.dao.SaveBlock(b)
	h.dao.SaveTransaction(transactions)
	h.dao.SaveTrade(trades)
	h.dao.SaveTransfer(transfers)
	h.dao.SavePool(pools)
}

func (h *Handler) processInstruction(in *types.InstructionNode, trades map[string]*db.Trade, transfers map[string]*db.Transfer, pools map[string]*db.Pool) () {
	switch in.Instruction.ProgramId {
	case types.System:
		h.processSystemTransfer(in)
	case types.Token:
		h.processTokenTransfer(in)
	case types.RaydiumAMM:
		h.processRaydiumAmmTrade(in)
	case types.RaydiumClmm:
		h.processRaydiumClmmTrade(in)
	case types.WhirlPool:
		h.processWhirlPoolTrade(in)
	default:
		for _, item := range in.Children {
			h.processInstruction(item, trades, transfers, pools)
		}
	}
}

func (h *Handler) processSystemTransfer(in *types.InstructionNode) *db.Transfer {
	type instruction struct {
		Info struct {
			Destination solana.PublicKey `json:"destination"`
			Lamports    uint64           `json:"lamports"`
			Source      solana.PublicKey `json:"source"`
		} `json:"info"`
		T string `json:"type"`
	}
	ins := in.Instruction
	var s instruction
	k, _ := ins.Parsed.MarshalJSON()
	json.Unmarshal(k, &s)
	transfer := &db.Transfer{
		Mint:   "",
		Amount: s.Info.Lamports,
		From:   s.Info.Source.String(),
		To:     s.Info.Destination.String(),
	}
	return transfer
}

func (h *Handler) processTokenTransfer(in *types.InstructionNode) *db.Transfer {
	type instruction struct {
		Info struct {
			Destination solana.PublicKey `json:"destination"`
			Lamports    uint64           `json:"lamports"`
			Source      solana.PublicKey `json:"source"`
			Authority   solana.PublicKey `json:"authority"`
			Mint        solana.PublicKey `json:"mint"`
			TokenAmount struct {
				Amount   decimal.Decimal
				Decimals uint64
			} `json:"tokenAmount"`
		} `json:"info"`
		T string `json:"type"`
	}
	ins := in.Instruction
	var s instruction
	k, _ := ins.Parsed.MarshalJSON()
	json.Unmarshal(k, &s)
	transfer := &db.Transfer{
		Mint:   "",
		Amount: s.Info.Lamports,
		From:   s.Info.Source.String(),
		To:     s.Info.Destination.String(),
	}
	return transfer
}

func (h *Handler) processRaydiumAmmTrade(in *types.InstructionNode) *db.Trade {
	inst := new(raydium_amm.Instruction)
	data := in.Instruction.Data
	err := ag_binary.NewBorshDecoder(data).Decode(inst)
	if err != nil {
		return nil
	}
	if inst.TypeID.Uint8() != raydium_amm.Instruction_Deposit {
		accounts := make([]*solana.AccountMeta, 0)
		for _, item := range in.Instruction.Accounts {
			accounts = append(accounts, &solana.AccountMeta{
				PublicKey:  item,
				IsWritable: false,
				IsSigner:   false,
			})
		}
		inst1 := inst.Impl.(*raydium_amm.Deposit)
		inst1.SetAccounts(accounts)
		//
		type instruction struct {
			Info struct {
				Destination solana.PublicKey `json:"destination"`
				Lamports    uint64           `json:"lamports"`
				Source      solana.PublicKey `json:"source"`
				Authority   solana.PublicKey `json:"authority"`
				Mint        solana.PublicKey `json:"mint"`
				TokenAmount struct {
					Amount   decimal.Decimal
					Decimals uint64
				} `json:"tokenAmount"`
			} `json:"info"`
			T string `json:"type"`
		}
		tokenAAmount := uint64(0)
		{
			ins := in.Instruction
			var s instruction
			k, _ := ins.Parsed.MarshalJSON()
			json.Unmarshal(k, &s)
			tokenAAmount = s.Info.Lamports
		}
		tokenBAmount := uint64(0)
		{
			ins := in.Instruction
			var s instruction
			k, _ := ins.Parsed.MarshalJSON()
			json.Unmarshal(k, &s)
			tokenBAmount = s.Info.Lamports
		}
		trade := &db.Trade{
			Pool:         inst1.GetAmmAccount().PublicKey.String(),
			Type:         "deposit",
			TokenAAmount: decimal.NewFromInt(int64(tokenAAmount)),
			TokenBAmount: decimal.NewFromInt(int64(tokenBAmount)),
			User:         inst1.GetAmmAuthorityAccount().PublicKey.String(),
		}
		return trade
	}
	return nil
}

func (h *Handler) processRaydiumClmmTrade(in *types.InstructionNode) *db.Trade {
	return nil
}

func (h *Handler) processWhirlPoolTrade(in *types.InstructionNode) *db.Trade {
	return nil
}
