package keeper

import (
	"context"

	"github.com/SkillChainLab/skillchain/x/profile/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowProfile(goCtx context.Context, req *types.QueryShowProfileRequest) (*types.QueryShowProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	profile, found := k.GetProfile(ctx, req.Username)
	if !found {
		return nil, status.Error(codes.NotFound, "profile not found")
	}

	return &types.QueryShowProfileResponse{
		Username: profile.Username,
		Bio:      profile.Bio,
	}, nil
}
