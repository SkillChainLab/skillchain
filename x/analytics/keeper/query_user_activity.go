package keeper

import (
	"context"

	"skillchain/x/analytics/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserActivityAll(ctx context.Context, req *types.QueryAllUserActivityRequest) (*types.QueryAllUserActivityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userActivitys []types.UserActivity

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	userActivityStore := prefix.NewStore(store, types.KeyPrefix(types.UserActivityKeyPrefix))

	pageRes, err := query.Paginate(userActivityStore, req.Pagination, func(key []byte, value []byte) error {
		var userActivity types.UserActivity
		if err := k.cdc.Unmarshal(value, &userActivity); err != nil {
			return err
		}

		userActivitys = append(userActivitys, userActivity)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserActivityResponse{UserActivity: userActivitys, Pagination: pageRes}, nil
}

func (k Keeper) UserActivity(ctx context.Context, req *types.QueryGetUserActivityRequest) (*types.QueryGetUserActivityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetUserActivity(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUserActivityResponse{UserActivity: val}, nil
}
