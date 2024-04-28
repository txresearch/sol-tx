package config

import "github.com/gagliardetto/solana-go"

type Trade struct {
	Symbol string
}

type Node struct {
	Enable bool
	Rpc    string
	Ws     string
}

type RecentBlockHash struct {
	Nodes []Node
}

type BlockHandler struct {
	Nodes []Node
}

type Market struct {
	Node     Node
	Programs []Program
}

type Program struct {
	Enable bool
	Name   string
	Id     solana.PublicKey
}

type Notify struct {
	Items map[string]struct {
		Enable bool   `json:"enable"`
		Url    string `json:"url"`
	}
}

type Sender struct {
	Enable     bool
	Name       string
	Nodes      []Node
	Parameters map[string]interface{}
}

type Snipe struct {
	Tokens map[solana.PublicKey]bool
	Pools  map[solana.PublicKey]bool
	Trade  struct {
		InputToken  solana.PublicKey
		InputAmount uint64
	}
	TransactionSender string
}

type Wallet struct {
	Key string
}

type Config struct {
	Mainnet         bool
	RecentBlockHash RecentBlockHash
	BlockHandler    BlockHandler
	Node            Node
	Market          Market
	Senders         []Sender
	Notify          Notify
	Wallet          Wallet
	Snipe           Snipe
}
