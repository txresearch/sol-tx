// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Verify is the `Verify` instruction.
type Verify struct {
	VerificationArgs *VerificationArgs

	// [0] = [SIGNER] authority
	// ··········· Creator to verify, collection update authority or delegate
	//
	// [1] = [] delegateRecord
	// ··········· Delegate record PDA
	//
	// [2] = [WRITE] metadata
	// ··········· Metadata account
	//
	// [3] = [] collectionMint
	// ··········· Mint of the Collection
	//
	// [4] = [WRITE] collectionMetadata
	// ··········· Metadata Account of the Collection
	//
	// [5] = [] collectionMasterEdition
	// ··········· Master Edition Account of the Collection Token
	//
	// [6] = [] systemProgram
	// ··········· System program
	//
	// [7] = [] sysvarInstructions
	// ··········· Instructions sysvar account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewVerifyInstructionBuilder creates a new `Verify` instruction builder.
func NewVerifyInstructionBuilder() *Verify {
	nd := &Verify{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetVerificationArgs sets the "verificationArgs" parameter.
func (inst *Verify) SetVerificationArgs(verificationArgs VerificationArgs) *Verify {
	inst.VerificationArgs = &verificationArgs
	return inst
}

// SetAuthorityAccount sets the "authority" account.
// Creator to verify, collection update authority or delegate
func (inst *Verify) SetAuthorityAccount(authority ag_solanago.PublicKey) *Verify {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// Creator to verify, collection update authority or delegate
func (inst *Verify) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetDelegateRecordAccount sets the "delegateRecord" account.
// Delegate record PDA
func (inst *Verify) SetDelegateRecordAccount(delegateRecord ag_solanago.PublicKey) *Verify {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(delegateRecord)
	return inst
}

// GetDelegateRecordAccount gets the "delegateRecord" account.
// Delegate record PDA
func (inst *Verify) GetDelegateRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMetadataAccount sets the "metadata" account.
// Metadata account
func (inst *Verify) SetMetadataAccount(metadata ag_solanago.PublicKey) *Verify {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata account
func (inst *Verify) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetCollectionMintAccount sets the "collectionMint" account.
// Mint of the Collection
func (inst *Verify) SetCollectionMintAccount(collectionMint ag_solanago.PublicKey) *Verify {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(collectionMint)
	return inst
}

// GetCollectionMintAccount gets the "collectionMint" account.
// Mint of the Collection
func (inst *Verify) GetCollectionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetCollectionMetadataAccount sets the "collectionMetadata" account.
// Metadata Account of the Collection
func (inst *Verify) SetCollectionMetadataAccount(collectionMetadata ag_solanago.PublicKey) *Verify {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(collectionMetadata).WRITE()
	return inst
}

// GetCollectionMetadataAccount gets the "collectionMetadata" account.
// Metadata Account of the Collection
func (inst *Verify) GetCollectionMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetCollectionMasterEditionAccount sets the "collectionMasterEdition" account.
// Master Edition Account of the Collection Token
func (inst *Verify) SetCollectionMasterEditionAccount(collectionMasterEdition ag_solanago.PublicKey) *Verify {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(collectionMasterEdition)
	return inst
}

// GetCollectionMasterEditionAccount gets the "collectionMasterEdition" account.
// Master Edition Account of the Collection Token
func (inst *Verify) GetCollectionMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// System program
func (inst *Verify) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Verify {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// System program
func (inst *Verify) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSysvarInstructionsAccount sets the "sysvarInstructions" account.
// Instructions sysvar account
func (inst *Verify) SetSysvarInstructionsAccount(sysvarInstructions ag_solanago.PublicKey) *Verify {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(sysvarInstructions)
	return inst
}

// GetSysvarInstructionsAccount gets the "sysvarInstructions" account.
// Instructions sysvar account
func (inst *Verify) GetSysvarInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst Verify) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Verify,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Verify) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Verify) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.VerificationArgs == nil {
			return errors.New("VerificationArgs parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.DelegateRecord is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.CollectionMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.CollectionMetadata is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.CollectionMasterEdition is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SysvarInstructions is not set")
		}
	}
	return nil
}

func (inst *Verify) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Verify")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("VerificationArgs", *inst.VerificationArgs))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("              authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         delegateRecord", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("               metadata", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         collectionMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     collectionMetadata", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("collectionMasterEdition", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("          systemProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("     sysvarInstructions", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj Verify) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `VerificationArgs` param:
	err = encoder.Encode(obj.VerificationArgs)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Verify) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `VerificationArgs`:
	err = decoder.Decode(&obj.VerificationArgs)
	if err != nil {
		return err
	}
	return nil
}

// NewVerifyInstruction declares a new Verify instruction with the provided parameters and accounts.
func NewVerifyInstruction(
	// Parameters:
	verificationArgs VerificationArgs,
	// Accounts:
	authority ag_solanago.PublicKey,
	delegateRecord ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	collectionMint ag_solanago.PublicKey,
	collectionMetadata ag_solanago.PublicKey,
	collectionMasterEdition ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	sysvarInstructions ag_solanago.PublicKey) *Verify {
	return NewVerifyInstructionBuilder().
		SetVerificationArgs(verificationArgs).
		SetAuthorityAccount(authority).
		SetDelegateRecordAccount(delegateRecord).
		SetMetadataAccount(metadata).
		SetCollectionMintAccount(collectionMint).
		SetCollectionMetadataAccount(collectionMetadata).
		SetCollectionMasterEditionAccount(collectionMasterEdition).
		SetSystemProgramAccount(systemProgram).
		SetSysvarInstructionsAccount(sysvarInstructions)
}