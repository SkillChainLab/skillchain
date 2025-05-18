package keeper

import (
	"context"

	"github.com/SkillChainLab/skillchain/x/job/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListJobsByCreator(goCtx context.Context, req *types.QueryListJobsByCreatorRequest) (*types.QueryListJobsByCreatorResponse, error) {
	if req == nil || req.Creator == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request: creator address required")
	}

	store := k.getStore(goCtx)
	var jobs []types.Job

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var job types.Job
		k.cdc.MustUnmarshal(iterator.Value(), &job)

		if job.Creator == req.Creator {
			jobs = append(jobs, job)
		}
	}

	return &types.QueryListJobsByCreatorResponse{
		Jobs: jobs,
	}, nil
} 