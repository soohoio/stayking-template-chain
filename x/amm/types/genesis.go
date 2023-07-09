package types

func NewGenesisState(params Params, lastPairID uint64, pairs []Pair) *GenesisState {
	return &GenesisState{
		Params:     params,
		LastPairId: lastPairID,
		Pairs:      pairs,
	}
}

func DefaultGenesis() *GenesisState {
	return NewGenesisState(DefaultParams(), 0, nil)
}

func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	return nil
}
