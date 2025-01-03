package block

import (
	"context"
	"encoding/json"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/hashicorp/go-hclog"
	"github.com/sol-tx/log"
	"strings"
	"time"
)

type Callback interface {
	OnBlock(block *rpc.GetParsedBlockResult) error
}

type Block struct {
	ctx        context.Context
	cb         Callback
	log        hclog.Logger
	rpcClients []*rpc.Client
	wsUrls     []string
	blockSub   *ws.ParsedBlockSubscription
}

func New(ctx context.Context, rpcUrls []string, wsUrls []string, cb Callback) *Block {
	rpcClients := make([]*rpc.Client, 0)
	for _, rpcUrl := range rpcUrls {
		rpcClients = append(rpcClients, rpc.New(rpcUrl))
	}
	b := &Block{
		ctx:        ctx,
		log:        log.NewLog("block.subscribe"),
		rpcClients: rpcClients,
		wsUrls:     wsUrls,
		cb:         cb,
	}
	return b
}

func (b *Block) Subscribe() {
	b.log.Info("start......")
	go b.subscribeBlock()
}

func (b *Block) Unsubscribe() {
	if b.blockSub != nil {
		b.blockSub.Unsubscribe()
	}
	b.log.Info("stop......")
}

func (b *Block) ReFetch(start uint64, end uint64) {
	go b.fetchBlock(start, end)
}

func (b *Block) subscribeBlock() {
	defer func() {
		b.log.Info("subscribe block exit......")
	}()
	process := func(used int) (exit bool) {
		wsClient, err := ws.Connect(b.ctx, b.wsUrls[used])
		if err != nil {
			b.log.Error("ws connect", "error", err)
			return false
		}
		rewards := false
		sub, err := wsClient.ParsedBlockSubscribe(
			ws.NewBlockSubscribeFilterAll(),
			&ws.BlockSubscribeOpts{
				Commitment:                     rpc.CommitmentFinalized,
				Encoding:                       solana.EncodingJSONParsed, // transaction meta will in response
				TransactionDetails:             rpc.TransactionDetailsFull,
				Rewards:                        &rewards,
				MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion1,
			})
		if err != nil {
			b.log.Error("block subscribe", "error", err)
			return false
		}
		b.blockSub = sub
		b.log.Info("subscribe block successful")
		for {
			got, err := b.blockSub.Recv(b.ctx)
			if err != nil {
				b.log.Error("block subscribe recv", "error", err)
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
			b.log.Info("receive block", "slot", got.Value.Slot, "block time", block.BlockTime.Time().UTC().Format("2006-01-02 15:04:05"))
			if b.cb != nil {
				b.cb.OnBlock(block)
			}
		}
	}
	used := 0
	for {
		b.log.Info("try to subscribe block")
		if exit := process(used); exit {
			return
		}
		used++
		used = used % len(b.wsUrls)
		b.log.Info("subscribe block failed, retry......")
		time.Sleep(time.Second * 5)
	}
}

func (b *Block) fetchBlock(start uint64, end uint64) {
	process := func(slot uint64, used int) (*rpc.GetParsedBlockResult, error) {
		rewards := false
		block, err := b.rpcClients[used].GetParsedBlockWithOpts(b.ctx, slot, &rpc.GetBlockOpts{
			Encoding:                       solana.EncodingJSONParsed, // transaction meta will in response
			TransactionDetails:             rpc.TransactionDetailsFull,
			Rewards:                        &rewards,
			Commitment:                     rpc.CommitmentConfirmed,
			MaxSupportedTransactionVersion: &rpc.MaxSupportedTransactionVersion1,
		})
		if err == nil {
			return block, nil
		}
		b.log.Warn("fetch block failed", "slot", slot, "err", err)

		type MyError struct {
			Code    int
			Message string
		}

		errJson, _ := json.Marshal(err)
		myError := MyError{}
		json.Unmarshal(errJson, &myError)

		if myError.Code == -32007 && strings.Contains(myError.Message, "was skipped") {
			b.log.Warn("block is skipped", "slot", slot)
			return nil, nil
		} else {
			return nil, err
		}
	}
	slot := start
	used := 0
	ticker := time.NewTicker(time.Second * 2)
	for {
		if slot > end {
			return
		}
		select {
		case <-ticker.C:
			b.log.Info("try to fetch block", "slot", slot)
			block, err := process(slot, used)
			if err != nil {
				used++
				used = used % len(b.rpcClients)
				continue
			}
			if block == nil {
				b.log.Info("the fetched block is empty", "slot", slot)
				slot++
				continue
			}
			b.log.Info("fetch block", "block height", *block.BlockHeight, "slot", slot, "block time", block.BlockTime.Time().UTC().Format("2006-01-02 15:04:05"))
			if b.cb != nil {
				b.cb.OnBlock(block)
			}
			slot++
		case <-b.ctx.Done():
			return
		}
	}
}
