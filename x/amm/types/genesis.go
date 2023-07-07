package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName  = "amm"
	StoreKey    = ModuleName
	RouterKey   = ModuleName
	MemStoreKey = ModuleName
)

func NewGenesisState(params Params) *GenesisState {
	return &GenesisState{
		Params: params,
	}
}

func DefaultGenesis() *GenesisState {
	return NewGenesisState(DefaultParams())
}

func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	return nil
}

func (k Keeper) InitGenesis(ctx sdk.Context, genState tyㅎpes.GenesisState) {

}
