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
	"github.com/sol-tx/program"
	"os"
)

type Handler struct {
	ctx             context.Context
	log             hclog.Logger
	blockSubscriber *blocksubscribe.BlockSubscribe
	updatedBlocks   chan *rpc.GetBlockResult
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
		updatedBlocks: make(chan *rpc.GetBlockResult, 1024),
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

func (t *Handler) OnBlock(block *rpc.GetBlockResult) error {
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

func subInstructions(ins []solana.CompiledInstruction, inTree *program.InstructionNode) {
	keys := inTree.Instruction.Accounts
	for i, in := range ins {
		programId := keys[in.ProgramIDIndex]
		accounts := make([]*solana.AccountMeta, 0)
		for _, a := range in.AccountsWithKey {
			for _, b := range keys {
				if b.PublicKey == a {
					accounts = append(accounts, b)
				}
			}
		}
		a := &program.Instruction{
			Program:  programId.PublicKey,
			Accounts: accounts,
			Data:     in.Data,
		}
		current := &program.InstructionNode{
			StackHeight: 1,
			Seq:         i,
			Instruction: a,
			Children:    nil,
		}
		inTree.Children = append(inTree.Children, current)
	}
}

func subInnerInstructions(ins []solana.CompiledInstruction, depth int, inTree *program.InstructionNode) {
	var current *program.InstructionNode
	for i, in := range ins {
		if in.StackHeight == depth+1 {
			subInnerInstructions(ins[i:], in.StackHeight, current)
		} else if in.StackHeight == depth-1 {
			return
		} else {
			keys := inTree.Instruction.Accounts
			accounts := make([]*solana.AccountMeta, 0)
			for _, a := range in.AccountsWithKey {
				for _, b := range keys {
					if b.PublicKey == a {
						accounts = append(accounts, b)
					}
				}
			}
			pid := keys[in.ProgramIDIndex].PublicKey
			a := &program.Instruction{
				Program:  pid,
				Accounts: accounts,
				Data:     in.Data,
			}
			current = &program.InstructionNode{
				StackHeight: in.StackHeight,
				Seq:         i,
				Instruction: a,
				Children:    nil,
			}
			inTree.Children = append(inTree.Children, current)
		}
	}
}

func trade(ins []*program.Instruction) {

}

func (t *Handler) parseBlock(block *rpc.GetBlockResult) (*db.Block, []*db.Transaction, []*db.Trade, []*db.Token, []*db.Pool) {
	t.log.Info("blockHandler", "hash", block.Blockhash, "time", block.BlockTime.Time().UTC().Format("2006-01-02 15:04:05"))
	// must with meta - json parsed
	// block
	b := &db.Block{
		Height: *block.BlockHeight,
		Time:   uint64(*block.BlockTime),
		Hash:   block.Blockhash.String(),
		Slot:   0,
	}
	// save transactions
	transactions := make([]*db.Transaction, 0)
	trades := make([]*db.Trade, 0)
	tokens := make([]*db.Token, 0)
	pools := make([]*db.Pool, 0)
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
		parsedTransaction, err := transaction.GetParsedTransaction()
		if err != nil {
			t.log.Error("GetParsedTransaction", "error", err)
			continue
		}
		message := parsedTransaction.Message
		instructions := message.Instructions
		if len(instructions) == 0 {
			t.log.Warn("no instruction")
			continue
		}
		keys, err := message.AccountMetaList()
		if err != nil {
			t.log.Error("GetAllKeys", "error", err)
			continue
		}
		if int(instructions[0].ProgramIDIndex) >= len(keys) {
			t.log.Error("program id invalid")
			continue
		}
		programId := keys[instructions[0].ProgramIDIndex].PublicKey
		if programId == program.Vote {
			continue
		}
		//
		inTree := &program.InstructionNode{
			StackHeight: 0,
			Seq:         0,
			Instruction: &program.Instruction{
				Program:  solana.PublicKey{},
				Accounts: keys,
				Data:     nil,
			},
			Children: nil,
		}
		subInstructions(instructions, inTree)
		//
		innerInstructions := meta.InnerInstructions
		if len(instructions) > 0 {
			for _, innerInstruction := range innerInstructions {
				pinn := inTree.Children[innerInstruction.Index]
				subInnerInstructions(innerInstruction.Instructions, 2, pinn)
			}
		}
	}
	return b, transactions, trades, tokens, pools
}

func (t *Handler) processBlock(block *rpc.GetBlockResult) {
	t.parseBlock(block)
}
