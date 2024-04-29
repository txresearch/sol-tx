package soltx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/go-hclog"
	"github.com/sol-tx/blocksubscribe"
	"github.com/sol-tx/config"
	"github.com/sol-tx/db"
	"github.com/sol-tx/log"
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

func (t *Handler) processBlock(block *rpc.GetBlockResult) {
	t.log.Info("blockHandler", "hash", block.Blockhash, "time", block.BlockTime.Time().UTC().Format("2006-01-02 15:04:05"))
}
