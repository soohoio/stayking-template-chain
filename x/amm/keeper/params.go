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
