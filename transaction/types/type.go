package types

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/shopspring/decimal"
)

type Meta struct {
	Accounts     []rpc.ParsedMessageAccount
	TokenOwner   map[solana.PublicKey]solana.PublicKey
	TokenMint    map[solana.PublicKey]solana.PublicKey
	PreBalance   map[solana.PublicKey]decimal.Decimal
	PostBalance  map[solana.PublicKey]decimal.Decimal
	ErrorMessage []byte
}

type Instruction interface {
	GetEvent() interface{}
	GetChildren(i int) Instruction
	GetInstruction() *rpc.ParsedInstruction
}

type ActionParser interface {
	Parse(in Instruction, meta *Meta) interface{}
}
