package keeper

import (
	"context"

	"skillchain/x/analytics/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetPlatformMetric set a specific platformMetric in the store from its index
func (k Keeper) SetPlatformMetric(ctx context.Context, platformMetric types.PlatformMetric) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlatformMetricKeyPrefix))
	b := k.cdc.MustMarshal(&platformMetric)
	store.Set(types.PlatformMetricKey(
		platformMetric.Index,
	), b)
}

// GetPlatformMetric returns a platformMetric from its index
func (k Keeper) GetPlatformMetric(
	ctx context.Context,
	index string,

) (val types.PlatformMetric, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlatformMetricKeyPrefix))

	b := store.Get(types.PlatformMetricKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePlatformMetric removes a platformMetric from the store
func (k Keeper) RemovePlatformMetric(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlatformMetricKeyPrefix))
	store.Delete(types.PlatformMetricKey(
		index,
	))
}

// GetAllPlatformMetric returns all platformMetric
func (k Keeper) GetAllPlatformMetric(ctx context.Context) (list []types.PlatformMetric) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlatformMetricKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PlatformMetric
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
