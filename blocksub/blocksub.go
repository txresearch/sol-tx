package blocksub

import (
	"context"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/hashicorp/go-hclog"
	"github.com/sol-tx/log"
	"time"
)

type Callback interface {
	OnBlock(block *rpc.GetBlockResult) error
}

type BlockSubscribe struct {
	ctx        context.Context
	log        hclog.Logger
	rpcClients []*rpc.Client
	wsUrls     []string
	cb         Callback
	blockSub   *ws.BlockSubscription
}

func New(ctx context.Context, rpcUrls []string, wsUrls []string, cb Callback) *BlockSubscribe {
	rpcClients := make([]*rpc.Client, 0)
	for _, rpcUrl := range rpcUrls {
		rpcClients = append(rpcClients, rpc.New(rpcUrl))
	}
	b := &BlockSubscribe{
		ctx:        ctx,
		log:        log.NewLog("block.subscribe"),
		rpcClients: rpcClients,
		wsUrls:     wsUrls,
		cb:         cb,
	}
	return b
}

func (b *BlockSubscribe) Start() {
	b.log.Info("start......")
	go b.subscribeBlock()
}

func (b *BlockSubscribe) Stop() {
	if b.blockSub != nil {
		b.blockSub.Unsubscribe()
	}
	b.log.Info("stop......")
}

func (b *BlockSubscribe) Reset(start uint64, end uint64) {
	go b.fetchBlock(start, end)
}

func (b *BlockSubscribe) subscribeBlock() {
	process := func(used int) (exit bool) {
		wsclient, err := ws.Connect(b.ctx, b.wsUrls[used])
		if err != nil {
			b.log.Error("ws connect", "error", err)
			return false
		}
		rewards := false
		version := uint64(0)
		sub, err := wsclient.BlockSubscribe(ws.NewBlockSubscribeFilterAll(), &ws.BlockSubscribeOpts{
			Commitment:                     rpc.CommitmentFinalized,
			Encoding:                       solana.EncodingJSONParsed, // transaction meta will in response
			TransactionDetails:             rpc.TransactionDetailsFull,
			Rewards:                        &rewards,
			MaxSupportedTransactionVersion: &version,
		})
		if err != nil {
			b.log.Error("BlockSubscribe", "error", err)
			return false
		}
		b.blockSub = sub
		b.log.Info("subscribe block successful")
		for {
			got, err := sub.Recv()
			if err != nil {
				b.log.Error("recv", "error", err)
				return false
			}
			if got == nil {
				// exit
				return true
			}
			b.log.Info("receive block", "slot", got.Context.Slot)
			if got.Value.Err != nil {
				b.log.Info("got value", "error", got.Value.Err)
				return false
			}
			block := got.Value.Block
			b.log.Info("receive block", "slot", got.Value.Slot, "block time", block.BlockTime.String())
			if b.cb != nil {
				b.cb.OnBlock(block)
			}
		}
	}
	used := 0
	for true {
		b.log.Info("try to subscribe block")
		if exit := process(used); exit {
			return
		}
		used++
		used = used % len(b.wsUrls)
		b.log.Info("subscribe failed, retry......")
		time.Sleep(time.Second * 5)
	}
}

func (b *BlockSubscribe) fetchBlock(start uint64, end uint64) {
	fetch := func(slot uint64) *rpc.GetBlockResult {
		rewards := false
		version := uint64(0)
		block, err := b.rpcClients[0].GetBlockWithOpts(b.ctx, slot, &rpc.GetBlockOpts{
			Encoding:                       solana.EncodingJSONParsed, // transaction meta will in response
			TransactionDetails:             rpc.TransactionDetailsFull,
			Rewards:                        &rewards,
			Commitment:                     rpc.CommitmentConfirmed,
			MaxSupportedTransactionVersion: &version,
		})
		if err != nil {
			b.log.Error("GetBlock", "error", err)
			return nil
		}
		return block
	}
	slot := start
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			block := fetch(slot)
			if block == nil {
				continue
			}
			b.log.Info("fetch block", "block height", block.BlockHeight, "block time", block.BlockTime.String())
			if b.cb != nil {
				b.cb.OnBlock(block)
			}
			slot++
			if slot > end {
				return
			}
		case <-b.ctx.Done():
			return
		}
	}
}
