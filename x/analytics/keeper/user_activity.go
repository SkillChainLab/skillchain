package keeper

import (
	"context"

	"skillchain/x/analytics/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetUserActivity set a specific userActivity in the store from its index
func (k Keeper) SetUserActivity(ctx context.Context, userActivity types.UserActivity) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserActivityKeyPrefix))
	b := k.cdc.MustMarshal(&userActivity)
	store.Set(types.UserActivityKey(
		userActivity.Index,
	), b)
}

// GetUserActivity returns a userActivity from its index
func (k Keeper) GetUserActivity(
	ctx context.Context,
	index string,

) (val types.UserActivity, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserActivityKeyPrefix))

	b := store.Get(types.UserActivityKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserActivity removes a userActivity from the store
func (k Keeper) RemoveUserActivity(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserActivityKeyPrefix))
	store.Delete(types.UserActivityKey(
		index,
	))
}

// GetAllUserActivity returns all userActivity
func (k Keeper) GetAllUserActivity(ctx context.Context) (list []types.UserActivity) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserActivityKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserActivity
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
