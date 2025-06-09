package keeper

import (
	"context"

	"skillchain/x/filestorage/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetFileRecord set a specific fileRecord in the store from its index
func (k Keeper) SetFileRecord(ctx context.Context, fileRecord types.FileRecord) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FileRecordKeyPrefix))
	b := k.cdc.MustMarshal(&fileRecord)
	store.Set(types.FileRecordKey(
		fileRecord.Index,
	), b)
}

// GetFileRecord returns a fileRecord from its index
func (k Keeper) GetFileRecord(
	ctx context.Context,
	index string,

) (val types.FileRecord, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FileRecordKeyPrefix))

	b := store.Get(types.FileRecordKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFileRecord removes a fileRecord from the store
func (k Keeper) RemoveFileRecord(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FileRecordKeyPrefix))
	store.Delete(types.FileRecordKey(
		index,
	))
}

// GetAllFileRecord returns all fileRecord
func (k Keeper) GetAllFileRecord(ctx context.Context) (list []types.FileRecord) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FileRecordKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FileRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
