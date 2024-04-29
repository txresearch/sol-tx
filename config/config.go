package config

import "github.com/gagliardetto/solana-go"

type Node struct {
	Enable bool
	Rpc    string
	Ws     string
}

type BlockSubscribe struct {
	Nodes []Node
}

type Program struct {
	Enable bool
	Name   string
	Id     solana.PublicKey
}

type Dao struct {
	User     string
	Password string
	Url      string
	Port     uint
	Scheme   string
	Debug    bool
}

type Config struct {
	Mainnet        bool
	BlockSubscribe BlockSubscribe
	Programs       []Program
	Dao            Dao
}
