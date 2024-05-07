package types

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"math"
)

var (
	Pow105 = uint64(math.Pow10(15))
	Pow104 = uint64(math.Pow10(12))
	Pow103 = uint64(math.Pow10(9))
	Pow102 = uint64(math.Pow10(6))
	Pow101 = uint64(math.Pow10(3))
)

var (
	Compute         = solana.MustPublicKeyFromBase58("ComputeBudget111111111111111111111111111111")
	Token           = solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	System          = solana.MustPublicKeyFromBase58("11111111111111111111111111111111")
	SysRent         = solana.MustPublicKeyFromBase58("SysvarRent111111111111111111111111111111111")
	AssociatedToken = solana.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL")
	TokenMetadata   = solana.MustPublicKeyFromBase58("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s")
	Token2022       = solana.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb")
	Memo            = solana.MustPublicKeyFromBase58("MemoSq4gqABAXKb96qnH8TysNcWxMyWCqXgDLGmfcHr")
	WhirlPool       = solana.MustPublicKeyFromBase58("whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc")
	RaydiumClmm     = solana.MustPublicKeyFromBase58("CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK")
	RaydiumAMM      = solana.MustPublicKeyFromBase58("675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8")
	Vote            = solana.MustPublicKeyFromBase58("Vote111111111111111111111111111111111111111")
)

type Transaction struct {
	Hash              solana.Signature `json:"hash"`
	TokenAccountOwner map[solana.PublicKey]solana.PublicKey
	TokenAccountMint  map[solana.PublicKey]solana.PublicKey
	Instructions      []*InstructionNode `json:"instructions"`
	Seq               int
}

type InstructionNode struct {
	Seq         int                    `json:"seq"`
	Instruction *rpc.ParsedInstruction `json:"instruction"`
	Children    []*InstructionNode     `json:"children,omitempty"`
}
