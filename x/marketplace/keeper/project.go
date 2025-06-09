package keeper

import (
	"context"

	"skillchain/x/marketplace/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetProject set a specific project in the store from its index
func (k Keeper) SetProject(ctx context.Context, project types.Project) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ProjectKeyPrefix))
	b := k.cdc.MustMarshal(&project)
	store.Set(types.ProjectKey(
		project.Index,
	), b)
}

// GetProject returns a project from its index
func (k Keeper) GetProject(
	ctx context.Context,
	index string,

) (val types.Project, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ProjectKeyPrefix))

	b := store.Get(types.ProjectKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProject removes a project from the store
func (k Keeper) RemoveProject(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ProjectKeyPrefix))
	store.Delete(types.ProjectKey(
		index,
	))
}

// GetAllProject returns all project
func (k Keeper) GetAllProject(ctx context.Context) (list []types.Project) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ProjectKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Project
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
