package soltx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/go-hclog"
	"github.com/sol-tx/config"
	"github.com/sol-tx/log"
	"os"
)

type Handler struct {
	ctx context.Context
	log hclog.Logger
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
		ctx: ctx,
		log: log.NewLog("sol-tx"),
	}
	//
	return t
}

func (t *Handler) Start() error {
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
	//t.updatedBlocks <- block
	return nil
}
