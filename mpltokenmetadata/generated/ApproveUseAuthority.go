// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ApproveUseAuthority is the `ApproveUseAuthority` instruction.
type ApproveUseAuthority struct {
	ApproveUseAuthorityArgs *ApproveUseAuthorityArgs

	// [0] = [WRITE] useAuthorityRecord
	// ··········· Use Authority Record PDA
	//
	// [1] = [WRITE, SIGNER] owner
	// ··········· Owner
	//
	// [2] = [WRITE, SIGNER] payer
	// ··········· Payer
	//
	// [3] = [] user
	// ··········· A Use Authority
	//
	// [4] = [WRITE] ownerTokenAccount
	// ··········· Owned Token Account Of Mint
	//
	// [5] = [] metadata
	// ··········· Metadata account
	//
	// [6] = [] mint
	// ··········· Mint of Metadata
	//
	// [7] = [] burner
	// ··········· Program As Signer (Burner)
	//
	// [8] = [] tokenProgram
	// ··········· Token program
	//
	// [9] = [] systemProgram
	// ··········· System program
	//
	// [10] = [] rent
	// ··········· Rent info
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewApproveUseAuthorityInstructionBuilder creates a new `ApproveUseAuthority` instruction builder.
func NewApproveUseAuthorityInstructionBuilder() *ApproveUseAuthority {
	nd := &ApproveUseAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetApproveUseAuthorityArgs sets the "approveUseAuthorityArgs" parameter.
func (inst *ApproveUseAuthority) SetApproveUseAuthorityArgs(approveUseAuthorityArgs ApproveUseAuthorityArgs) *ApproveUseAuthority {
	inst.ApproveUseAuthorityArgs = &approveUseAuthorityArgs
	return inst
}

// SetUseAuthorityRecordAccount sets the "useAuthorityRecord" account.
// Use Authority Record PDA
func (inst *ApproveUseAuthority) SetUseAuthorityRecordAccount(useAuthorityRecord ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(useAuthorityRecord).WRITE()
	return inst
}

// GetUseAuthorityRecordAccount gets the "useAuthorityRecord" account.
// Use Authority Record PDA
func (inst *ApproveUseAuthority) GetUseAuthorityRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
// Owner
func (inst *ApproveUseAuthority) SetOwnerAccount(owner ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
// Owner
func (inst *ApproveUseAuthority) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *ApproveUseAuthority) SetPayerAccount(payer ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *ApproveUseAuthority) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserAccount sets the "user" account.
// A Use Authority
func (inst *ApproveUseAuthority) SetUserAccount(user ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(user)
	return inst
}

// GetUserAccount gets the "user" account.
// A Use Authority
func (inst *ApproveUseAuthority) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetOwnerTokenAccountAccount sets the "ownerTokenAccount" account.
// Owned Token Account Of Mint
func (inst *ApproveUseAuthority) SetOwnerTokenAccountAccount(ownerTokenAccount ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(ownerTokenAccount).WRITE()
	return inst
}

// GetOwnerTokenAccountAccount gets the "ownerTokenAccount" account.
// Owned Token Account Of Mint
func (inst *ApproveUseAuthority) GetOwnerTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMetadataAccount sets the "metadata" account.
// Metadata account
func (inst *ApproveUseAuthority) SetMetadataAccount(metadata ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata account
func (inst *ApproveUseAuthority) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMintAccount sets the "mint" account.
// Mint of Metadata
func (inst *ApproveUseAuthority) SetMintAccount(mint ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
// Mint of Metadata
func (inst *ApproveUseAuthority) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetBurnerAccount sets the "burner" account.
// Program As Signer (Burner)
func (inst *ApproveUseAuthority) SetBurnerAccount(burner ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(burner)
	return inst
}

// GetBurnerAccount gets the "burner" account.
// Program As Signer (Burner)
func (inst *ApproveUseAuthority) GetBurnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *ApproveUseAuthority) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *ApproveUseAuthority) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// System program
func (inst *ApproveUseAuthority) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// System program
func (inst *ApproveUseAuthority) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetRentAccount sets the "rent" account.
// Rent info
func (inst *ApproveUseAuthority) SetRentAccount(rent ag_solanago.PublicKey) *ApproveUseAuthority {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
// Rent info
func (inst *ApproveUseAuthority) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst ApproveUseAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ApproveUseAuthority,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ApproveUseAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ApproveUseAuthority) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ApproveUseAuthorityArgs == nil {
			return errors.New("ApproveUseAuthorityArgs parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.UseAuthorityRecord is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.OwnerTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Burner is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *ApproveUseAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ApproveUseAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ApproveUseAuthorityArgs", *inst.ApproveUseAuthorityArgs))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("useAuthorityRecord", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("              user", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        ownerToken", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          metadata", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("              mint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("            burner", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("      tokenProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("     systemProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("              rent", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj ApproveUseAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ApproveUseAuthorityArgs` param:
	err = encoder.Encode(obj.ApproveUseAuthorityArgs)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ApproveUseAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ApproveUseAuthorityArgs`:
	err = decoder.Decode(&obj.ApproveUseAuthorityArgs)
	if err != nil {
		return err
	}
	return nil
}

// NewApproveUseAuthorityInstruction declares a new ApproveUseAuthority instruction with the provided parameters and accounts.
func NewApproveUseAuthorityInstruction(
	// Parameters:
	approveUseAuthorityArgs ApproveUseAuthorityArgs,
	// Accounts:
	useAuthorityRecord ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	ownerTokenAccount ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	burner ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *ApproveUseAuthority {
	return NewApproveUseAuthorityInstructionBuilder().
		SetApproveUseAuthorityArgs(approveUseAuthorityArgs).
		SetUseAuthorityRecordAccount(useAuthorityRecord).
		SetOwnerAccount(owner).
		SetPayerAccount(payer).
		SetUserAccount(user).
		SetOwnerTokenAccountAccount(ownerTokenAccount).
		SetMetadataAccount(metadata).
		SetMintAccount(mint).
		SetBurnerAccount(burner).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}