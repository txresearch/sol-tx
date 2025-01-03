package transaction

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/sol-tx/transaction/types"
)

type Instruction struct {
	Seq         int
	Instruction *rpc.ParsedInstruction
	Event       interface{}
	Children    []*Instruction
}

func (in *Instruction) GetEvent() interface{} {
	return in.Event
}
func (in *Instruction) GetChildren(i int) types.Instruction {
	if i >= len(in.Children) {
		return nil
	}
	return in.Children[i]
}
func (in *Instruction) GetInstruction() *rpc.ParsedInstruction {
	return in.Instruction
}

func (in *Instruction) split(subIns []*rpc.ParsedInstruction) []int {
	currentHeight := subIns[0].StackHeight
	split := make([]int, 0)
	for index, item := range subIns {
		if item.StackHeight == currentHeight {
			split = append(split, index)
		}
	}
	return split
}

func (in *Instruction) parseInstructions(subIns []*rpc.ParsedInstruction) {
	if len(subIns) == 0 {
		return
	}
	// ins split by stack height
	split := in.split(subIns)
	split = append(split, len(subIns))
	for i := 0; i < len(split)-1; i++ {
		index1 := split[i]
		index2 := split[i+1]
		current := &Instruction{
			Seq:         i + 1,
			Instruction: subIns[index1],
			Children:    nil,
		}
		in.Children = append(in.Children, current)
		current.parseInstructions(subIns[index1+1 : index2])
	}
}

func (in *Instruction) instructionActions(parsers map[solana.PublicKey]types.ActionParser, meta *types.Meta) {
	for _, child := range in.Children {
		child.instructionActions(parsers, meta)
	}
	parser, ok := parsers[in.Instruction.ProgramId]
	if !ok {
		return
	}
	event := parser.Parse(in, meta)
	if event != nil {
		in.Event = event
	}
}
