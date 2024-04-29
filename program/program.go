package program

import "github.com/gagliardetto/solana-go"

var (
	Compute         = solana.MustPublicKeyFromBase58("ComputeBudget111111111111111111111111111111")
	Token           = solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	System          = solana.MustPublicKeyFromBase58("11111111111111111111111111111111")
	SysRent         = solana.MustPublicKeyFromBase58("SysvarRent111111111111111111111111111111111")
	AssociatedToken = solana.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL")
	TokenMetadata   = solana.MustPublicKeyFromBase58("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s")
	Token2022       = solana.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb")
	Memo            = solana.MustPublicKeyFromBase58("MemoSq4gqABAXKb96qnH8TysNcWxMyWCqXgDLGmfcHr")
	Whirl           = solana.MustPublicKeyFromBase58("whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc")
	RaydiumClmm     = solana.MustPublicKeyFromBase58("CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK")
	RaydiumAMM      = solana.MustPublicKeyFromBase58("675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8")
	Vote            = solana.MustPublicKeyFromBase58("Vote111111111111111111111111111111111111111")
)

type Instruction struct {
	Program  solana.PublicKey
	Accounts []*solana.AccountMeta
	Data     solana.Base58
}

type InstructionNode struct {
	StackHeight int
	Seq         int
	Instruction *Instruction
	Children    []*InstructionNode
}
