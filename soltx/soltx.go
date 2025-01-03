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
	"github.com/sol-tx/block"
	"github.com/sol-tx/config"
	"github.com/sol-tx/db"
	"github.com/sol-tx/log"
	raydium_amm "github.com/sol-tx/raydiumamm/generated"
	amm_v3 "github.com/sol-tx/raydiumclmm/generated"
	"github.com/sol-tx/transaction"
	"github.com/sol-tx/types"
	whirlpool "github.com/sol-tx/whirlpool/generated"
	"os"
	"time"
)

type Handler struct {
	ctx           context.Context
	log           hclog.Logger
	blockHandler  *block.Block
	updatedBlocks chan *rpc.GetParsedBlockResult
	dao           *db.Dao
}

func newBlockHandle(ctx context.Context, cfg config.BlockHandler, cb block.Callback) *block.Block {
	rpcUrls := make([]string, 0)
	wsUrls := make([]string, 0)
	for _, node := range cfg.Nodes {
		if node.Enable {
			rpcUrls = append(rpcUrls, node.Rpc)
			wsUrls = append(wsUrls, node.Ws)
		}
	}
	return block.New(ctx, rpcUrls, wsUrls, cb)
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
	t := &Handler{
		ctx:           ctx,
		log:           log.NewLog("sol-tx"),
		updatedBlocks: make(chan *rpc.GetParsedBlockResult, 1024),
	}
	t.blockHandler = newBlockHandle(ctx, cfg.BlockHandler, t)
	return t
}

func (h *Handler) Start() error {
	h.blockHandler.Subscribe()
	go h.process()
	return nil
}

func (h *Handler) Service() {
	h.Start()
	<-h.ctx.Done()
	h.Stop()
}

func (h *Handler) Stop() error {
	h.blockHandler.Unsubscribe()
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
		case blockResult := <-h.updatedBlocks:
			h.processBlock(blockResult)
		case <-h.ctx.Done():
			return
		}
	}
}

func (h *Handler) parseBlock(block *rpc.GetParsedBlockResult) []*transaction.Transaction {
	h.log.Info("block handler", "hash", block.Blockhash, "height", *block.BlockHeight, "time", block.BlockTime.Time().UTC().Format("2006-01-02 15:04:05"))
	myTxs := make([]*transaction.Transaction, 0)
	for i, tx := range block.Transactions {
		myTx := transaction.New()
		err := myTx.Parse(&tx)
		if err != nil {
			h.log.Error("failed to parse transaction", "error", err)
			continue
		}
		myTx.Seq = i + 1
		if len(myTx.Instructions) == 0 {
			// no instruction
			continue
		}
		myTxs = append(myTxs, myTx)
	}
	return myTxs
}

func baseTxSeq(depth int) uint64 {
	switch depth {
	case 5:
		return types.Pow101
	case 4:
		return types.Pow102
	case 3:
		return types.Pow103
	case 2:
		return types.Pow102
	case 1:
		return types.Pow105
	default:
		return 1
	}
}

func (h *Handler) processBlock(block *rpc.GetParsedBlockResult) {
	h.log.Info("process block begin", "height", *block.BlockHeight, "time", time.Now().UnixNano())
	defer func() {
		h.log.Info("process block end", "height", *block.BlockHeight, "time", time.Now().UnixNano())
	}()
	b := &db.Block{
		Height: *block.BlockHeight,
		Time:   uint64(*block.BlockTime),
		Hash:   block.Blockhash.String(),
		Slot:   0,
	}
	txes := h.parseBlock(block)
	hash2Transactions := make(map[string]*db.Transaction, 0)
	for _, tx := range txes {
		t := &db.Transaction{
			Hash:        tx.Hash.String(),
			BlockHeight: *block.BlockHeight,
			BlockSeq:    uint64(tx.Seq),
			Time:        uint64(*block.BlockTime),
		}
		hash2Transactions[tx.Hash.String()] = t
	}
	// save hash2Transactions
	hash2Transactions := make(map[string]*db.Transaction, 0)
	id2Trades := make(map[string]*db.Trade, 0)
	id2Transfers := make(map[string]*db.Transfer, 0)
	hash2Pools := make(map[string]*db.Pool, 0)
	for _, tx := range txes {
		t := &db.Transaction{
			Hash:        tx.Hash.String(),
			BlockHeight: *block.BlockHeight,
			BlockSeq:    uint64(tx.Seq),
			Time:        uint64(*block.BlockTime),
		}
		hash2Transactions[tx.Hash.String()] = t
		for _, in := range tx.Instructions {
			h.processInstruction(in, tx.TokenAccountOwner, tx.TokenAccountMint, b, t, 0, id2Trades, id2Transfers, hash2Pools)
		}
	}
	// save all
	blocks := make([]*db.Block, 0)
	blocks = append(blocks, b)
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
	if len(blocks) > 0 {
		h.dao.SaveBlock(blocks)
	}
	if len(transactions) > 0 {
		h.dao.SaveTransaction(transactions)
	}
	if len(trades) > 0 {
		h.dao.SaveTrade(trades)
	}
	if len(transfers) > 0 {
		h.dao.SaveTransfer(transfers)
	}
	if len(pools) > 0 {
		h.dao.SavePool(pools)
	}
}

func (h *Handler) processInstruction(inn *transaction.Instruction,
	tokenAccountOwner map[solana.PublicKey]solana.PublicKey, tokenAccountMint map[solana.PublicKey]solana.PublicKey,
	b *db.Block, t *db.Transaction, parentSeq uint64,
	trades map[string]*db.Trade, transfers map[string]*db.Transfer, pools map[string]*db.Pool) {
	if t.Hash == "BpF8KatKR2KBzfsU3jqphcQnwFR2UdQqenUHzJDypWfuMgaNj6WXWCbWTXJdvCZUfzMpXyB6mLpdLC2MEmavb1n" {
		h.log.Info("xxxxxxx")
	}
	switch inn.Instruction.ProgramId {
	case types.System:
		h.processSystemTransfer(inn, tokenAccountOwner, tokenAccountMint, b, t, parentSeq, trades, transfers, pools)
	case types.Token:
		h.processTokenTransfer(inn, tokenAccountOwner, tokenAccountMint, b, t, parentSeq, trades, transfers, pools)
	case types.RaydiumAMM:
		h.processRaydiumAmmTrade(inn, tokenAccountOwner, tokenAccountMint, b, t, parentSeq, trades, transfers, pools)
	case types.RaydiumClmm:
		h.processRaydiumClmmTrade(inn, tokenAccountOwner, tokenAccountMint, b, t, parentSeq, trades, transfers, pools)
	case types.WhirlPool:
		h.processWhirlPoolTrade(inn, tokenAccountOwner, tokenAccountMint, b, t, parentSeq, trades, transfers, pools)
	default:
		for _, item := range inn.Children {
			seq := parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq)
			h.processInstruction(item, tokenAccountOwner, tokenAccountMint, b, t, seq, trades, transfers, pools)
		}
	}
}

func (h *Handler) processSystemTransfer(inn *types.InstructionNode,
	tokenAccountOwner map[solana.PublicKey]solana.PublicKey, tokenAccountMint map[solana.PublicKey]solana.PublicKey,
	b *db.Block, t *db.Transaction, parentSeq uint64,
	trades map[string]*db.Trade, transfers map[string]*db.Transfer, pools map[string]*db.Pool) {
	type instruction struct {
		Info struct {
			Destination solana.PublicKey `json:"destination"`
			Lamports    uint64           `json:"lamports"`
			Source      solana.PublicKey `json:"source"`
		} `json:"info"`
		T string `json:"type"`
	}
	inJson, _ := inn.Instruction.Parsed.MarshalJSON()
	var in instruction
	json.Unmarshal(inJson, &in)
	switch in.T {
	case "transfer":
		transfer := &db.Transfer{
			BlockHeight: b.Height,
			BlockSeq:    t.BlockSeq,
			TxHash:      t.Hash,
			TxSeq:       parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Mint:        "11111111111111111111111111111111",
			Amount:      in.Info.Lamports,
			From:        in.Info.Source.String(),
			To:          in.Info.Destination.String(),
		}
		transfers[fmt.Sprintf("%s_%d", transfer.TxHash, transfer.TxSeq)] = transfer
	}
}

func (h *Handler) processTokenTransfer(inn *types.InstructionNode,
	tokenAccountOwner map[solana.PublicKey]solana.PublicKey, tokenAccountMint map[solana.PublicKey]solana.PublicKey,
	b *db.Block, t *db.Transaction, parentSeq uint64,
	trades map[string]*db.Trade, transfers map[string]*db.Transfer, pools map[string]*db.Pool) {
	type instruction struct {
		Info struct {
			Destination solana.PublicKey `json:"destination"`
			Lamports    uint64           `json:"amount,string"`
			Source      solana.PublicKey `json:"source"`
			Authority   solana.PublicKey `json:"authority"`
			Mint        solana.PublicKey `json:"mint"`
			TokenAmount struct {
				Amount   uint64 `json:"amount,string"`
				Decimals uint64 `json:"decimals"`
			} `json:"tokenAmount"`
		} `json:"info"`
		T string `json:"type"`
	}
	inJson, _ := inn.Instruction.Parsed.MarshalJSON()
	var in instruction
	json.Unmarshal(inJson, &in)
	switch in.T {
	case "transfer":
		mint := tokenAccountMint[in.Info.Source]
		from := in.Info.Source
		if k, ok := tokenAccountOwner[in.Info.Source]; ok {
			from = k
		}
		to := in.Info.Destination
		if k, ok := tokenAccountOwner[in.Info.Destination]; ok {
			to = k
		}
		transfer := &db.Transfer{
			BlockHeight: b.Height,
			BlockSeq:    t.BlockSeq,
			TxHash:      t.Hash,
			TxSeq:       parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Mint:        mint.String(),
			Amount:      in.Info.Lamports,
			From:        from.String(),
			To:          to.String(),
		}
		transfers[fmt.Sprintf("%s_%d", transfer.TxHash, transfer.TxSeq)] = transfer
	case "transferChecked":
		mint := tokenAccountMint[in.Info.Source]
		from := in.Info.Source
		if k, ok := tokenAccountOwner[in.Info.Source]; ok {
			from = k
		}
		to := in.Info.Destination
		if k, ok := tokenAccountOwner[in.Info.Destination]; ok {
			to = k
		}
		transfer := &db.Transfer{
			BlockHeight: b.Height,
			BlockSeq:    t.BlockSeq,
			TxHash:      t.Hash,
			TxSeq:       parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Mint:        mint.String(),
			Amount:      in.Info.TokenAmount.Amount,
			From:        from.String(),
			To:          to.String(),
		}
		transfers[fmt.Sprintf("%s_%d", transfer.TxHash, transfer.TxSeq)] = transfer
	}
}

func (h *Handler) getTokenTransfer(inn *types.InstructionNode,
	tokenAccountOwner map[solana.PublicKey]solana.PublicKey, tokenAccountMint map[solana.PublicKey]solana.PublicKey) *db.Transfer {
	type instruction struct {
		Info struct {
			Destination solana.PublicKey `json:"destination"`
			Lamports    uint64           `json:"amount,string"`
			Source      solana.PublicKey `json:"source"`
			Authority   solana.PublicKey `json:"authority"`
			Mint        solana.PublicKey `json:"mint"`
			TokenAmount struct {
				Amount   uint64 `json:"amount,string"`
				Decimals uint64 `json:"decimals"`
			} `json:"tokenAmount"`
		} `json:"info"`
		T string `json:"type"`
	}
	inJson, _ := inn.Instruction.Parsed.MarshalJSON()
	var in instruction
	json.Unmarshal(inJson, &in)
	switch in.T {
	case "transfer":
		mint := tokenAccountMint[in.Info.Source]
		from := in.Info.Source
		if k, ok := tokenAccountOwner[in.Info.Source]; ok {
			from = k
		}
		to := in.Info.Destination
		if k, ok := tokenAccountOwner[in.Info.Destination]; ok {
			to = k
		}
		transfer := &db.Transfer{
			Mint:   mint.String(),
			Amount: in.Info.Lamports,
			From:   from.String(),
			To:     to.String(),
		}
		return transfer
	case "transferChecked":
		mint := tokenAccountMint[in.Info.Source]
		from := in.Info.Source
		if k, ok := tokenAccountOwner[in.Info.Source]; ok {
			from = k
		}
		to := in.Info.Destination
		if k, ok := tokenAccountOwner[in.Info.Destination]; ok {
			to = k
		}
		transfer := &db.Transfer{
			Mint:   mint.String(),
			Amount: in.Info.TokenAmount.Amount,
			From:   from.String(),
			To:     to.String(),
		}
		return transfer
	default:
		return nil
	}
}

func (h *Handler) processRaydiumAmmTrade(inn *types.InstructionNode,
	tokenAccountOwner map[solana.PublicKey]solana.PublicKey, tokenAccountMint map[solana.PublicKey]solana.PublicKey,
	b *db.Block, t *db.Transaction, parentSeq uint64,
	trades map[string]*db.Trade, transfers map[string]*db.Transfer, pools map[string]*db.Pool) {
	inst := new(raydium_amm.Instruction)
	err := ag_binary.NewBorshDecoder(inn.Instruction.Data).Decode(inst)
	if err != nil {
		return
	}
	accounts := make([]*solana.AccountMeta, 0)
	for _, item := range inn.Instruction.Accounts {
		accounts = append(accounts, &solana.AccountMeta{
			PublicKey:  item,
			IsWritable: false,
			IsSigner:   false,
		})
	}
	switch inst.TypeID.Uint8() {
	case raydium_amm.Instruction_Initialize2:
		inst1 := inst.Impl.(*raydium_amm.Initialize2)
		inst1.SetAccounts(accounts)
		//
		t1 := h.getTokenTransfer(inn.Children[0], tokenAccountOwner, tokenAccountMint)
		t2 := h.getTokenTransfer(inn.Children[1], tokenAccountOwner, tokenAccountMint)
		//
		trade := &db.Trade{
			BlockHeight:  b.Height,
			BlockSeq:     t.BlockSeq,
			TxHash:       t.Hash,
			TxSeq:        parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Pool:         inst1.GetAmmAccount().PublicKey.String(),
			Type:         db.CreatePool,
			TokenAAmount: decimal.NewFromInt(int64(t1.Amount)),
			TokenBAmount: decimal.NewFromInt(int64(t2.Amount)),
			User:         inst1.GetUserWalletAccount().PublicKey.String(),
		}
		trades[fmt.Sprintf("%s_%d", trade.TxHash, trade.TxSeq)] = trade
		//
		pool := &db.Pool{
			Hash:     inst1.GetAmmAccount().PublicKey.String(),
			MintA:    inst1.GetCoinMintAccount().PublicKey.String(),
			MintB:    inst1.GetPcMintAccount().PublicKey.String(),
			MintLp:   inst1.GetLpMintAccount().PublicKey.String(),
			VaultA:   inst1.GetPoolCoinTokenAccountAccount().PublicKey.String(),
			VaultB:   inst1.GetPoolPcTokenAccountAccount().PublicKey.String(),
			VaultLp:  "",
			ReserveA: 0,
			ReserveB: 0,
		}
		pools[pool.Hash] = pool
	case raydium_amm.Instruction_Deposit:
		inst1 := inst.Impl.(*raydium_amm.Deposit)
		inst1.SetAccounts(accounts)
		//
		t1 := h.getTokenTransfer(inn.Children[0], tokenAccountOwner, tokenAccountMint)
		t2 := h.getTokenTransfer(inn.Children[1], tokenAccountOwner, tokenAccountMint)
		//
		trade := &db.Trade{
			BlockHeight:  b.Height,
			BlockSeq:     t.BlockSeq,
			TxHash:       t.Hash,
			TxSeq:        parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Pool:         inst1.GetAmmAccount().PublicKey.String(),
			Type:         db.AddLiquidity,
			TokenAAmount: decimal.NewFromInt(int64(t1.Amount)),
			TokenBAmount: decimal.NewFromInt(int64(t2.Amount)),
			User:         inst1.GetUserOwnerAccount().PublicKey.String(),
		}
		trades[fmt.Sprintf("%s_%d", trade.TxHash, trade.TxSeq)] = trade
	case raydium_amm.Instruction_SwapBaseIn:
		inst1 := inst.Impl.(*raydium_amm.SwapBaseIn)
		inst1.SetAccounts(accounts)
		if len(inn.Children) < 2 {
			return
		}
		//
		t1 := h.getTokenTransfer(inn.Children[0], tokenAccountOwner, tokenAccountMint)
		t2 := h.getTokenTransfer(inn.Children[1], tokenAccountOwner, tokenAccountMint)
		//
		trade := &db.Trade{
			BlockHeight:  b.Height,
			BlockSeq:     t.BlockSeq,
			TxHash:       t.Hash,
			TxSeq:        parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Pool:         inst1.GetAmmAccount().PublicKey.String(),
			Type:         db.Swap,
			TokenAAmount: decimal.NewFromInt(int64(t1.Amount)),
			TokenBAmount: decimal.NewFromInt(int64(t2.Amount)),
			User:         inst1.GetUserSourceOwnerAccount().PublicKey.String(),
		}
		trades[fmt.Sprintf("%s_%d", trade.TxHash, trade.TxSeq)] = trade
	case raydium_amm.Instruction_SwapBaseOut:
		inst1 := inst.Impl.(*raydium_amm.SwapBaseOut)
		inst1.SetAccounts(accounts)
		//
		t1 := h.getTokenTransfer(inn.Children[0], tokenAccountOwner, tokenAccountMint)
		t2 := h.getTokenTransfer(inn.Children[1], tokenAccountOwner, tokenAccountMint)
		//
		trade := &db.Trade{
			BlockHeight:  b.Height,
			BlockSeq:     t.BlockSeq,
			TxHash:       t.Hash,
			TxSeq:        parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Pool:         inst1.GetAmmAccount().PublicKey.String(),
			Type:         db.Swap,
			TokenAAmount: decimal.NewFromInt(int64(t1.Amount)),
			TokenBAmount: decimal.NewFromInt(int64(t2.Amount)),
			User:         inst1.GetUserSourceOwnerAccount().PublicKey.String(),
		}
		trades[fmt.Sprintf("%s_%d", trade.TxHash, trade.TxSeq)] = trade
	case raydium_amm.Instruction_Withdraw:
		inst1 := inst.Impl.(*raydium_amm.Withdraw)
		inst1.SetAccounts(accounts)
		//
		t1 := h.getTokenTransfer(inn.Children[0], tokenAccountOwner, tokenAccountMint)
		t2 := h.getTokenTransfer(inn.Children[1], tokenAccountOwner, tokenAccountMint)
		//
		trade := &db.Trade{
			BlockHeight:  b.Height,
			BlockSeq:     t.BlockSeq,
			TxHash:       t.Hash,
			TxSeq:        parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Pool:         inst1.GetAmmAccount().PublicKey.String(),
			Type:         db.RemoveLiquidity,
			TokenAAmount: decimal.NewFromInt(int64(t1.Amount)),
			TokenBAmount: decimal.NewFromInt(int64(t2.Amount)),
			User:         inst1.GetUserOwnerAccount().PublicKey.String(),
		}
		trades[fmt.Sprintf("%s_%d", trade.TxHash, trade.TxSeq)] = trade
	}
}

func (h *Handler) processRaydiumClmmTrade(inn *types.InstructionNode,
	tokenAccountOwner map[solana.PublicKey]solana.PublicKey, tokenAccountMint map[solana.PublicKey]solana.PublicKey,
	b *db.Block, t *db.Transaction, parentSeq uint64,
	trades map[string]*db.Trade, transfers map[string]*db.Transfer, pools map[string]*db.Pool) {
	inst := new(amm_v3.Instruction)
	err := ag_binary.NewBorshDecoder(inn.Instruction.Data).Decode(inst)
	if err != nil {
		return
	}
	accounts := make([]*solana.AccountMeta, 0)
	for _, item := range inn.Instruction.Accounts {
		accounts = append(accounts, &solana.AccountMeta{
			PublicKey:  item,
			IsWritable: false,
			IsSigner:   false,
		})
	}
	switch inst.TypeID {
	case amm_v3.Instruction_CreatePool:
		inst1 := inst.Impl.(*amm_v3.CreatePool)
		inst1.SetAccounts(accounts)
		pool := &db.Pool{
			Hash:     inst1.GetPoolStateAccount().PublicKey.String(),
			MintA:    inst1.GetTokenMint0Account().PublicKey.String(),
			MintB:    inst1.GetTokenMint1Account().PublicKey.String(),
			MintLp:   inst1.GetTokenVault1Account().PublicKey.String(),
			VaultA:   inst1.GetTokenVault1Account().PublicKey.String(),
			VaultB:   inst1.GetTokenVault1Account().PublicKey.String(),
			VaultLp:  "",
			ReserveA: 0,
			ReserveB: 0,
		}
		pools[pool.Hash] = pool
	case amm_v3.Instruction_IncreaseLiquidityV2:
		inst1 := inst.Impl.(*amm_v3.IncreaseLiquidityV2)
		inst1.SetAccounts(accounts)
		//
		t1 := h.getTokenTransfer(inn.Children[0], tokenAccountOwner, tokenAccountMint)
		t2 := h.getTokenTransfer(inn.Children[1], tokenAccountOwner, tokenAccountMint)
		//
		trade := &db.Trade{
			BlockHeight:  b.Height,
			BlockSeq:     t.BlockSeq,
			TxHash:       t.Hash,
			TxSeq:        parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Pool:         inst1.GetPoolStateAccount().PublicKey.String(),
			Type:         db.AddLiquidity,
			TokenAAmount: decimal.NewFromInt(int64(t1.Amount)),
			TokenBAmount: decimal.NewFromInt(int64(t2.Amount)),
			User:         inst1.Get(0).PublicKey.String(),
		}
		trades[fmt.Sprintf("%s_%d", trade.TxHash, trade.TxSeq)] = trade
	case amm_v3.Instruction_DecreaseLiquidityV2:
		inst1 := inst.Impl.(*amm_v3.DecreaseLiquidityV2)
		inst1.SetAccounts(accounts)
		//
		t1 := h.getTokenTransfer(inn.Children[0], tokenAccountOwner, tokenAccountMint)
		t2 := h.getTokenTransfer(inn.Children[1], tokenAccountOwner, tokenAccountMint)
		//
		trade := &db.Trade{
			BlockHeight:  b.Height,
			BlockSeq:     t.BlockSeq,
			TxHash:       t.Hash,
			TxSeq:        parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Pool:         inst1.GetPoolStateAccount().PublicKey.String(),
			Type:         db.RemoveLiquidity,
			TokenAAmount: decimal.NewFromInt(int64(t1.Amount)),
			TokenBAmount: decimal.NewFromInt(int64(t2.Amount)),
			User:         inst1.Get(0).PublicKey.String(),
		}
		trades[fmt.Sprintf("%s_%d", trade.TxHash, trade.TxSeq)] = trade
	case amm_v3.Instruction_Swap:
		inst1 := inst.Impl.(*amm_v3.Swap)
		inst1.SetAccounts(accounts)
		//
		t1 := h.getTokenTransfer(inn.Children[0], tokenAccountOwner, tokenAccountMint)
		t2 := h.getTokenTransfer(inn.Children[1], tokenAccountOwner, tokenAccountMint)
		//
		trade := &db.Trade{
			BlockHeight:  b.Height,
			BlockSeq:     t.BlockSeq,
			TxHash:       t.Hash,
			TxSeq:        parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Pool:         inst1.GetPoolStateAccount().PublicKey.String(),
			Type:         db.Swap,
			TokenAAmount: decimal.NewFromInt(int64(t1.Amount)),
			TokenBAmount: decimal.NewFromInt(int64(t2.Amount)),
			User:         inst1.GetPayerAccount().PublicKey.String(),
		}
		trades[fmt.Sprintf("%s_%d", trade.TxHash, trade.TxSeq)] = trade
	case amm_v3.Instruction_SwapV2:
		inst1 := inst.Impl.(*amm_v3.SwapV2)
		inst1.SetAccounts(accounts)
		//
		t1 := h.getTokenTransfer(inn.Children[0], tokenAccountOwner, tokenAccountMint)
		t2 := h.getTokenTransfer(inn.Children[1], tokenAccountOwner, tokenAccountMint)
		//
		trade := &db.Trade{
			BlockHeight:  b.Height,
			BlockSeq:     t.BlockSeq,
			TxHash:       t.Hash,
			TxSeq:        parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Pool:         inst1.GetPoolStateAccount().PublicKey.String(),
			Type:         db.Swap,
			TokenAAmount: decimal.NewFromInt(int64(t1.Amount)),
			TokenBAmount: decimal.NewFromInt(int64(t2.Amount)),
			User:         inst1.Get(0).PublicKey.String(),
		}
		trades[fmt.Sprintf("%s_%d", trade.TxHash, trade.TxSeq)] = trade
	}
}

func (h *Handler) processWhirlPoolTrade(inn *types.InstructionNode,
	tokenAccountOwner map[solana.PublicKey]solana.PublicKey, tokenAccountMint map[solana.PublicKey]solana.PublicKey,
	b *db.Block, t *db.Transaction, parentSeq uint64,
	trades map[string]*db.Trade, transfers map[string]*db.Transfer, pools map[string]*db.Pool) {
	inst := new(whirlpool.Instruction)
	err := ag_binary.NewBorshDecoder(inn.Instruction.Data).Decode(inst)
	if err != nil {
		return
	}
	accounts := make([]*solana.AccountMeta, 0)
	for _, item := range inn.Instruction.Accounts {
		accounts = append(accounts, &solana.AccountMeta{
			PublicKey:  item,
			IsWritable: false,
			IsSigner:   false,
		})
	}
	switch inst.TypeID {
	case whirlpool.Instruction_InitializePool:
		inst1 := inst.Impl.(*whirlpool.InitializePool)
		inst1.SetAccounts(accounts)
		//
		trade := &db.Trade{
			BlockHeight:  b.Height,
			BlockSeq:     t.BlockSeq,
			TxHash:       t.Hash,
			TxSeq:        parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Pool:         inst1.GetWhirlpoolAccount().PublicKey.String(),
			Type:         db.Swap,
			TokenAAmount: decimal.NewFromInt(0),
			TokenBAmount: decimal.NewFromInt(0),
			User:         inst1.GetFunderAccount().PublicKey.String(),
		}
		trades[fmt.Sprintf("%s_%d", trade.TxHash, trade.TxSeq)] = trade
	case whirlpool.Instruction_Swap:
		inst1 := inst.Impl.(*whirlpool.Swap)
		inst1.SetAccounts(accounts)
		//
		t1 := h.getTokenTransfer(inn.Children[0], tokenAccountOwner, tokenAccountMint)
		t2 := h.getTokenTransfer(inn.Children[1], tokenAccountOwner, tokenAccountMint)
		//
		trade := &db.Trade{
			BlockHeight:  b.Height,
			BlockSeq:     t.BlockSeq,
			TxHash:       t.Hash,
			TxSeq:        parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Pool:         inst1.GetWhirlpoolAccount().PublicKey.String(),
			Type:         db.Swap,
			TokenAAmount: decimal.NewFromInt(int64(t1.Amount)),
			TokenBAmount: decimal.NewFromInt(int64(t2.Amount)),
			User:         inst1.GetTokenAuthorityAccount().PublicKey.String(),
		}
		trades[fmt.Sprintf("%s_%d", trade.TxHash, trade.TxSeq)] = trade
	case whirlpool.Instruction_IncreaseLiquidity:
		inst1 := inst.Impl.(*whirlpool.IncreaseLiquidity)
		inst1.SetAccounts(accounts)
		//
		t1 := h.getTokenTransfer(inn.Children[0], tokenAccountOwner, tokenAccountMint)
		t2 := h.getTokenTransfer(inn.Children[1], tokenAccountOwner, tokenAccountMint)
		//
		trade := &db.Trade{
			BlockHeight:  b.Height,
			BlockSeq:     t.BlockSeq,
			TxHash:       t.Hash,
			TxSeq:        parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Pool:         inst1.GetWhirlpoolAccount().PublicKey.String(),
			Type:         db.AddLiquidity,
			TokenAAmount: decimal.NewFromInt(int64(t1.Amount)),
			TokenBAmount: decimal.NewFromInt(int64(t2.Amount)),
			User:         inst1.GetPositionAuthorityAccount().PublicKey.String(),
		}
		trades[fmt.Sprintf("%s_%d", trade.TxHash, trade.TxSeq)] = trade
	case whirlpool.Instruction_DecreaseLiquidity:
		inst1 := inst.Impl.(*whirlpool.DecreaseLiquidity)
		inst1.SetAccounts(accounts)
		//
		t1 := h.getTokenTransfer(inn.Children[0], tokenAccountOwner, tokenAccountMint)
		t2 := h.getTokenTransfer(inn.Children[1], tokenAccountOwner, tokenAccountMint)
		//
		trade := &db.Trade{
			BlockHeight:  b.Height,
			BlockSeq:     t.BlockSeq,
			TxHash:       t.Hash,
			TxSeq:        parentSeq + baseTxSeq(inn.Instruction.StackHeight)*uint64(inn.Seq),
			Pool:         inst1.GetWhirlpoolAccount().PublicKey.String(),
			Type:         db.RemoveLiquidity,
			TokenAAmount: decimal.NewFromInt(int64(t1.Amount)),
			TokenBAmount: decimal.NewFromInt(int64(t2.Amount)),
			User:         inst1.GetPositionAuthorityAccount().PublicKey.String(),
		}
		trades[fmt.Sprintf("%s_%d", trade.TxHash, trade.TxSeq)] = trade
	}
}
