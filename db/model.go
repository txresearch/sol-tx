package db

import "github.com/shopspring/decimal"

type Block struct {
	Height uint64 `gorm:"type:varchar;not null;uniqueIndex:block_index"`
	Hash   string
	Time   uint64
	Slot   uint64
}

type Transaction struct {
	Hash        string `gorm:"type:varchar;not null;uniqueIndex:transaction_index"`
	BlockHeight uint64
	Time        uint64
}

type Trade struct {
	TxHash       string `gorm:"type:varchar;not null;uniqueIndex:trade_index"`
	TxSeq        uint64 `gorm:"type:uint;not null;uniqueIndex:trade_index"`
	BlockHeight  uint64
	BlockSeq     uint64
	PoolHash     string
	Type         string
	TokenAHash   string
	TokenAAmount decimal.Decimal
	TokenBHash   string
	TokenBAmount decimal.Decimal
	User         string
}

type Token struct {
	Hash    string `gorm:"type:varchar;not null;uniqueIndex:token_index"`
	Name    string
	Symbol  string
	Decimal uint64
}

type UserOwner struct {
	User  string `gorm:"type:varchar;not null;uniqueIndex:user_index"`
	Owner string
}

type Pool struct {
	Hash    string `gorm:"type:varchar;not null;uniqueIndex:pool_index"`
	TokenA  uint64
	TokenB  uint64
	TokenLp uint64
}
