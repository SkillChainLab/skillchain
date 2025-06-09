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

func (k Keeper) PlatformMetricAll(ctx context.Context, req *types.QueryAllPlatformMetricRequest) (*types.QueryAllPlatformMetricResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var platformMetrics []types.PlatformMetric

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	platformMetricStore := prefix.NewStore(store, types.KeyPrefix(types.PlatformMetricKeyPrefix))

	pageRes, err := query.Paginate(platformMetricStore, req.Pagination, func(key []byte, value []byte) error {
		var platformMetric types.PlatformMetric
		if err := k.cdc.Unmarshal(value, &platformMetric); err != nil {
			return err
		}

		platformMetrics = append(platformMetrics, platformMetric)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPlatformMetricResponse{PlatformMetric: platformMetrics, Pagination: pageRes}, nil
}

func (k Keeper) PlatformMetric(ctx context.Context, req *types.QueryGetPlatformMetricRequest) (*types.QueryGetPlatformMetricResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetPlatformMetric(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPlatformMetricResponse{PlatformMetric: val}, nil
}
