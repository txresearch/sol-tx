// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package generated

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetRewardEmissionsSuperAuthority is the `setRewardEmissionsSuperAuthority` instruction.
type SetRewardEmissionsSuperAuthority struct {

	// [0] = [WRITE] whirlpoolsConfig
	//
	// [1] = [SIGNER] rewardEmissionsSuperAuthority
	//
	// [2] = [] newRewardEmissionsSuperAuthority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetRewardEmissionsSuperAuthorityInstructionBuilder creates a new `SetRewardEmissionsSuperAuthority` instruction builder.
func NewSetRewardEmissionsSuperAuthorityInstructionBuilder() *SetRewardEmissionsSuperAuthority {
	nd := &SetRewardEmissionsSuperAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetWhirlpoolsConfigAccount sets the "whirlpoolsConfig" account.
func (inst *SetRewardEmissionsSuperAuthority) SetWhirlpoolsConfigAccount(whirlpoolsConfig ag_solanago.PublicKey) *SetRewardEmissionsSuperAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(whirlpoolsConfig).WRITE()
	return inst
}

// GetWhirlpoolsConfigAccount gets the "whirlpoolsConfig" account.
func (inst *SetRewardEmissionsSuperAuthority) GetWhirlpoolsConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetRewardEmissionsSuperAuthorityAccount sets the "rewardEmissionsSuperAuthority" account.
func (inst *SetRewardEmissionsSuperAuthority) SetRewardEmissionsSuperAuthorityAccount(rewardEmissionsSuperAuthority ag_solanago.PublicKey) *SetRewardEmissionsSuperAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(rewardEmissionsSuperAuthority).SIGNER()
	return inst
}

// GetRewardEmissionsSuperAuthorityAccount gets the "rewardEmissionsSuperAuthority" account.
func (inst *SetRewardEmissionsSuperAuthority) GetRewardEmissionsSuperAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetNewRewardEmissionsSuperAuthorityAccount sets the "newRewardEmissionsSuperAuthority" account.
func (inst *SetRewardEmissionsSuperAuthority) SetNewRewardEmissionsSuperAuthorityAccount(newRewardEmissionsSuperAuthority ag_solanago.PublicKey) *SetRewardEmissionsSuperAuthority {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(newRewardEmissionsSuperAuthority)
	return inst
}

// GetNewRewardEmissionsSuperAuthorityAccount gets the "newRewardEmissionsSuperAuthority" account.
func (inst *SetRewardEmissionsSuperAuthority) GetNewRewardEmissionsSuperAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst SetRewardEmissionsSuperAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetRewardEmissionsSuperAuthority,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetRewardEmissionsSuperAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetRewardEmissionsSuperAuthority) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.WhirlpoolsConfig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.RewardEmissionsSuperAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.NewRewardEmissionsSuperAuthority is not set")
		}
	}
	return nil
}

func (inst *SetRewardEmissionsSuperAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetRewardEmissionsSuperAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                whirlpoolsConfig", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   rewardEmissionsSuperAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("newRewardEmissionsSuperAuthority", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj SetRewardEmissionsSuperAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *SetRewardEmissionsSuperAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewSetRewardEmissionsSuperAuthorityInstruction declares a new SetRewardEmissionsSuperAuthority instruction with the provided parameters and accounts.
func NewSetRewardEmissionsSuperAuthorityInstruction(
	// Accounts:
	whirlpoolsConfig ag_solanago.PublicKey,
	rewardEmissionsSuperAuthority ag_solanago.PublicKey,
	newRewardEmissionsSuperAuthority ag_solanago.PublicKey) *SetRewardEmissionsSuperAuthority {
	return NewSetRewardEmissionsSuperAuthorityInstructionBuilder().
		SetWhirlpoolsConfigAccount(whirlpoolsConfig).
		SetRewardEmissionsSuperAuthorityAccount(rewardEmissionsSuperAuthority).
		SetNewRewardEmissionsSuperAuthorityAccount(newRewardEmissionsSuperAuthority)
}
