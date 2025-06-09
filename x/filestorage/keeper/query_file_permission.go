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

func (k Keeper) FilePermissionAll(ctx context.Context, req *types.QueryAllFilePermissionRequest) (*types.QueryAllFilePermissionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var filePermissions []types.FilePermission

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	filePermissionStore := prefix.NewStore(store, types.KeyPrefix(types.FilePermissionKeyPrefix))

	pageRes, err := query.Paginate(filePermissionStore, req.Pagination, func(key []byte, value []byte) error {
		var filePermission types.FilePermission
		if err := k.cdc.Unmarshal(value, &filePermission); err != nil {
			return err
		}

		filePermissions = append(filePermissions, filePermission)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFilePermissionResponse{FilePermission: filePermissions, Pagination: pageRes}, nil
}

func (k Keeper) FilePermission(ctx context.Context, req *types.QueryGetFilePermissionRequest) (*types.QueryGetFilePermissionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetFilePermission(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetFilePermissionResponse{FilePermission: val}, nil
}
