package db

import "github.com/shopspring/decimal"

type Trade struct {
	Id           uint64
	TxHash       string
	TxSeq        uint64
	BlockHeight  uint64
	BlockSeq     uint64
	PoolId       uint64
	Type         string
	TokenAId     uint64
	TokenAAmount decimal.Decimal
	TokenBId     uint64
	TokenBAmount decimal.Decimal
	User         string
}

type Token struct {
	Id      uint64
	Name    string
	Symbol  string
	Decimal uint64
}

type UserOwner struct {
	User  string
	Owner string
}

type Pool struct {
	Id      uint64
	TokenA  uint64
	TokenB  uint64
	TokenLp uint64
}
