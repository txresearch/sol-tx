// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package raydium_amm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Deposit is the `deposit` instruction.
type Deposit struct {
	MaxCoinAmount *uint64
	MaxPcAmount   *uint64
	BaseSide      *uint64

	// [0] = [] tokenProgram
	//
	// [1] = [WRITE] amm
	//
	// [2] = [] ammAuthority
	//
	// [3] = [] ammOpenOrders
	//
	// [4] = [WRITE] ammTargetOrders
	//
	// [5] = [WRITE] lpMintAddress
	//
	// [6] = [WRITE] poolCoinTokenAccount
	//
	// [7] = [WRITE] poolPcTokenAccount
	//
	// [8] = [] serumMarket
	//
	// [9] = [WRITE] userCoinTokenAccount
	//
	// [10] = [WRITE] userPcTokenAccount
	//
	// [11] = [WRITE] userLpTokenAccount
	//
	// [12] = [SIGNER] userOwner
	//
	// [13] = [] serumEventQueue
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDepositInstructionBuilder creates a new `Deposit` instruction builder.
func NewDepositInstructionBuilder() *Deposit {
	nd := &Deposit{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	return nd
}

// SetMaxCoinAmount sets the "maxCoinAmount" parameter.
func (inst *Deposit) SetMaxCoinAmount(maxCoinAmount uint64) *Deposit {
	inst.MaxCoinAmount = &maxCoinAmount
	return inst
}

// SetMaxPcAmount sets the "maxPcAmount" parameter.
func (inst *Deposit) SetMaxPcAmount(maxPcAmount uint64) *Deposit {
	inst.MaxPcAmount = &maxPcAmount
	return inst
}

// SetBaseSide sets the "baseSide" parameter.
func (inst *Deposit) SetBaseSide(baseSide uint64) *Deposit {
	inst.BaseSide = &baseSide
	return inst
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *Deposit) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *Deposit) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAmmAccount sets the "amm" account.
func (inst *Deposit) SetAmmAccount(amm ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(amm).WRITE()
	return inst
}

// GetAmmAccount gets the "amm" account.
func (inst *Deposit) GetAmmAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAmmAuthorityAccount sets the "ammAuthority" account.
func (inst *Deposit) SetAmmAuthorityAccount(ammAuthority ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(ammAuthority)
	return inst
}

// GetAmmAuthorityAccount gets the "ammAuthority" account.
func (inst *Deposit) GetAmmAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAmmOpenOrdersAccount sets the "ammOpenOrders" account.
func (inst *Deposit) SetAmmOpenOrdersAccount(ammOpenOrders ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(ammOpenOrders)
	return inst
}

// GetAmmOpenOrdersAccount gets the "ammOpenOrders" account.
func (inst *Deposit) GetAmmOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAmmTargetOrdersAccount sets the "ammTargetOrders" account.
func (inst *Deposit) SetAmmTargetOrdersAccount(ammTargetOrders ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(ammTargetOrders).WRITE()
	return inst
}

// GetAmmTargetOrdersAccount gets the "ammTargetOrders" account.
func (inst *Deposit) GetAmmTargetOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetLpMintAddressAccount sets the "lpMintAddress" account.
func (inst *Deposit) SetLpMintAddressAccount(lpMintAddress ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(lpMintAddress).WRITE()
	return inst
}

// GetLpMintAddressAccount gets the "lpMintAddress" account.
func (inst *Deposit) GetLpMintAddressAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPoolCoinTokenAccountAccount sets the "poolCoinTokenAccount" account.
func (inst *Deposit) SetPoolCoinTokenAccountAccount(poolCoinTokenAccount ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(poolCoinTokenAccount).WRITE()
	return inst
}

// GetPoolCoinTokenAccountAccount gets the "poolCoinTokenAccount" account.
func (inst *Deposit) GetPoolCoinTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetPoolPcTokenAccountAccount sets the "poolPcTokenAccount" account.
func (inst *Deposit) SetPoolPcTokenAccountAccount(poolPcTokenAccount ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(poolPcTokenAccount).WRITE()
	return inst
}

// GetPoolPcTokenAccountAccount gets the "poolPcTokenAccount" account.
func (inst *Deposit) GetPoolPcTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSerumMarketAccount sets the "serumMarket" account.
func (inst *Deposit) SetSerumMarketAccount(serumMarket ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(serumMarket)
	return inst
}

// GetSerumMarketAccount gets the "serumMarket" account.
func (inst *Deposit) GetSerumMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetUserCoinTokenAccountAccount sets the "userCoinTokenAccount" account.
func (inst *Deposit) SetUserCoinTokenAccountAccount(userCoinTokenAccount ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(userCoinTokenAccount).WRITE()
	return inst
}

// GetUserCoinTokenAccountAccount gets the "userCoinTokenAccount" account.
func (inst *Deposit) GetUserCoinTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetUserPcTokenAccountAccount sets the "userPcTokenAccount" account.
func (inst *Deposit) SetUserPcTokenAccountAccount(userPcTokenAccount ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(userPcTokenAccount).WRITE()
	return inst
}

// GetUserPcTokenAccountAccount gets the "userPcTokenAccount" account.
func (inst *Deposit) GetUserPcTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetUserLpTokenAccountAccount sets the "userLpTokenAccount" account.
func (inst *Deposit) SetUserLpTokenAccountAccount(userLpTokenAccount ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(userLpTokenAccount).WRITE()
	return inst
}

// GetUserLpTokenAccountAccount gets the "userLpTokenAccount" account.
func (inst *Deposit) GetUserLpTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetUserOwnerAccount sets the "userOwner" account.
func (inst *Deposit) SetUserOwnerAccount(userOwner ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(userOwner).SIGNER()
	return inst
}

// GetUserOwnerAccount gets the "userOwner" account.
func (inst *Deposit) GetUserOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSerumEventQueueAccount sets the "serumEventQueue" account.
func (inst *Deposit) SetSerumEventQueueAccount(serumEventQueue ag_solanago.PublicKey) *Deposit {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(serumEventQueue)
	return inst
}

// GetSerumEventQueueAccount gets the "serumEventQueue" account.
func (inst *Deposit) GetSerumEventQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst Deposit) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_Deposit),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Deposit) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Deposit) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MaxCoinAmount == nil {
			return errors.New("MaxCoinAmount parameter is not set")
		}
		if inst.MaxPcAmount == nil {
			return errors.New("MaxPcAmount parameter is not set")
		}
		if inst.BaseSide == nil {
			return errors.New("BaseSide parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Amm is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AmmAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.AmmOpenOrders is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.AmmTargetOrders is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.LpMintAddress is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PoolCoinTokenAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.PoolPcTokenAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SerumMarket is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.UserCoinTokenAccount is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.UserPcTokenAccount is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.UserLpTokenAccount is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.UserOwner is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.SerumEventQueue is not set")
		}
	}
	return nil
}

func (inst *Deposit) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Deposit")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MaxCoinAmount", *inst.MaxCoinAmount))
						paramsBranch.Child(ag_format.Param("  MaxPcAmount", *inst.MaxPcAmount))
						paramsBranch.Child(ag_format.Param("     BaseSide", *inst.BaseSide))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   tokenProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            amm", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   ammAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("  ammOpenOrders", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("ammTargetOrders", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("  lpMintAddress", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("  poolCoinToken", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("    poolPcToken", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("    serumMarket", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("  userCoinToken", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("    userPcToken", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("    userLpToken", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("      userOwner", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("serumEventQueue", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj Deposit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MaxCoinAmount` param:
	err = encoder.Encode(obj.MaxCoinAmount)
	if err != nil {
		return err
	}
	// Serialize `MaxPcAmount` param:
	err = encoder.Encode(obj.MaxPcAmount)
	if err != nil {
		return err
	}
	// Serialize `BaseSide` param:
	err = encoder.Encode(obj.BaseSide)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Deposit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MaxCoinAmount`:
	err = decoder.Decode(&obj.MaxCoinAmount)
	if err != nil {
		return err
	}
	// Deserialize `MaxPcAmount`:
	err = decoder.Decode(&obj.MaxPcAmount)
	if err != nil {
		return err
	}
	// Deserialize `BaseSide`:
	err = decoder.Decode(&obj.BaseSide)
	if err != nil {
		return err
	}
	return nil
}

// NewDepositInstruction declares a new Deposit instruction with the provided parameters and accounts.
func NewDepositInstruction(
	// Parameters:
	maxCoinAmount uint64,
	maxPcAmount uint64,
	baseSide uint64,
	// Accounts:
	tokenProgram ag_solanago.PublicKey,
	amm ag_solanago.PublicKey,
	ammAuthority ag_solanago.PublicKey,
	ammOpenOrders ag_solanago.PublicKey,
	ammTargetOrders ag_solanago.PublicKey,
	lpMintAddress ag_solanago.PublicKey,
	poolCoinTokenAccount ag_solanago.PublicKey,
	poolPcTokenAccount ag_solanago.PublicKey,
	serumMarket ag_solanago.PublicKey,
	userCoinTokenAccount ag_solanago.PublicKey,
	userPcTokenAccount ag_solanago.PublicKey,
	userLpTokenAccount ag_solanago.PublicKey,
	userOwner ag_solanago.PublicKey,
	serumEventQueue ag_solanago.PublicKey) *Deposit {
	return NewDepositInstructionBuilder().
		SetMaxCoinAmount(maxCoinAmount).
		SetMaxPcAmount(maxPcAmount).
		SetBaseSide(baseSide).
		SetTokenProgramAccount(tokenProgram).
		SetAmmAccount(amm).
		SetAmmAuthorityAccount(ammAuthority).
		SetAmmOpenOrdersAccount(ammOpenOrders).
		SetAmmTargetOrdersAccount(ammTargetOrders).
		SetLpMintAddressAccount(lpMintAddress).
		SetPoolCoinTokenAccountAccount(poolCoinTokenAccount).
		SetPoolPcTokenAccountAccount(poolPcTokenAccount).
		SetSerumMarketAccount(serumMarket).
		SetUserCoinTokenAccountAccount(userCoinTokenAccount).
		SetUserPcTokenAccountAccount(userPcTokenAccount).
		SetUserLpTokenAccountAccount(userLpTokenAccount).
		SetUserOwnerAccount(userOwner).
		SetSerumEventQueueAccount(serumEventQueue)
}
