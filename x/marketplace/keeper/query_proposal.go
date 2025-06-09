package keeper

import (
	"context"

	"skillchain/x/marketplace/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProposalAll(ctx context.Context, req *types.QueryAllProposalRequest) (*types.QueryAllProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proposals []types.Proposal

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	proposalStore := prefix.NewStore(store, types.KeyPrefix(types.ProposalKeyPrefix))

	pageRes, err := query.Paginate(proposalStore, req.Pagination, func(key []byte, value []byte) error {
		var proposal types.Proposal
		if err := k.cdc.Unmarshal(value, &proposal); err != nil {
			return err
		}

		proposals = append(proposals, proposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProposalResponse{Proposal: proposals, Pagination: pageRes}, nil
}

func (k Keeper) Proposal(ctx context.Context, req *types.QueryGetProposalRequest) (*types.QueryGetProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProposalResponse{Proposal: val}, nil
}
