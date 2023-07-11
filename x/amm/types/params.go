package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyFeeRate             = []byte("FeeRate")
	KeyMinInitialLiquidity = []byte("MinInitialLiquidity")
)

var (
	DefaultFeeRate             = sdk.NewDecWithPrec(3, 3) // 0.3%
	DefaultMinInitialLiquidity = sdk.NewInt(1000)
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(feeRate sdk.Dec, minInitialLiquidity sdk.Int) Params {
	return Params{
		FeeRate:             feeRate,
		MinInitialLiquidity: minInitialLiquidity,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultFeeRate, DefaultMinInitialLiquidity)
}

// ParamSetPairs get the params.ParamSet
func (params *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyFeeRate, &params.FeeRate, validateFeeRate),
		paramtypes.NewParamSetPair(KeyMinInitialLiquidity, &params.MinInitialLiquidity, validateMinInitialLiquidity),
	}
}

// Validate validates the set of params
func (params Params) Validate() error {
	if err := validateFeeRate(params.FeeRate); err != nil {
		return err
	}
	if err := validateMinInitialLiquidity(params.MinInitialLiquidity); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("fee rate must not be negative: %s", v)
	}
	return nil
}

func validateMinInitialLiquidity(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("min initial liquidity must not be negative: %s", v)
	}
	return nil
}
