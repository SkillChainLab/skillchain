package keeper

import (
	"context"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListJobApplications(ctx context.Context, req *types.QueryListJobApplicationsRequest) (*types.QueryListJobApplicationsResponse, error) {
	store := k.getStore(ctx)
	applications := []*types.Application{}

	// Get all applications for the specified job
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var application types.Application
		k.cdc.MustUnmarshal(iterator.Value(), &application)

		if application.JobId == req.JobId {
			applications = append(applications, &application)
		}
	}

	return &types.QueryListJobApplicationsResponse{
		Applications: applications,
	}, nil
}
