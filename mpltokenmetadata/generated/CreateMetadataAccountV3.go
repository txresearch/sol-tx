// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateMetadataAccountV3 is the `CreateMetadataAccountV3` instruction.
type CreateMetadataAccountV3 struct {
	CreateMetadataAccountArgsV3 *CreateMetadataAccountArgsV3

	// [0] = [WRITE] metadata
	// ··········· Metadata key (pda of ['metadata', program id, mint id])
	//
	// [1] = [] mint
	// ··········· Mint of token asset
	//
	// [2] = [SIGNER] mintAuthority
	// ··········· Mint authority
	//
	// [3] = [WRITE, SIGNER] payer
	// ··········· payer
	//
	// [4] = [] updateAuthority
	// ··········· update authority info
	//
	// [5] = [] systemProgram
	// ··········· System program
	//
	// [6] = [] rent
	// ··········· Rent info
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateMetadataAccountV3InstructionBuilder creates a new `CreateMetadataAccountV3` instruction builder.
func NewCreateMetadataAccountV3InstructionBuilder() *CreateMetadataAccountV3 {
	nd := &CreateMetadataAccountV3{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetCreateMetadataAccountArgsV3 sets the "createMetadataAccountArgsV3" parameter.
func (inst *CreateMetadataAccountV3) SetCreateMetadataAccountArgsV3(createMetadataAccountArgsV3 CreateMetadataAccountArgsV3) *CreateMetadataAccountV3 {
	inst.CreateMetadataAccountArgsV3 = &createMetadataAccountArgsV3
	return inst
}

// SetMetadataAccount sets the "metadata" account.
// Metadata key (pda of ['metadata', program id, mint id])
func (inst *CreateMetadataAccountV3) SetMetadataAccount(metadata ag_solanago.PublicKey) *CreateMetadataAccountV3 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata key (pda of ['metadata', program id, mint id])
func (inst *CreateMetadataAccountV3) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAccount sets the "mint" account.
// Mint of token asset
func (inst *CreateMetadataAccountV3) SetMintAccount(mint ag_solanago.PublicKey) *CreateMetadataAccountV3 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
// Mint of token asset
func (inst *CreateMetadataAccountV3) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintAuthorityAccount sets the "mintAuthority" account.
// Mint authority
func (inst *CreateMetadataAccountV3) SetMintAuthorityAccount(mintAuthority ag_solanago.PublicKey) *CreateMetadataAccountV3 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mintAuthority).SIGNER()
	return inst
}

// GetMintAuthorityAccount gets the "mintAuthority" account.
// Mint authority
func (inst *CreateMetadataAccountV3) GetMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
// payer
func (inst *CreateMetadataAccountV3) SetPayerAccount(payer ag_solanago.PublicKey) *CreateMetadataAccountV3 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// payer
func (inst *CreateMetadataAccountV3) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
// update authority info
func (inst *CreateMetadataAccountV3) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *CreateMetadataAccountV3 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(updateAuthority)
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
// update authority info
func (inst *CreateMetadataAccountV3) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// System program
func (inst *CreateMetadataAccountV3) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateMetadataAccountV3 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// System program
func (inst *CreateMetadataAccountV3) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetRentAccount sets the "rent" account.
// Rent info
func (inst *CreateMetadataAccountV3) SetRentAccount(rent ag_solanago.PublicKey) *CreateMetadataAccountV3 {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
// Rent info
func (inst *CreateMetadataAccountV3) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst CreateMetadataAccountV3) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateMetadataAccountV3,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateMetadataAccountV3) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateMetadataAccountV3) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.CreateMetadataAccountArgsV3 == nil {
			return errors.New("CreateMetadataAccountArgsV3 parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MintAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *CreateMetadataAccountV3) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateMetadataAccountV3")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("CreateMetadataAccountArgsV3", *inst.CreateMetadataAccountArgsV3))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       metadata", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  mintAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("updateAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("  systemProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           rent", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj CreateMetadataAccountV3) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CreateMetadataAccountArgsV3` param:
	err = encoder.Encode(obj.CreateMetadataAccountArgsV3)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateMetadataAccountV3) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CreateMetadataAccountArgsV3`:
	err = decoder.Decode(&obj.CreateMetadataAccountArgsV3)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateMetadataAccountV3Instruction declares a new CreateMetadataAccountV3 instruction with the provided parameters and accounts.
func NewCreateMetadataAccountV3Instruction(
	// Parameters:
	createMetadataAccountArgsV3 CreateMetadataAccountArgsV3,
	// Accounts:
	metadata ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	mintAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *CreateMetadataAccountV3 {
	return NewCreateMetadataAccountV3InstructionBuilder().
		SetCreateMetadataAccountArgsV3(createMetadataAccountArgsV3).
		SetMetadataAccount(metadata).
		SetMintAccount(mint).
		SetMintAuthorityAccount(mintAuthority).
		SetPayerAccount(payer).
		SetUpdateAuthorityAccount(updateAuthority).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}