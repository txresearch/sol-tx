package soltx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/go-hclog"
	"github.com/sol-tx/blocksubscribe"
	"github.com/sol-tx/config"
	"github.com/sol-tx/db"
	"github.com/sol-tx/log"
	"github.com/sol-tx/types"
	"os"
)

type Handler struct {
	ctx             context.Context
	log             hclog.Logger
	blockSubscriber *blocksubscribe.BlockSubscribe
	updatedBlocks   chan *rpc.GetParsedBlockResult
	dao             *db.Dao
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
	t.blockSubscriber = newBlockSubscribe(ctx, cfg.BlockSubscribe, t)
	t.dao = newDao(ctx, cfg.Dao)
	return t
}

func (t *Handler) Start() error {
	t.blockSubscriber.Start()
	go t.process()
	return nil
}

func (t *Handler) Service() {
	t.Start()
	<-t.ctx.Done()
	t.Stop()
}

func (t *Handler) Stop() error {
	return nil
}

func (t *Handler) OnBlock(block *rpc.GetParsedBlockResult) error {
	t.updatedBlocks <- block
	return nil
}

func (t *Handler) process() {
	defer func() {
		t.log.Info("blockHandler process exit")
	}()
	for {
		select {
		case block := <-t.updatedBlocks:
			t.processBlock(block)
		case <-t.ctx.Done():
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

func (t *Handler) parseBlock(block *rpc.GetParsedBlockResult) []*types.Transaction {
	t.log.Info("blockHandler", "hash", block.Blockhash, "time", block.BlockTime.Time().UTC().Format("2006-01-02 15:04:05"))
	// must with meta - json parsed
	txTrees := make([]*types.Transaction, 0)
	for _, transaction := range block.Transactions {
		meta := transaction.Meta
		if meta == nil {
			t.log.Error("transaction meta is missing")
			continue
		}
		if meta.Err != nil {
			t.log.Warn("transaction failed, ignore this one")
			continue
		}
		parsedTransaction := transaction.Transaction
		message := parsedTransaction.Message
		instructions := message.Instructions
		if len(instructions) == 0 {
			t.log.Warn("no instruction")
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

func (t *Handler) processBlock(block *rpc.GetParsedBlockResult) {
	tree := t.parseBlock(block)
	/*
		b := &db.Block{
			Height: *block.BlockHeight,
			Time:   uint64(*block.BlockTime),
			Hash:   block.Blockhash.String(),
			Slot:   0,
		}

	*/
	// save transactions
	transactions := make(map[string]*db.Transaction, 0)
	trades := make(map[string]*db.Trade, 0)
	transfers := make(map[string]*db.Transfer, 0)
	//tokens := make(map[string]*db.Token, 0)
	pools := make(map[string]*db.Pool, 0)
	for _, item := range tree {
		transactions[item.Hash.String()] = &db.Transaction{
			Hash:        item.Hash.String(),
			BlockHeight: *block.BlockHeight,
			Time:        uint64(*block.BlockTime),
		}
		for _, item1 := range item.Instructions {
			t.processInstruction(item1, trades, transfers, pools)
		}
	}
}

func (t *Handler) processInstruction(in *types.InstructionNode, trades map[string]*db.Trade, transfers map[string]*db.Transfer, pools map[string]*db.Pool) () {
	switch in.Instruction.ProgramId {
	case types.System:
		t.processSystemTransfer(in.Instruction)
	case types.Token:
		t.processTokenTransfer(in.Instruction)
	case types.RaydiumAMM:
		t.processRaydiumAmmTrade(in)
	case types.RaydiumClmm:
		t.processRaydiumClmmTrade(in)
	case types.Whirl:
		t.processWhirlTrade(in)
	default:
		for _, item := range in.Children {
			t.processInstruction(item, trades, transfers, pools)
		}
	}
}

func (t *Handler) processSystemTransfer(in *rpc.ParsedInstruction) *db.Transfer {
	return nil
}

func (t *Handler) processTokenTransfer(in *rpc.ParsedInstruction) *db.Transfer {
	return nil
}

func (t *Handler) processRaydiumAmmTrade(in *types.InstructionNode) *db.Trade {
	return nil
}

func (t *Handler) processRaydiumClmmTrade(in *types.InstructionNode) *db.Trade {
	return nil
}

func (t *Handler) processWhirlTrade(in *types.InstructionNode) *db.Trade {
	return nil
}
