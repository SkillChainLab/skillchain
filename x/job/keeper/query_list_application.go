package keeper

import (
	"context"

	"github.com/SkillChainLab/skillchain/x/job/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListApplication(goCtx context.Context, req *types.QueryListApplicationRequest) (*types.QueryListApplicationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	store := k.getStore(goCtx)

	var applications []types.Application

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var app types.Application
		k.cdc.MustUnmarshal(iterator.Value(), &app)

		if app.JobId == req.JobId {
			applications = append(applications, app)
		}
	}

	return &types.QueryListApplicationResponse{Applications: applications}, nil
}
