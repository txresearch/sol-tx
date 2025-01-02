package block

import (
	"context"
	"github.com/gagliardetto/solana-go/rpc"
	"testing"
	"time"
)

type BlockCallback struct {
}

func (cb *BlockCallback) OnBlock(block *rpc.GetParsedBlockResult) error {
	return nil
}

func TestBlock_Start(t *testing.T) {
	var cb BlockCallback
	bs := New(context.Background(), []string{rpc.MainNetBeta_RPC}, []string{rpc.MainNetBeta_WS}, &cb)
	bs.Subscribe()
	//
	time.Sleep(time.Second * 100000)
}

func TestBlock_ReFetch(t *testing.T) {
	var cb BlockCallback
	bs := New(context.Background(), []string{rpc.MainNetBeta_RPC}, []string{rpc.MainNetBeta_WS}, &cb)
	bs.ReFetch(311363595, 311363601)
	//
	time.Sleep(time.Second * 100000)
}
