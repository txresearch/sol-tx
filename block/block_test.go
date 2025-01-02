package block

import (
	"context"
	"github.com/gagliardetto/solana-go/rpc"
	"testing"
	"time"
)

type BlockSubscribeCallback struct {
}

func (cb *BlockSubscribeCallback) OnBlock(block *rpc.GetParsedBlockResult) error {
	return nil
}

func TestBlockSubscribe_Start(t *testing.T) {
	var cb BlockSubscribeCallback
	bs := New(context.Background(), []string{rpc.MainNetBeta_RPC}, []string{rpc.MainNetBeta_WS}, &cb)
	bs.Subscribe()
	//
	time.Sleep(time.Second * 100000)
}

func TestBlockSubscribe_Reset(t *testing.T) {
	var cb BlockSubscribeCallback
	bs := New(context.Background(), []string{rpc.MainNetBeta_RPC}, []string{rpc.MainNetBeta_WS}, &cb)
	bs.ReFetch(311363595, 311363601)
	//
	time.Sleep(time.Second * 100000)
}
