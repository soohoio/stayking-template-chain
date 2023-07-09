package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v3"
)

var _ paramstypes.ParamSet = (*Params)(nil)

var (
	KeyFeeRate             = []byte("FeeRate")
	KeyMinInitialLiquidity = []byte("MinInitialLiquidity")
)

var (
	DefaultFeeRate             = sdk.NewDecWithPrec(3, 3) // 0.3%
	DefaultMinInitialLiquidity = sdk.NewInt(1000)
)

func NewParams(feeRate sdk.Dec, minInitialLiquidity sdk.Int) Params {
	return Params{
		FeeRate:             feeRate,
		MinInitialLiquidity: minInitialLiquidity,
	}
}

func DefaultParams() Params {
	return NewParams(DefaultFeeRate, DefaultMinInitialLiquidity)
}

func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (params *Params) Validate() error {
	if err := validagteFeeRate(params.FeeRate); err != nil {
		return err
	}

	return nil
}

func (params *Params) String() string {
	out, _ := yaml.Marshal(params)

	return string(out)
}

func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{}
}

func validagteFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("fee rate must not be negative : %s", v)
	}
	return nil
}
