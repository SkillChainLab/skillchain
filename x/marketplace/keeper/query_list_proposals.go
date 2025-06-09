package keeper

import (
	"context"

	"skillchain/x/marketplace/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListProposals(goCtx context.Context, req *types.QueryListProposalsRequest) (*types.QueryListProposalsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryListProposalsResponse{}, nil
}
