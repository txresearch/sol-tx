package transaction

import (
	"encoding/json"
	"errors"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	Hash         solana.Signature
	Accounts     []rpc.ParsedMessageAccount
	TokenOwner   map[solana.PublicKey]solana.PublicKey
	TokenMint    map[solana.PublicKey]solana.PublicKey
	PreBalance   map[solana.PublicKey]decimal.Decimal
	PostBalance  map[solana.PublicKey]decimal.Decimal
	Instructions []*InstructionNode
	Slot         uint64
	BlockTime    time.Time
	ErrorMessage []byte
	Seq          int
}

type InstructionNode struct {
	Seq         int
	Instruction *rpc.ParsedInstruction
	Children    []*InstructionNode
}

func NewTransaction() *Transaction {
	return &Transaction{
		TokenOwner:  make(map[solana.PublicKey]solana.PublicKey),
		TokenMint:   make(map[solana.PublicKey]solana.PublicKey),
		PreBalance:  make(map[solana.PublicKey]decimal.Decimal),
		PostBalance: make(map[solana.PublicKey]decimal.Decimal),
	}
}

func (t *Transaction) Parse(tx *rpc.ParsedTransactionWithMeta) error {
	if tx.Meta == nil || tx.Transaction == nil {
		return errors.New("transaction meta or data is missing")
	}
	meta := tx.Meta
	transaction := tx.Transaction
	t.Hash = transaction.Signatures[0]
	t.Slot = tx.Slot
	t.BlockTime = tx.BlockTime.Time()
	if meta.Err != nil {
		// if failed, ignore this transaction
		errJson, _ := json.Marshal(meta.Err)
		t.ErrorMessage = errJson
		return nil
	}
	message := transaction.Message
	instructions := message.Instructions
	if len(instructions) == 0 {
		return nil
	}
	if instructions[0].ProgramId == solana.VoteProgramID {
		return nil
	}
	// account infos
	t.Accounts = message.AccountKeys
	for _, item := range meta.PostTokenBalances {
		account := t.Accounts[item.AccountIndex]
		t.TokenOwner[account.PublicKey] = *item.Owner
		t.TokenMint[account.PublicKey] = item.Mint
		t.PostBalance[account.PublicKey], _ = decimal.NewFromString(item.UiTokenAmount.Amount)
	}
	for _, item := range meta.PreTokenBalances {
		account := t.Accounts[item.AccountIndex]
		t.PreBalance[account.PublicKey], _ = decimal.NewFromString(item.UiTokenAmount.Amount)
	}
	for index, instruction := range instructions {
		instruction.StackHeight = 1
		current := &InstructionNode{
			Seq:         index + 1,
			Instruction: instruction,
			Children:    nil,
		}
		t.Instructions = append(t.Instructions, current)
	}
	innerInstructions := meta.InnerInstructions
	for _, innerInstruction := range innerInstructions {
		parent := t.Instructions[innerInstruction.Index]
		t.parseInnerInstructions(innerInstruction.Instructions, parent)
	}
	return nil
}

func (t *Transaction) split(ins []*rpc.ParsedInstruction) []int {
	currentHeight := ins[0].StackHeight
	split := make([]int, 0)
	for index, item := range ins {
		if item.StackHeight == currentHeight {
			split = append(split, index)
		}
	}
	return split
}

func (t *Transaction) parseInnerInstructions(ins []*rpc.ParsedInstruction, parent *InstructionNode) {
	if len(ins) == 0 {
		return
	}
	// ins split by stack height
	splited := t.split(ins)
	splited = append(splited, len(ins))
	for i := 0; i < len(splited)-1; i++ {
		index1 := splited[i]
		index2 := splited[i+1]
		current := &InstructionNode{
			Seq:         i + 1,
			Instruction: ins[index1],
			Children:    nil,
		}
		parent.Children = append(parent.Children, current)
		t.parseInnerInstructions(ins[index1+1:index2], current)
	}
}
