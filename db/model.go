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
	Pool         string
	Type         string
	TokenAAmount decimal.Decimal
	TokenBAmount decimal.Decimal
	User         string
}

type Transfer struct {
	TxHash      string `gorm:"type:varchar;not null;uniqueIndex:trade_index"`
	TxSeq       uint64 `gorm:"type:uint;not null;uniqueIndex:trade_index"`
	BlockHeight uint64
	BlockSeq    uint64
	Mint        string
	Amount      uint64
	User        string
}

type Mint struct {
	Hash        string `gorm:"type:varchar;not null;uniqueIndex:token_index"`
	Owner       string
	Name        string
	Symbol      string
	Decimal     uint64
	TotalSupply uint64
}

type Token struct {
	User  string `gorm:"type:varchar;not null;uniqueIndex:user_index"`
	Owner string
}

type Pool struct {
	Hash     string `gorm:"type:varchar;not null;uniqueIndex:pool_index"`
	MintA    string
	MintB    string
	MintLp   string
	VaultA   string
	VaultB   string
	VaultLp  string
	ReserveA uint64
	ReserveB uint64
}
