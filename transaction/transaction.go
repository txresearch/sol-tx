package transaction

import (
	"encoding/json"
	"errors"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/shopspring/decimal"
	"github.com/sol-tx/transaction/types"
	"time"
)

type Transaction struct {
	Hash         solana.Signature
	Instructions []*Instruction
	Meta         types.Meta
	Slot         uint64
	BlockTime    time.Time
	Seq          int
}

func New() *Transaction {
	return &Transaction{
		Meta: types.Meta{
			TokenOwner:  make(map[solana.PublicKey]solana.PublicKey),
			TokenMint:   make(map[solana.PublicKey]solana.PublicKey),
			PreBalance:  make(map[solana.PublicKey]decimal.Decimal),
			PostBalance: make(map[solana.PublicKey]decimal.Decimal),
		},
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
		t.Meta.ErrorMessage = errJson
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
	t.Meta.Accounts = message.AccountKeys
	for _, item := range meta.PostTokenBalances {
		account := t.Meta.Accounts[item.AccountIndex]
		t.Meta.TokenOwner[account.PublicKey] = *item.Owner
		t.Meta.TokenMint[account.PublicKey] = item.Mint
		t.Meta.PostBalance[account.PublicKey], _ = decimal.NewFromString(item.UiTokenAmount.Amount)
	}
	for _, item := range meta.PreTokenBalances {
		account := t.Meta.Accounts[item.AccountIndex]
		t.Meta.PreBalance[account.PublicKey], _ = decimal.NewFromString(item.UiTokenAmount.Amount)
	}
	for index, instruction := range instructions {
		instruction.StackHeight = 1
		current := &Instruction{
			Seq:         index + 1,
			Instruction: instruction,
			Children:    nil,
		}
		t.Instructions = append(t.Instructions, current)
	}
	innerInstructions := meta.InnerInstructions
	for _, innerInstruction := range innerInstructions {
		parent := t.Instructions[innerInstruction.Index]
		parent.parseInstructions(innerInstruction.Instructions)
	}
	return nil
}

func (t *Transaction) ParseActions(parsers map[solana.PublicKey]types.ActionParser) error {
	for _, instruction := range t.Instructions {
		t.instructionActions(instruction, parsers)
	}
	return nil
}

func (t *Transaction) instructionActions(in *Instruction, parsers map[solana.PublicKey]types.ActionParser) {
	for _, child := range in.Children {
		t.instructionActions(child, parsers)
	}
	parser, ok := parsers[in.Instruction.ProgramId]
	if !ok {
		return
	}
	event := parser.Parse(in, &t.Meta)
	if event != nil {
		in.Event = event
	}
}
