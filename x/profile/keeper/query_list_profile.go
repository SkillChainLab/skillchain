package keeper

import (
	"context"

	"github.com/SkillChainLab/skillchain/x/profile/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListProfile(goCtx context.Context, req *types.QueryListProfileRequest) (*types.QueryListProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	profiles := k.GetAllProfile(goCtx)

	return &types.QueryListProfileResponse{
		Profiles: profiles,
	}, nil
}
