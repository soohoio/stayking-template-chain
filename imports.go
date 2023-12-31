package stayking

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	staykingapp "github.com/soohoio/stayking-template-chain/app"
)

var (
	_ = storetypes.MemoryStoreKey{}
	_ = sdkerrors.Wrap
	_ = tmproto.Header{}
	_ = staykingapp.Setup
	_ = codectypes.AminoJSONPacker{}
	_ = cryptocodec.RegisterCrypto
	_ = authzcodec.Amino
	_ = simtypes.RandSubsetCoins
	_ = abci.RequestEndBlock{}
)
