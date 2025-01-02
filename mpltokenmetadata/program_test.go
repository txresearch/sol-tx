package mpltokenmetadata

import (
	"context"
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	token_metadata "github.com/sol-tx/mpltokenmetadata/generated"
	"testing"
)

func TestProgram_CreateTokenMeta(t *testing.T) {
	mint := solana.MustPublicKeyFromBase58("ArvwdkRujBSmyCxPPKHqgDhJbmh4mWnmZPnh7rW2ceeN")
	tokenMeta, _, _ := solana.FindTokenMetadataAddress(mint)
	fmt.Printf("%s\n", tokenMeta)
	//
}

func TestProgram_TokenMetaAddress(t *testing.T) {
	//mint := solana.MustPublicKeyFromBase58("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")
	mint := solana.MustPublicKeyFromBase58("So11111111111111111111111111111111111111112")
	tokenMeta, _, _ := solana.FindTokenMetadataAddress(mint)
	fmt.Printf("%s\n", tokenMeta)

	rpcClient := rpc.New(rpc.DevNet_RPC)
	w := wallet.New("H78q1dYMDDJsrcuuVgTPvXNW15Hx4T45AzFyNxXToZrtxKTXVckDkL9YcHuefJZnmhEYUNS5iQDJu8q2r8TTkML")

	//
	client := rpc.New(rpc.MainNetBeta_RPC)
	account, _ := client.GetAccountInfo(context.Background(), tokenMeta)
	data := account.Value.Data.GetBinary()
	meta := token_metadata.Metadata{}
	err := meta.UnmarshalWithDecoder(ag_binary.NewBinDecoder(data))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", meta)

}
