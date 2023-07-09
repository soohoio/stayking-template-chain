package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking-template-chain/x/amm/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
	return
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

func (k Keeper) GetFeeRate(ctx sdk.Context) (feeRate sdk.Dec) {
	k.paramSpace.Get(ctx, types.KeyFeeRate, &feeRate)
	return
}

func (k Keeper) GetMinInitialLiquidity(ctx sdk.Context) (minInitialLiquidity sdk.Int) {
	k.paramSpace.Get(ctx, types.KeyMinInitialLiquidity, &minInitialLiquidity)
	return
}
