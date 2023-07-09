package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking-template-chain/util"
)

const (
	ModuleName  = "amm"
	StoreKey    = ModuleName
	RouterKey   = ModuleName
	MemStoreKey = ModuleName
)

var (
	LastPairIDKey      = []byte{0x01}
	PairKeyPrefix      = []byte{0x02}
	PairIndexKeyPrefix = []byte{0x03}
)

func GetPairKey(pairID uint64) []byte {
	return append(PairKeyPrefix, sdk.Uint64ToBigEndian(pairID)...)
}

func GetPairIndexKey(denom0, denom1 string) []byte {
	return append(
		append(PairIndexKeyPrefix, util.LengthPrefix([]byte(denom0))...), denom1...)
}
