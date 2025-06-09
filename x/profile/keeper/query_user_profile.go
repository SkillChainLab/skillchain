package keeper

import (
	"context"

	"skillchain/x/profile/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserProfileAll(ctx context.Context, req *types.QueryAllUserProfileRequest) (*types.QueryAllUserProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userProfiles []types.UserProfile

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	userProfileStore := prefix.NewStore(store, types.KeyPrefix(types.UserProfileKeyPrefix))

	pageRes, err := query.Paginate(userProfileStore, req.Pagination, func(key []byte, value []byte) error {
		var userProfile types.UserProfile
		if err := k.cdc.Unmarshal(value, &userProfile); err != nil {
			return err
		}

		userProfiles = append(userProfiles, userProfile)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserProfileResponse{UserProfile: userProfiles, Pagination: pageRes}, nil
}

func (k Keeper) UserProfile(ctx context.Context, req *types.QueryGetUserProfileRequest) (*types.QueryGetUserProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetUserProfile(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUserProfileResponse{UserProfile: val}, nil
}
