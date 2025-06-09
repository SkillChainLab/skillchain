package keeper

import (
	"context"

	"skillchain/x/notifications/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetNotification set a specific notification in the store from its index
func (k Keeper) SetNotification(ctx context.Context, notification types.Notification) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NotificationKeyPrefix))
	b := k.cdc.MustMarshal(&notification)
	store.Set(types.NotificationKey(
		notification.Index,
	), b)
}

// GetNotification returns a notification from its index
func (k Keeper) GetNotification(
	ctx context.Context,
	index string,

) (val types.Notification, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NotificationKeyPrefix))

	b := store.Get(types.NotificationKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNotification removes a notification from the store
func (k Keeper) RemoveNotification(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NotificationKeyPrefix))
	store.Delete(types.NotificationKey(
		index,
	))
}

// GetAllNotification returns all notification
func (k Keeper) GetAllNotification(ctx context.Context) (list []types.Notification) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NotificationKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Notification
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
