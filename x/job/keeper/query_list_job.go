package keeper

import (
	"context"

	"github.com/SkillChainLab/skillchain/x/job/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListJob(goCtx context.Context, req *types.QueryListJobRequest) (*types.QueryListJobResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.getStore(ctx)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var jobs []types.Job

	for ; iterator.Valid(); iterator.Next() {
		var job types.Job
		k.cdc.MustUnmarshal(iterator.Value(), &job)
		jobs = append(jobs, job)
	}

	return &types.QueryListJobResponse{Jobs: jobs}, nil
}
