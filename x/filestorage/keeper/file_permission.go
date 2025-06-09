package keeper

import (
	"context"

	"skillchain/x/filestorage/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetFilePermission set a specific filePermission in the store from its index
func (k Keeper) SetFilePermission(ctx context.Context, filePermission types.FilePermission) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FilePermissionKeyPrefix))
	b := k.cdc.MustMarshal(&filePermission)
	store.Set(types.FilePermissionKey(
		filePermission.Index,
	), b)
}

// GetFilePermission returns a filePermission from its index
func (k Keeper) GetFilePermission(
	ctx context.Context,
	index string,

) (val types.FilePermission, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FilePermissionKeyPrefix))

	b := store.Get(types.FilePermissionKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFilePermission removes a filePermission from the store
func (k Keeper) RemoveFilePermission(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FilePermissionKeyPrefix))
	store.Delete(types.FilePermissionKey(
		index,
	))
}

// GetAllFilePermission returns all filePermission
func (k Keeper) GetAllFilePermission(ctx context.Context) (list []types.FilePermission) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FilePermissionKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FilePermission
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
