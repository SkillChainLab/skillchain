package keeper

import (
	"context"
	"strconv"

	"github.com/SkillChainLab/skillchain/x/job/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowJob(goCtx context.Context, req *types.QueryShowJobRequest) (*types.QueryShowJobResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	id, err := strconv.ParseUint(req.Id, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid job ID")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.getStore(ctx)
	bz := store.Get(types.JobKey(id))
	if bz == nil {
		return nil, status.Error(codes.NotFound, "job not found")
	}

	var job types.Job
	k.cdc.MustUnmarshal(bz, &job)

	return &types.QueryShowJobResponse{Job: &job}, nil
}
