package keeper

import (
	context "context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking-template-chain/x/amm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(k Keeper) Querier {
	return Querier{}
}

func (q Querier) Params(_ctx context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(_ctx)

	return &types.QueryParamsResponse{Params: q.GetParams(ctx)}, nil
}
