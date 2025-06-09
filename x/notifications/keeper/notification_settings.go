package keeper

import (
	"context"

	"skillchain/x/notifications/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetNotificationSettings set a specific notificationSettings in the store from its index
func (k Keeper) SetNotificationSettings(ctx context.Context, notificationSettings types.NotificationSettings) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NotificationSettingsKeyPrefix))
	b := k.cdc.MustMarshal(&notificationSettings)
	store.Set(types.NotificationSettingsKey(
		notificationSettings.Index,
	), b)
}

// GetNotificationSettings returns a notificationSettings from its index
func (k Keeper) GetNotificationSettings(
	ctx context.Context,
	index string,

) (val types.NotificationSettings, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NotificationSettingsKeyPrefix))

	b := store.Get(types.NotificationSettingsKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNotificationSettings removes a notificationSettings from the store
func (k Keeper) RemoveNotificationSettings(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NotificationSettingsKeyPrefix))
	store.Delete(types.NotificationSettingsKey(
		index,
	))
}

// GetAllNotificationSettings returns all notificationSettings
func (k Keeper) GetAllNotificationSettings(ctx context.Context) (list []types.NotificationSettings) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NotificationSettingsKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NotificationSettings
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
