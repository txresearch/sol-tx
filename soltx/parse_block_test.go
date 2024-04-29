package soltx

import (
	"context"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"testing"
)

func Test_ParseBlock(t *testing.T) {
	client := rpc.New(rpc.MainNetBeta_RPC)
	rewards := false
	version := uint64(0)
	client.GetParsedTransaction()
	r, err := client.GetBlockWithOpts(
		context.Background(),
		262286706,
		&rpc.GetBlockOpts{
			Encoding:                       solana.EncodingJSONParsed,
			TransactionDetails:             rpc.TransactionDetailsFull,
			Rewards:                        &rewards,
			Commitment:                     rpc.CommitmentConfirmed,
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
