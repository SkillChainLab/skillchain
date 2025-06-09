package keeper

import (
	"context"

	"skillchain/x/marketplace/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetMilestone set a specific milestone in the store from its index
func (k Keeper) SetMilestone(ctx context.Context, milestone types.Milestone) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MilestoneKeyPrefix))
	b := k.cdc.MustMarshal(&milestone)
	store.Set(types.MilestoneKey(
		milestone.Index,
	), b)
}

// GetMilestone returns a milestone from its index
func (k Keeper) GetMilestone(
	ctx context.Context,
	index string,

) (val types.Milestone, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MilestoneKeyPrefix))

	b := store.Get(types.MilestoneKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMilestone removes a milestone from the store
func (k Keeper) RemoveMilestone(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MilestoneKeyPrefix))
	store.Delete(types.MilestoneKey(
		index,
	))
}

// GetAllMilestone returns all milestone
func (k Keeper) GetAllMilestone(ctx context.Context) (list []types.Milestone) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MilestoneKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Milestone
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
