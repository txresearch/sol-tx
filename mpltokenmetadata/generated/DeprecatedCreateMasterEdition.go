// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// DeprecatedCreateMasterEdition is the `DeprecatedCreateMasterEdition` instruction.
type DeprecatedCreateMasterEdition struct {

	// [0] = [WRITE] edition
	// ··········· Unallocated edition V1 account with address as pda of ['metadata', program id, mint, 'edition']
	//
	// [1] = [WRITE] mint
	// ··········· Metadata mint
	//
	// [2] = [WRITE] printingMint
	// ··········· Printing mint - A mint you control that can mint tokens that can be exchanged for limited editions of your master edition via the MintNewEditionFromMasterEditionViaToken endpoint
	//
	// [3] = [WRITE] oneTimePrintingAuthorizationMint
	// ··········· One time authorization printing mint - A mint you control that prints tokens that gives the bearer permission to mint any number of tokens from the printing mint one time via an endpoint with the token-metadata program for your metadata. Also burns the token.
	//
	// [4] = [SIGNER] updateAuthority
	// ··········· Current Update authority key
	//
	// [5] = [SIGNER] printingMintAuthority
	// ··········· Printing mint authority - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY.
	//
	// [6] = [SIGNER] mintAuthority
	// ··········· Mint authority on the metadata's mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
	//
	// [7] = [] metadata
	// ··········· Metadata account
	//
	// [8] = [SIGNER] payer
	// ··········· payer
	//
	// [9] = [] tokenProgram
	// ··········· Token program
	//
	// [10] = [] systemProgram
	// ··········· System program
	//
	// [11] = [] rent
	// ··········· Rent info
	//
	// [12] = [SIGNER] oneTimePrintingAuthorizationMintAuthority
	// ··········· One time authorization printing mint authority - must be provided if using max supply. THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY.
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDeprecatedCreateMasterEditionInstructionBuilder creates a new `DeprecatedCreateMasterEdition` instruction builder.
func NewDeprecatedCreateMasterEditionInstructionBuilder() *DeprecatedCreateMasterEdition {
	nd := &DeprecatedCreateMasterEdition{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	return nd
}

// SetEditionAccount sets the "edition" account.
// Unallocated edition V1 account with address as pda of ['metadata', program id, mint, 'edition']
func (inst *DeprecatedCreateMasterEdition) SetEditionAccount(edition ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(edition).WRITE()
	return inst
}

// GetEditionAccount gets the "edition" account.
// Unallocated edition V1 account with address as pda of ['metadata', program id, mint, 'edition']
func (inst *DeprecatedCreateMasterEdition) GetEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAccount sets the "mint" account.
// Metadata mint
func (inst *DeprecatedCreateMasterEdition) SetMintAccount(mint ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
// Metadata mint
func (inst *DeprecatedCreateMasterEdition) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPrintingMintAccount sets the "printingMint" account.
// Printing mint - A mint you control that can mint tokens that can be exchanged for limited editions of your master edition via the MintNewEditionFromMasterEditionViaToken endpoint
func (inst *DeprecatedCreateMasterEdition) SetPrintingMintAccount(printingMint ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(printingMint).WRITE()
	return inst
}

// GetPrintingMintAccount gets the "printingMint" account.
// Printing mint - A mint you control that can mint tokens that can be exchanged for limited editions of your master edition via the MintNewEditionFromMasterEditionViaToken endpoint
func (inst *DeprecatedCreateMasterEdition) GetPrintingMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOneTimePrintingAuthorizationMintAccount sets the "oneTimePrintingAuthorizationMint" account.
// One time authorization printing mint - A mint you control that prints tokens that gives the bearer permission to mint any number of tokens from the printing mint one time via an endpoint with the token-metadata program for your metadata. Also burns the token.
func (inst *DeprecatedCreateMasterEdition) SetOneTimePrintingAuthorizationMintAccount(oneTimePrintingAuthorizationMint ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(oneTimePrintingAuthorizationMint).WRITE()
	return inst
}

// GetOneTimePrintingAuthorizationMintAccount gets the "oneTimePrintingAuthorizationMint" account.
// One time authorization printing mint - A mint you control that prints tokens that gives the bearer permission to mint any number of tokens from the printing mint one time via an endpoint with the token-metadata program for your metadata. Also burns the token.
func (inst *DeprecatedCreateMasterEdition) GetOneTimePrintingAuthorizationMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
// Current Update authority key
func (inst *DeprecatedCreateMasterEdition) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(updateAuthority).SIGNER()
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
// Current Update authority key
func (inst *DeprecatedCreateMasterEdition) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetPrintingMintAuthorityAccount sets the "printingMintAuthority" account.
// Printing mint authority - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY.
func (inst *DeprecatedCreateMasterEdition) SetPrintingMintAuthorityAccount(printingMintAuthority ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(printingMintAuthority).SIGNER()
	return inst
}

// GetPrintingMintAuthorityAccount gets the "printingMintAuthority" account.
// Printing mint authority - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY.
func (inst *DeprecatedCreateMasterEdition) GetPrintingMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMintAuthorityAccount sets the "mintAuthority" account.
// Mint authority on the metadata's mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
func (inst *DeprecatedCreateMasterEdition) SetMintAuthorityAccount(mintAuthority ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(mintAuthority).SIGNER()
	return inst
}

// GetMintAuthorityAccount gets the "mintAuthority" account.
// Mint authority on the metadata's mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
func (inst *DeprecatedCreateMasterEdition) GetMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetMetadataAccount sets the "metadata" account.
// Metadata account
func (inst *DeprecatedCreateMasterEdition) SetMetadataAccount(metadata ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata account
func (inst *DeprecatedCreateMasterEdition) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetPayerAccount sets the "payer" account.
// payer
func (inst *DeprecatedCreateMasterEdition) SetPayerAccount(payer ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// payer
func (inst *DeprecatedCreateMasterEdition) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *DeprecatedCreateMasterEdition) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *DeprecatedCreateMasterEdition) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// System program
func (inst *DeprecatedCreateMasterEdition) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// System program
func (inst *DeprecatedCreateMasterEdition) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetRentAccount sets the "rent" account.
// Rent info
func (inst *DeprecatedCreateMasterEdition) SetRentAccount(rent ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
// Rent info
func (inst *DeprecatedCreateMasterEdition) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetOneTimePrintingAuthorizationMintAuthorityAccount sets the "oneTimePrintingAuthorizationMintAuthority" account.
// One time authorization printing mint authority - must be provided if using max supply. THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY.
func (inst *DeprecatedCreateMasterEdition) SetOneTimePrintingAuthorizationMintAuthorityAccount(oneTimePrintingAuthorizationMintAuthority ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(oneTimePrintingAuthorizationMintAuthority).SIGNER()
	return inst
}

// GetOneTimePrintingAuthorizationMintAuthorityAccount gets the "oneTimePrintingAuthorizationMintAuthority" account.
// One time authorization printing mint authority - must be provided if using max supply. THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY.
func (inst *DeprecatedCreateMasterEdition) GetOneTimePrintingAuthorizationMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

func (inst DeprecatedCreateMasterEdition) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DeprecatedCreateMasterEdition,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DeprecatedCreateMasterEdition) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DeprecatedCreateMasterEdition) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Edition is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PrintingMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.OneTimePrintingAuthorizationMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.PrintingMintAuthority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.MintAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.OneTimePrintingAuthorizationMintAuthority is not set")
		}
	}
	return nil
}

func (inst *DeprecatedCreateMasterEdition) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DeprecatedCreateMasterEdition")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                                  edition", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                                     mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                             printingMint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         oneTimePrintingAuthorizationMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                          updateAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                    printingMintAuthority", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                            mintAuthority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                                 metadata", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                                    payer", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                             tokenProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                            systemProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                                     rent", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("oneTimePrintingAuthorizationMintAuthority", inst.AccountMetaSlice.Get(12)))
					})
				})
		})
}

func (obj DeprecatedCreateMasterEdition) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *DeprecatedCreateMasterEdition) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewDeprecatedCreateMasterEditionInstruction declares a new DeprecatedCreateMasterEdition instruction with the provided parameters and accounts.
func NewDeprecatedCreateMasterEditionInstruction(
	// Accounts:
	edition ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	printingMint ag_solanago.PublicKey,
	oneTimePrintingAuthorizationMint ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	printingMintAuthority ag_solanago.PublicKey,
	mintAuthority ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	oneTimePrintingAuthorizationMintAuthority ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	return NewDeprecatedCreateMasterEditionInstructionBuilder().
		SetEditionAccount(edition).
		SetMintAccount(mint).
		SetPrintingMintAccount(printingMint).
		SetOneTimePrintingAuthorizationMintAccount(oneTimePrintingAuthorizationMint).
		SetUpdateAuthorityAccount(updateAuthority).
		SetPrintingMintAuthorityAccount(printingMintAuthority).
		SetMintAuthorityAccount(mintAuthority).
		SetMetadataAccount(metadata).
		SetPayerAccount(payer).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetOneTimePrintingAuthorizationMintAuthorityAccount(oneTimePrintingAuthorizationMintAuthority)
}