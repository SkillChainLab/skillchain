package keeper

import (
	"context"

	"skillchain/x/profile/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetSkillEndorsement set a specific skillEndorsement in the store from its index
func (k Keeper) SetSkillEndorsement(ctx context.Context, skillEndorsement types.SkillEndorsement) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SkillEndorsementKeyPrefix))
	b := k.cdc.MustMarshal(&skillEndorsement)
	store.Set(types.SkillEndorsementKey(
		skillEndorsement.Index,
	), b)
}

// GetSkillEndorsement returns a skillEndorsement from its index
func (k Keeper) GetSkillEndorsement(
	ctx context.Context,
	index string,

) (val types.SkillEndorsement, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SkillEndorsementKeyPrefix))

	b := store.Get(types.SkillEndorsementKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSkillEndorsement removes a skillEndorsement from the store
func (k Keeper) RemoveSkillEndorsement(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SkillEndorsementKeyPrefix))
	store.Delete(types.SkillEndorsementKey(
		index,
	))
}

// GetAllSkillEndorsement returns all skillEndorsement
func (k Keeper) GetAllSkillEndorsement(ctx context.Context) (list []types.SkillEndorsement) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SkillEndorsementKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SkillEndorsement
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
