package keeper

import (
	"context"

	"skillchain/x/marketplace/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetProposal set a specific proposal in the store from its index
func (k Keeper) SetProposal(ctx context.Context, proposal types.Proposal) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ProposalKeyPrefix))
	b := k.cdc.MustMarshal(&proposal)
	store.Set(types.ProposalKey(
		proposal.Index,
	), b)
}

// GetProposal returns a proposal from its index
func (k Keeper) GetProposal(
	ctx context.Context,
	index string,

) (val types.Proposal, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ProposalKeyPrefix))

	b := store.Get(types.ProposalKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProposal removes a proposal from the store
func (k Keeper) RemoveProposal(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ProposalKeyPrefix))
	store.Delete(types.ProposalKey(
		index,
	))
}

// GetAllProposal returns all proposal
func (k Keeper) GetAllProposal(ctx context.Context) (list []types.Proposal) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ProposalKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Proposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
