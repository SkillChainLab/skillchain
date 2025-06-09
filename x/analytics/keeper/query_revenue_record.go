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

func (k Keeper) RevenueRecordAll(ctx context.Context, req *types.QueryAllRevenueRecordRequest) (*types.QueryAllRevenueRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var revenueRecords []types.RevenueRecord

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	revenueRecordStore := prefix.NewStore(store, types.KeyPrefix(types.RevenueRecordKeyPrefix))

	pageRes, err := query.Paginate(revenueRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var revenueRecord types.RevenueRecord
		if err := k.cdc.Unmarshal(value, &revenueRecord); err != nil {
			return err
		}

		revenueRecords = append(revenueRecords, revenueRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRevenueRecordResponse{RevenueRecord: revenueRecords, Pagination: pageRes}, nil
}

func (k Keeper) RevenueRecord(ctx context.Context, req *types.QueryGetRevenueRecordRequest) (*types.QueryGetRevenueRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetRevenueRecord(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRevenueRecordResponse{RevenueRecord: val}, nil
}
