package soltx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"os"
	"testing"
)

func Test_ParseBlock(t *testing.T) {
	client := rpc.New(rpc.MainNetBeta_RPC)
	rewards := false
	version := uint64(0)
	r, err := client.GetParsedBlockWithOpts(
		context.Background(),
		263991594,
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
	rJson, _ := json.MarshalIndent(r, "", "    ")
	os.WriteFile(fmt.Sprintf("block.json"), rJson, 0644)
	//fmt.Printf("%v\n", r)
	h := New(context.Background(), "./../env")
	txTrees := h.parseBlock(r)
	//
	for _, txTree := range txTrees {
		txTreeJson, _ := json.MarshalIndent(txTree, "", "    ")
		os.WriteFile(fmt.Sprintf("%s.json", txTree.Hash.String()), txTreeJson, 0644)
	}
}

func Test_ProcessBlock(t *testing.T) {
	client := rpc.New(rpc.MainNetBeta_RPC)
	rewards := false
	version := uint64(0)
	r, err := client.GetParsedBlockWithOpts(
		context.Background(),
		264352375,
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
	//fmt.Printf("%v\n", r)
	h := New(context.Background(), "./../env")
	h.processBlock(r)
}
