package keeper

import (
	"context"

	"github.com/SkillChainLab/skillchain/x/job/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListMyApplications(goCtx context.Context, req *types.QueryListMyApplicationsRequest) (*types.QueryListMyApplicationsResponse, error) {
	if req == nil || req.Applicant == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request: applicant required")
	}

	store := k.getStore(goCtx)
	var applications []*types.Application

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var app types.Application
		k.cdc.MustUnmarshal(iterator.Value(), &app)

		if app.Applicant == req.Applicant {
			copy := app
			applications = append(applications, &copy)
		}
	}

	return &types.QueryListMyApplicationsResponse{
		Applications: applications,
	}, nil
}
