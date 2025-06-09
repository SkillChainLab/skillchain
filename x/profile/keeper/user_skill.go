package keeper

import (
	"context"

	"skillchain/x/profile/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetUserSkill set a specific userSkill in the store from its index
func (k Keeper) SetUserSkill(ctx context.Context, userSkill types.UserSkill) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserSkillKeyPrefix))
	b := k.cdc.MustMarshal(&userSkill)
	store.Set(types.UserSkillKey(
		userSkill.Index,
	), b)
}

// GetUserSkill returns a userSkill from its index
func (k Keeper) GetUserSkill(
	ctx context.Context,
	index string,

) (val types.UserSkill, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserSkillKeyPrefix))

	b := store.Get(types.UserSkillKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserSkill removes a userSkill from the store
func (k Keeper) RemoveUserSkill(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserSkillKeyPrefix))
	store.Delete(types.UserSkillKey(
		index,
	))
}

// GetAllUserSkill returns all userSkill
func (k Keeper) GetAllUserSkill(ctx context.Context) (list []types.UserSkill) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserSkillKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserSkill
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
