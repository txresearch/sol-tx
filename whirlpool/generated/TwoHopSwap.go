// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package generated

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// TwoHopSwap is the `twoHopSwap` instruction.
type TwoHopSwap struct {
	Amount                 *uint64
	OtherAmountThreshold   *uint64
	AmountSpecifiedIsInput *bool
	AToBOne                *bool
	AToBTwo                *bool
	SqrtPriceLimitOne      *ag_binary.Uint128
	SqrtPriceLimitTwo      *ag_binary.Uint128

	// [0] = [] tokenProgram
	//
	// [1] = [SIGNER] tokenAuthority
	//
	// [2] = [WRITE] whirlpoolOne
	//
	// [3] = [WRITE] whirlpoolTwo
	//
	// [4] = [WRITE] tokenOwnerAccountOneA
	//
	// [5] = [WRITE] tokenVaultOneA
	//
	// [6] = [WRITE] tokenOwnerAccountOneB
	//
	// [7] = [WRITE] tokenVaultOneB
	//
	// [8] = [WRITE] tokenOwnerAccountTwoA
	//
	// [9] = [WRITE] tokenVaultTwoA
	//
	// [10] = [WRITE] tokenOwnerAccountTwoB
	//
	// [11] = [WRITE] tokenVaultTwoB
	//
	// [12] = [WRITE] tickArrayOne0
	//
	// [13] = [WRITE] tickArrayOne1
	//
	// [14] = [WRITE] tickArrayOne2
	//
	// [15] = [WRITE] tickArrayTwo0
	//
	// [16] = [WRITE] tickArrayTwo1
	//
	// [17] = [WRITE] tickArrayTwo2
	//
	// [18] = [] oracleOne
	//
	// [19] = [] oracleTwo
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewTwoHopSwapInstructionBuilder creates a new `TwoHopSwap` instruction builder.
func NewTwoHopSwapInstructionBuilder() *TwoHopSwap {
	nd := &TwoHopSwap{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 20),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *TwoHopSwap) SetAmount(amount uint64) *TwoHopSwap {
	inst.Amount = &amount
	return inst
}

// SetOtherAmountThreshold sets the "otherAmountThreshold" parameter.
func (inst *TwoHopSwap) SetOtherAmountThreshold(otherAmountThreshold uint64) *TwoHopSwap {
	inst.OtherAmountThreshold = &otherAmountThreshold
	return inst
}

// SetAmountSpecifiedIsInput sets the "amountSpecifiedIsInput" parameter.
func (inst *TwoHopSwap) SetAmountSpecifiedIsInput(amountSpecifiedIsInput bool) *TwoHopSwap {
	inst.AmountSpecifiedIsInput = &amountSpecifiedIsInput
	return inst
}

// SetAToBOne sets the "aToBOne" parameter.
func (inst *TwoHopSwap) SetAToBOne(aToBOne bool) *TwoHopSwap {
	inst.AToBOne = &aToBOne
	return inst
}

// SetAToBTwo sets the "aToBTwo" parameter.
func (inst *TwoHopSwap) SetAToBTwo(aToBTwo bool) *TwoHopSwap {
	inst.AToBTwo = &aToBTwo
	return inst
}

// SetSqrtPriceLimitOne sets the "sqrtPriceLimitOne" parameter.
func (inst *TwoHopSwap) SetSqrtPriceLimitOne(sqrtPriceLimitOne ag_binary.Uint128) *TwoHopSwap {
	inst.SqrtPriceLimitOne = &sqrtPriceLimitOne
	return inst
}

// SetSqrtPriceLimitTwo sets the "sqrtPriceLimitTwo" parameter.
func (inst *TwoHopSwap) SetSqrtPriceLimitTwo(sqrtPriceLimitTwo ag_binary.Uint128) *TwoHopSwap {
	inst.SqrtPriceLimitTwo = &sqrtPriceLimitTwo
	return inst
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *TwoHopSwap) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *TwoHopSwap) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTokenAuthorityAccount sets the "tokenAuthority" account.
func (inst *TwoHopSwap) SetTokenAuthorityAccount(tokenAuthority ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(tokenAuthority).SIGNER()
	return inst
}

// GetTokenAuthorityAccount gets the "tokenAuthority" account.
func (inst *TwoHopSwap) GetTokenAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetWhirlpoolOneAccount sets the "whirlpoolOne" account.
func (inst *TwoHopSwap) SetWhirlpoolOneAccount(whirlpoolOne ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(whirlpoolOne).WRITE()
	return inst
}

// GetWhirlpoolOneAccount gets the "whirlpoolOne" account.
func (inst *TwoHopSwap) GetWhirlpoolOneAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetWhirlpoolTwoAccount sets the "whirlpoolTwo" account.
func (inst *TwoHopSwap) SetWhirlpoolTwoAccount(whirlpoolTwo ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(whirlpoolTwo).WRITE()
	return inst
}

// GetWhirlpoolTwoAccount gets the "whirlpoolTwo" account.
func (inst *TwoHopSwap) GetWhirlpoolTwoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenOwnerAccountOneAAccount sets the "tokenOwnerAccountOneA" account.
func (inst *TwoHopSwap) SetTokenOwnerAccountOneAAccount(tokenOwnerAccountOneA ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenOwnerAccountOneA).WRITE()
	return inst
}

// GetTokenOwnerAccountOneAAccount gets the "tokenOwnerAccountOneA" account.
func (inst *TwoHopSwap) GetTokenOwnerAccountOneAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenVaultOneAAccount sets the "tokenVaultOneA" account.
func (inst *TwoHopSwap) SetTokenVaultOneAAccount(tokenVaultOneA ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenVaultOneA).WRITE()
	return inst
}

// GetTokenVaultOneAAccount gets the "tokenVaultOneA" account.
func (inst *TwoHopSwap) GetTokenVaultOneAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenOwnerAccountOneBAccount sets the "tokenOwnerAccountOneB" account.
func (inst *TwoHopSwap) SetTokenOwnerAccountOneBAccount(tokenOwnerAccountOneB ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenOwnerAccountOneB).WRITE()
	return inst
}

// GetTokenOwnerAccountOneBAccount gets the "tokenOwnerAccountOneB" account.
func (inst *TwoHopSwap) GetTokenOwnerAccountOneBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenVaultOneBAccount sets the "tokenVaultOneB" account.
func (inst *TwoHopSwap) SetTokenVaultOneBAccount(tokenVaultOneB ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenVaultOneB).WRITE()
	return inst
}

// GetTokenVaultOneBAccount gets the "tokenVaultOneB" account.
func (inst *TwoHopSwap) GetTokenVaultOneBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenOwnerAccountTwoAAccount sets the "tokenOwnerAccountTwoA" account.
func (inst *TwoHopSwap) SetTokenOwnerAccountTwoAAccount(tokenOwnerAccountTwoA ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenOwnerAccountTwoA).WRITE()
	return inst
}

// GetTokenOwnerAccountTwoAAccount gets the "tokenOwnerAccountTwoA" account.
func (inst *TwoHopSwap) GetTokenOwnerAccountTwoAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenVaultTwoAAccount sets the "tokenVaultTwoA" account.
func (inst *TwoHopSwap) SetTokenVaultTwoAAccount(tokenVaultTwoA ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenVaultTwoA).WRITE()
	return inst
}

// GetTokenVaultTwoAAccount gets the "tokenVaultTwoA" account.
func (inst *TwoHopSwap) GetTokenVaultTwoAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenOwnerAccountTwoBAccount sets the "tokenOwnerAccountTwoB" account.
func (inst *TwoHopSwap) SetTokenOwnerAccountTwoBAccount(tokenOwnerAccountTwoB ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenOwnerAccountTwoB).WRITE()
	return inst
}

// GetTokenOwnerAccountTwoBAccount gets the "tokenOwnerAccountTwoB" account.
func (inst *TwoHopSwap) GetTokenOwnerAccountTwoBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenVaultTwoBAccount sets the "tokenVaultTwoB" account.
func (inst *TwoHopSwap) SetTokenVaultTwoBAccount(tokenVaultTwoB ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenVaultTwoB).WRITE()
	return inst
}

// GetTokenVaultTwoBAccount gets the "tokenVaultTwoB" account.
func (inst *TwoHopSwap) GetTokenVaultTwoBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetTickArrayOne0Account sets the "tickArrayOne0" account.
func (inst *TwoHopSwap) SetTickArrayOne0Account(tickArrayOne0 ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(tickArrayOne0).WRITE()
	return inst
}

// GetTickArrayOne0Account gets the "tickArrayOne0" account.
func (inst *TwoHopSwap) GetTickArrayOne0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetTickArrayOne1Account sets the "tickArrayOne1" account.
func (inst *TwoHopSwap) SetTickArrayOne1Account(tickArrayOne1 ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(tickArrayOne1).WRITE()
	return inst
}

// GetTickArrayOne1Account gets the "tickArrayOne1" account.
func (inst *TwoHopSwap) GetTickArrayOne1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetTickArrayOne2Account sets the "tickArrayOne2" account.
func (inst *TwoHopSwap) SetTickArrayOne2Account(tickArrayOne2 ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(tickArrayOne2).WRITE()
	return inst
}

// GetTickArrayOne2Account gets the "tickArrayOne2" account.
func (inst *TwoHopSwap) GetTickArrayOne2Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetTickArrayTwo0Account sets the "tickArrayTwo0" account.
func (inst *TwoHopSwap) SetTickArrayTwo0Account(tickArrayTwo0 ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(tickArrayTwo0).WRITE()
	return inst
}

// GetTickArrayTwo0Account gets the "tickArrayTwo0" account.
func (inst *TwoHopSwap) GetTickArrayTwo0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetTickArrayTwo1Account sets the "tickArrayTwo1" account.
func (inst *TwoHopSwap) SetTickArrayTwo1Account(tickArrayTwo1 ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(tickArrayTwo1).WRITE()
	return inst
}

// GetTickArrayTwo1Account gets the "tickArrayTwo1" account.
func (inst *TwoHopSwap) GetTickArrayTwo1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetTickArrayTwo2Account sets the "tickArrayTwo2" account.
func (inst *TwoHopSwap) SetTickArrayTwo2Account(tickArrayTwo2 ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(tickArrayTwo2).WRITE()
	return inst
}

// GetTickArrayTwo2Account gets the "tickArrayTwo2" account.
func (inst *TwoHopSwap) GetTickArrayTwo2Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetOracleOneAccount sets the "oracleOne" account.
func (inst *TwoHopSwap) SetOracleOneAccount(oracleOne ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(oracleOne)
	return inst
}

// GetOracleOneAccount gets the "oracleOne" account.
func (inst *TwoHopSwap) GetOracleOneAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

// SetOracleTwoAccount sets the "oracleTwo" account.
func (inst *TwoHopSwap) SetOracleTwoAccount(oracleTwo ag_solanago.PublicKey) *TwoHopSwap {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(oracleTwo)
	return inst
}

// GetOracleTwoAccount gets the "oracleTwo" account.
func (inst *TwoHopSwap) GetOracleTwoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(19)
}

func (inst TwoHopSwap) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_TwoHopSwap,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst TwoHopSwap) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *TwoHopSwap) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
		if inst.OtherAmountThreshold == nil {
			return errors.New("OtherAmountThreshold parameter is not set")
		}
		if inst.AmountSpecifiedIsInput == nil {
			return errors.New("AmountSpecifiedIsInput parameter is not set")
		}
		if inst.AToBOne == nil {
			return errors.New("AToBOne parameter is not set")
		}
		if inst.AToBTwo == nil {
			return errors.New("AToBTwo parameter is not set")
		}
		if inst.SqrtPriceLimitOne == nil {
			return errors.New("SqrtPriceLimitOne parameter is not set")
		}
		if inst.SqrtPriceLimitTwo == nil {
			return errors.New("SqrtPriceLimitTwo parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TokenAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.WhirlpoolOne is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.WhirlpoolTwo is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenOwnerAccountOneA is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenVaultOneA is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenOwnerAccountOneB is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenVaultOneB is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenOwnerAccountTwoA is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenVaultTwoA is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenOwnerAccountTwoB is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenVaultTwoB is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.TickArrayOne0 is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.TickArrayOne1 is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.TickArrayOne2 is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.TickArrayTwo0 is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.TickArrayTwo1 is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.TickArrayTwo2 is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.OracleOne is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.OracleTwo is not set")
		}
	}
	return nil
}

func (inst *TwoHopSwap) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("TwoHopSwap")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=7]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("                Amount", *inst.Amount))
						paramsBranch.Child(ag_format.Param("  OtherAmountThreshold", *inst.OtherAmountThreshold))
						paramsBranch.Child(ag_format.Param("AmountSpecifiedIsInput", *inst.AmountSpecifiedIsInput))
						paramsBranch.Child(ag_format.Param("               AToBOne", *inst.AToBOne))
						paramsBranch.Child(ag_format.Param("               AToBTwo", *inst.AToBTwo))
						paramsBranch.Child(ag_format.Param("     SqrtPriceLimitOne", *inst.SqrtPriceLimitOne))
						paramsBranch.Child(ag_format.Param("     SqrtPriceLimitTwo", *inst.SqrtPriceLimitTwo))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=20]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         tokenProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       tokenAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         whirlpoolOne", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         whirlpoolTwo", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("tokenOwnerAccountOneA", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("       tokenVaultOneA", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("tokenOwnerAccountOneB", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       tokenVaultOneB", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("tokenOwnerAccountTwoA", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("       tokenVaultTwoA", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("tokenOwnerAccountTwoB", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("       tokenVaultTwoB", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("        tickArrayOne0", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("        tickArrayOne1", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("        tickArrayOne2", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("        tickArrayTwo0", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("        tickArrayTwo1", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("        tickArrayTwo2", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("            oracleOne", inst.AccountMetaSlice.Get(18)))
						accountsBranch.Child(ag_format.Meta("            oracleTwo", inst.AccountMetaSlice.Get(19)))
					})
				})
		})
}

func (obj TwoHopSwap) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `OtherAmountThreshold` param:
	err = encoder.Encode(obj.OtherAmountThreshold)
	if err != nil {
		return err
	}
	// Serialize `AmountSpecifiedIsInput` param:
	err = encoder.Encode(obj.AmountSpecifiedIsInput)
	if err != nil {
		return err
	}
	// Serialize `AToBOne` param:
	err = encoder.Encode(obj.AToBOne)
	if err != nil {
		return err
	}
	// Serialize `AToBTwo` param:
	err = encoder.Encode(obj.AToBTwo)
	if err != nil {
		return err
	}
	// Serialize `SqrtPriceLimitOne` param:
	err = encoder.Encode(obj.SqrtPriceLimitOne)
	if err != nil {
		return err
	}
	// Serialize `SqrtPriceLimitTwo` param:
	err = encoder.Encode(obj.SqrtPriceLimitTwo)
	if err != nil {
		return err
	}
	return nil
}
func (obj *TwoHopSwap) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `OtherAmountThreshold`:
	err = decoder.Decode(&obj.OtherAmountThreshold)
	if err != nil {
		return err
	}
	// Deserialize `AmountSpecifiedIsInput`:
	err = decoder.Decode(&obj.AmountSpecifiedIsInput)
	if err != nil {
		return err
	}
	// Deserialize `AToBOne`:
	err = decoder.Decode(&obj.AToBOne)
	if err != nil {
		return err
	}
	// Deserialize `AToBTwo`:
	err = decoder.Decode(&obj.AToBTwo)
	if err != nil {
		return err
	}
	// Deserialize `SqrtPriceLimitOne`:
	err = decoder.Decode(&obj.SqrtPriceLimitOne)
	if err != nil {
		return err
	}
	// Deserialize `SqrtPriceLimitTwo`:
	err = decoder.Decode(&obj.SqrtPriceLimitTwo)
	if err != nil {
		return err
	}
	return nil
}

// NewTwoHopSwapInstruction declares a new TwoHopSwap instruction with the provided parameters and accounts.
func NewTwoHopSwapInstruction(
	// Parameters:
	amount uint64,
	otherAmountThreshold uint64,
	amountSpecifiedIsInput bool,
	aToBOne bool,
	aToBTwo bool,
	sqrtPriceLimitOne ag_binary.Uint128,
	sqrtPriceLimitTwo ag_binary.Uint128,
	// Accounts:
	tokenProgram ag_solanago.PublicKey,
	tokenAuthority ag_solanago.PublicKey,
	whirlpoolOne ag_solanago.PublicKey,
	whirlpoolTwo ag_solanago.PublicKey,
	tokenOwnerAccountOneA ag_solanago.PublicKey,
	tokenVaultOneA ag_solanago.PublicKey,
	tokenOwnerAccountOneB ag_solanago.PublicKey,
	tokenVaultOneB ag_solanago.PublicKey,
	tokenOwnerAccountTwoA ag_solanago.PublicKey,
	tokenVaultTwoA ag_solanago.PublicKey,
	tokenOwnerAccountTwoB ag_solanago.PublicKey,
	tokenVaultTwoB ag_solanago.PublicKey,
	tickArrayOne0 ag_solanago.PublicKey,
	tickArrayOne1 ag_solanago.PublicKey,
	tickArrayOne2 ag_solanago.PublicKey,
	tickArrayTwo0 ag_solanago.PublicKey,
	tickArrayTwo1 ag_solanago.PublicKey,
	tickArrayTwo2 ag_solanago.PublicKey,
	oracleOne ag_solanago.PublicKey,
	oracleTwo ag_solanago.PublicKey) *TwoHopSwap {
	return NewTwoHopSwapInstructionBuilder().
		SetAmount(amount).
		SetOtherAmountThreshold(otherAmountThreshold).
		SetAmountSpecifiedIsInput(amountSpecifiedIsInput).
		SetAToBOne(aToBOne).
		SetAToBTwo(aToBTwo).
		SetSqrtPriceLimitOne(sqrtPriceLimitOne).
		SetSqrtPriceLimitTwo(sqrtPriceLimitTwo).
		SetTokenProgramAccount(tokenProgram).
		SetTokenAuthorityAccount(tokenAuthority).
		SetWhirlpoolOneAccount(whirlpoolOne).
		SetWhirlpoolTwoAccount(whirlpoolTwo).
		SetTokenOwnerAccountOneAAccount(tokenOwnerAccountOneA).
		SetTokenVaultOneAAccount(tokenVaultOneA).
		SetTokenOwnerAccountOneBAccount(tokenOwnerAccountOneB).
		SetTokenVaultOneBAccount(tokenVaultOneB).
		SetTokenOwnerAccountTwoAAccount(tokenOwnerAccountTwoA).
		SetTokenVaultTwoAAccount(tokenVaultTwoA).
		SetTokenOwnerAccountTwoBAccount(tokenOwnerAccountTwoB).
		SetTokenVaultTwoBAccount(tokenVaultTwoB).
		SetTickArrayOne0Account(tickArrayOne0).
		SetTickArrayOne1Account(tickArrayOne1).
		SetTickArrayOne2Account(tickArrayOne2).
		SetTickArrayTwo0Account(tickArrayTwo0).
		SetTickArrayTwo1Account(tickArrayTwo1).
		SetTickArrayTwo2Account(tickArrayTwo2).
		SetOracleOneAccount(oracleOne).
		SetOracleTwoAccount(oracleTwo)
}