package keeper

import (
	"context"

	"skillchain/x/marketplace/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetJobPosting set a specific jobPosting in the store from its index
func (k Keeper) SetJobPosting(ctx context.Context, jobPosting types.JobPosting) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.JobPostingKeyPrefix))
	b := k.cdc.MustMarshal(&jobPosting)
	store.Set(types.JobPostingKey(
		jobPosting.Index,
	), b)
}

// GetJobPosting returns a jobPosting from its index
func (k Keeper) GetJobPosting(
	ctx context.Context,
	index string,

) (val types.JobPosting, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.JobPostingKeyPrefix))

	b := store.Get(types.JobPostingKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveJobPosting removes a jobPosting from the store
func (k Keeper) RemoveJobPosting(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.JobPostingKeyPrefix))
	store.Delete(types.JobPostingKey(
		index,
	))
}

// GetAllJobPosting returns all jobPosting
func (k Keeper) GetAllJobPosting(ctx context.Context) (list []types.JobPosting) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.JobPostingKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.JobPosting
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
