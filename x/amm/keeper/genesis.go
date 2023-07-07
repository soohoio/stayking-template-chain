package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking-template-chain/x/amm/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetParams(ctx),
	)
}
