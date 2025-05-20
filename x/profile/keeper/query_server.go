package keeper

import (
	"context"

	"github.com/SkillChainLab/skillchain/x/profile/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QueryServerImpl struct {
	Keeper
}

var _ types.QueryServer = QueryServerImpl{}

// ShowProfile implements the Query/ShowProfile gRPC method
func (k QueryServerImpl) ShowProfile(ctx context.Context, req *types.QueryShowProfileRequest) (*types.QueryShowProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	profile, found := k.GetProfile(ctx, req.Username)
	if !found {
		return nil, status.Error(codes.NotFound, "profile not found")
	}

	return &types.QueryShowProfileResponse{
		Profile: &profile,
	}, nil
}

// ListProfile implements the Query/ListProfile gRPC method
func (k QueryServerImpl) ListProfile(ctx context.Context, req *types.QueryListProfileRequest) (*types.QueryListProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	profiles := k.GetAllProfile(ctx)

	return &types.QueryListProfileResponse{
		Profiles: profiles,
	}, nil
}

func (k QueryServerImpl) ShowProfileByAddress(ctx context.Context, req *types.QueryShowProfileByAddressRequest) (*types.QueryShowProfileByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Implement profile lookup by address
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (k QueryServerImpl) ListProfiles(ctx context.Context, req *types.QueryListProfilesRequest) (*types.QueryListProfilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	profiles := k.GetAllProfile(ctx)

	return &types.QueryListProfilesResponse{
		Profiles: profiles,
	}, nil
}
