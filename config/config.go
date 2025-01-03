package config

import "github.com/gagliardetto/solana-go"

type Node struct {
	Enable bool
	Rpc    string
	Ws     string
}

type BlockHandler struct {
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
	Mainnet      bool
	BlockHandler BlockHandler
	Programs     []Program
	Dao          Dao
}
