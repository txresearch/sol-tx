// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// TransferOutOfEscrow is the `TransferOutOfEscrow` instruction.
type TransferOutOfEscrow struct {
	TransferOutOfEscrowArgs *TransferOutOfEscrowArgs

	// [0] = [] escrow
	// ··········· Escrow account
	//
	// [1] = [WRITE] metadata
	// ··········· Metadata account
	//
	// [2] = [WRITE, SIGNER] payer
	// ··········· Wallet paying for the transaction and new account
	//
	// [3] = [] attributeMint
	// ··········· Mint account for the new attribute
	//
	// [4] = [WRITE] attributeSrc
	// ··········· Token account source for the new attribute
	//
	// [5] = [WRITE] attributeDst
	// ··········· Token account, owned by TM, destination for the new attribute
	//
	// [6] = [] escrowMint
	// ··········· Mint account that the escrow is attached
	//
	// [7] = [] escrowAccount
	// ··········· Token account that holds the token the escrow is attached to
	//
	// [8] = [] systemProgram
	// ··········· System program
	//
	// [9] = [] ataProgram
	// ··········· Associated Token program
	//
	// [10] = [] tokenProgram
	// ··········· Token program
	//
	// [11] = [] sysvarInstructions
	// ··········· Instructions sysvar account
	//
	// [12] = [SIGNER] authority
	// ··········· Authority/creator of the escrow account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewTransferOutOfEscrowInstructionBuilder creates a new `TransferOutOfEscrow` instruction builder.
func NewTransferOutOfEscrowInstructionBuilder() *TransferOutOfEscrow {
	nd := &TransferOutOfEscrow{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	return nd
}

// SetTransferOutOfEscrowArgs sets the "transferOutOfEscrowArgs" parameter.
func (inst *TransferOutOfEscrow) SetTransferOutOfEscrowArgs(transferOutOfEscrowArgs TransferOutOfEscrowArgs) *TransferOutOfEscrow {
	inst.TransferOutOfEscrowArgs = &transferOutOfEscrowArgs
	return inst
}

// SetEscrowAccount sets the "escrow" account.
// Escrow account
func (inst *TransferOutOfEscrow) SetEscrowAccount(escrow ag_solanago.PublicKey) *TransferOutOfEscrow {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(escrow)
	return inst
}

// GetEscrowAccount gets the "escrow" account.
// Escrow account
func (inst *TransferOutOfEscrow) GetEscrowAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMetadataAccount sets the "metadata" account.
// Metadata account
func (inst *TransferOutOfEscrow) SetMetadataAccount(metadata ag_solanago.PublicKey) *TransferOutOfEscrow {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata account
func (inst *TransferOutOfEscrow) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
// Wallet paying for the transaction and new account
func (inst *TransferOutOfEscrow) SetPayerAccount(payer ag_solanago.PublicKey) *TransferOutOfEscrow {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Wallet paying for the transaction and new account
func (inst *TransferOutOfEscrow) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAttributeMintAccount sets the "attributeMint" account.
// Mint account for the new attribute
func (inst *TransferOutOfEscrow) SetAttributeMintAccount(attributeMint ag_solanago.PublicKey) *TransferOutOfEscrow {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(attributeMint)
	return inst
}

// GetAttributeMintAccount gets the "attributeMint" account.
// Mint account for the new attribute
func (inst *TransferOutOfEscrow) GetAttributeMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAttributeSrcAccount sets the "attributeSrc" account.
// Token account source for the new attribute
func (inst *TransferOutOfEscrow) SetAttributeSrcAccount(attributeSrc ag_solanago.PublicKey) *TransferOutOfEscrow {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(attributeSrc).WRITE()
	return inst
}

// GetAttributeSrcAccount gets the "attributeSrc" account.
// Token account source for the new attribute
func (inst *TransferOutOfEscrow) GetAttributeSrcAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAttributeDstAccount sets the "attributeDst" account.
// Token account, owned by TM, destination for the new attribute
func (inst *TransferOutOfEscrow) SetAttributeDstAccount(attributeDst ag_solanago.PublicKey) *TransferOutOfEscrow {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(attributeDst).WRITE()
	return inst
}

// GetAttributeDstAccount gets the "attributeDst" account.
// Token account, owned by TM, destination for the new attribute
func (inst *TransferOutOfEscrow) GetAttributeDstAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetEscrowMintAccount sets the "escrowMint" account.
// Mint account that the escrow is attached
func (inst *TransferOutOfEscrow) SetEscrowMintAccount(escrowMint ag_solanago.PublicKey) *TransferOutOfEscrow {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(escrowMint)
	return inst
}

// GetEscrowMintAccount gets the "escrowMint" account.
// Mint account that the escrow is attached
func (inst *TransferOutOfEscrow) GetEscrowMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetEscrowAccountAccount sets the "escrowAccount" account.
// Token account that holds the token the escrow is attached to
func (inst *TransferOutOfEscrow) SetEscrowAccountAccount(escrowAccount ag_solanago.PublicKey) *TransferOutOfEscrow {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(escrowAccount)
	return inst
}

// GetEscrowAccountAccount gets the "escrowAccount" account.
// Token account that holds the token the escrow is attached to
func (inst *TransferOutOfEscrow) GetEscrowAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// System program
func (inst *TransferOutOfEscrow) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *TransferOutOfEscrow {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// System program
func (inst *TransferOutOfEscrow) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetAtaProgramAccount sets the "ataProgram" account.
// Associated Token program
func (inst *TransferOutOfEscrow) SetAtaProgramAccount(ataProgram ag_solanago.PublicKey) *TransferOutOfEscrow {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(ataProgram)
	return inst
}

// GetAtaProgramAccount gets the "ataProgram" account.
// Associated Token program
func (inst *TransferOutOfEscrow) GetAtaProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *TransferOutOfEscrow) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *TransferOutOfEscrow {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *TransferOutOfEscrow) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSysvarInstructionsAccount sets the "sysvarInstructions" account.
// Instructions sysvar account
func (inst *TransferOutOfEscrow) SetSysvarInstructionsAccount(sysvarInstructions ag_solanago.PublicKey) *TransferOutOfEscrow {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(sysvarInstructions)
	return inst
}

// GetSysvarInstructionsAccount gets the "sysvarInstructions" account.
// Instructions sysvar account
func (inst *TransferOutOfEscrow) GetSysvarInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetAuthorityAccount sets the "authority" account.
// Authority/creator of the escrow account
func (inst *TransferOutOfEscrow) SetAuthorityAccount(authority ag_solanago.PublicKey) *TransferOutOfEscrow {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// Authority/creator of the escrow account
func (inst *TransferOutOfEscrow) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

func (inst TransferOutOfEscrow) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_TransferOutOfEscrow,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst TransferOutOfEscrow) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *TransferOutOfEscrow) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.TransferOutOfEscrowArgs == nil {
			return errors.New("TransferOutOfEscrowArgs parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Escrow is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.AttributeMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.AttributeSrc is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AttributeDst is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.EscrowMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.EscrowAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.AtaProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SysvarInstructions is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *TransferOutOfEscrow) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("TransferOutOfEscrow")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("TransferOutOfEscrowArgs", *inst.TransferOutOfEscrowArgs))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("            escrow", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          metadata", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("     attributeMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      attributeSrc", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("      attributeDst", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("        escrowMint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("            escrow", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("     systemProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("        ataProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("      tokenProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("sysvarInstructions", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("         authority", inst.AccountMetaSlice.Get(12)))
					})
				})
		})
}

func (obj TransferOutOfEscrow) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TransferOutOfEscrowArgs` param:
	err = encoder.Encode(obj.TransferOutOfEscrowArgs)
	if err != nil {
		return err
	}
	return nil
}
func (obj *TransferOutOfEscrow) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TransferOutOfEscrowArgs`:
	err = decoder.Decode(&obj.TransferOutOfEscrowArgs)
	if err != nil {
		return err
	}
	return nil
}

// NewTransferOutOfEscrowInstruction declares a new TransferOutOfEscrow instruction with the provided parameters and accounts.
func NewTransferOutOfEscrowInstruction(
	// Parameters:
	transferOutOfEscrowArgs TransferOutOfEscrowArgs,
	// Accounts:
	escrow ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	attributeMint ag_solanago.PublicKey,
	attributeSrc ag_solanago.PublicKey,
	attributeDst ag_solanago.PublicKey,
	escrowMint ag_solanago.PublicKey,
	escrowAccount ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	ataProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	sysvarInstructions ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *TransferOutOfEscrow {
	return NewTransferOutOfEscrowInstructionBuilder().
		SetTransferOutOfEscrowArgs(transferOutOfEscrowArgs).
		SetEscrowAccount(escrow).
		SetMetadataAccount(metadata).
		SetPayerAccount(payer).
		SetAttributeMintAccount(attributeMint).
		SetAttributeSrcAccount(attributeSrc).
		SetAttributeDstAccount(attributeDst).
		SetEscrowMintAccount(escrowMint).
		SetEscrowAccountAccount(escrowAccount).
		SetSystemProgramAccount(systemProgram).
		SetAtaProgramAccount(ataProgram).
		SetTokenProgramAccount(tokenProgram).
		SetSysvarInstructionsAccount(sysvarInstructions).
		SetAuthorityAccount(authority)
}