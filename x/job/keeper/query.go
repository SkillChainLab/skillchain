package keeper

import (
	"context"
	"fmt"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListJobApplications(ctx context.Context, req *types.QueryListJobApplicationsRequest) (*types.QueryListJobApplicationsResponse, error) {
	store := k.getStore(ctx)
	applications := []*types.Application{}

	// Create the prefix for the job ID
	prefix := fmt.Sprintf("Application/value/%d", req.JobId)
	iterator := store.Iterator([]byte(prefix), nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var application types.Application
		k.cdc.MustUnmarshal(iterator.Value(), &application)
		applications = append(applications, &application)
	}

	return &types.QueryListJobApplicationsResponse{
		Applications: applications,
	}, nil
}
