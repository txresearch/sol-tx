package soltx

import (
	"context"
	"github.com/gagliardetto/solana-go"
	rpc2 "github.com/gagliardetto/solana-go/rpc"
	"testing"
)

func Test_ParseBlock(t *testing.T) {
	client := rpc2.New(rpc2.MainNetBeta_RPC)
	rewards := false
	version := uint64(0)
	r, err := client.GetBlockWithOpts(
		context.Background(),
		262286706,
		&rpc2.GetBlockOpts{
			Encoding:                       solana.EncodingJSONParsed,
			TransactionDetails:             rpc2.TransactionDetailsFull,
			Rewards:                        &rewards,
			Commitment:                     rpc2.CommitmentConfirmed,
			MaxSupportedTransactionVersion: &version,
		},
	)
	if err != nil {
		panic(err)
	}
	//
	h := New(context.Background(), "./../env")
	h.parseBlock(r)
}
