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
		Profile: &profile,
	}, nil
}

func (k Keeper) ShowProfileByAddress(goCtx context.Context, req *types.QueryShowProfileByAddressRequest) (*types.QueryShowProfileByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get all profiles
	profiles := k.GetAllProfile(ctx)

	// Find profile by creator address
	for _, profile := range profiles {
		if profile.Creator == req.Creator {
			return &types.QueryShowProfileByAddressResponse{
				Profile: &profile,
			}, nil
		}
	}

	return nil, status.Error(codes.NotFound, "profile not found")
}

func (k Keeper) ListProfiles(goCtx context.Context, req *types.QueryListProfilesRequest) (*types.QueryListProfilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	profiles := k.GetAllProfile(ctx)

	return &types.QueryListProfilesResponse{
		Profiles: profiles,
	}, nil
}
