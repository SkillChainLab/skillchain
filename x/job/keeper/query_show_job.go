package keeper

import (
	"context"

	"github.com/SkillChainLab/skillchain/x/job/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowJob(goCtx context.Context, req *types.QueryShowJobRequest) (*types.QueryShowJobResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryShowJobResponse{}, nil
}
