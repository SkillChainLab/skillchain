package keeper

import (
	"context"

	"skillchain/x/filestorage/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FileRecordAll(ctx context.Context, req *types.QueryAllFileRecordRequest) (*types.QueryAllFileRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var fileRecords []types.FileRecord

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	fileRecordStore := prefix.NewStore(store, types.KeyPrefix(types.FileRecordKeyPrefix))

	pageRes, err := query.Paginate(fileRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var fileRecord types.FileRecord
		if err := k.cdc.Unmarshal(value, &fileRecord); err != nil {
			return err
		}

		fileRecords = append(fileRecords, fileRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFileRecordResponse{FileRecord: fileRecords, Pagination: pageRes}, nil
}

func (k Keeper) FileRecord(ctx context.Context, req *types.QueryGetFileRecordRequest) (*types.QueryGetFileRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetFileRecord(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetFileRecordResponse{FileRecord: val}, nil
}
