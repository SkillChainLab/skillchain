package keeper

import (
	"context"

	"skillchain/x/analytics/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetRevenueRecord set a specific revenueRecord in the store from its index
func (k Keeper) SetRevenueRecord(ctx context.Context, revenueRecord types.RevenueRecord) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RevenueRecordKeyPrefix))
	b := k.cdc.MustMarshal(&revenueRecord)
	store.Set(types.RevenueRecordKey(
		revenueRecord.Index,
	), b)
}

// GetRevenueRecord returns a revenueRecord from its index
func (k Keeper) GetRevenueRecord(
	ctx context.Context,
	index string,

) (val types.RevenueRecord, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RevenueRecordKeyPrefix))

	b := store.Get(types.RevenueRecordKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRevenueRecord removes a revenueRecord from the store
func (k Keeper) RemoveRevenueRecord(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RevenueRecordKeyPrefix))
	store.Delete(types.RevenueRecordKey(
		index,
	))
}

// GetAllRevenueRecord returns all revenueRecord
func (k Keeper) GetAllRevenueRecord(ctx context.Context) (list []types.RevenueRecord) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RevenueRecordKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RevenueRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
