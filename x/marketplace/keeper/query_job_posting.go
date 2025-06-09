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

func (k Keeper) JobPostingAll(ctx context.Context, req *types.QueryAllJobPostingRequest) (*types.QueryAllJobPostingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var jobPostings []types.JobPosting

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	jobPostingStore := prefix.NewStore(store, types.KeyPrefix(types.JobPostingKeyPrefix))

	pageRes, err := query.Paginate(jobPostingStore, req.Pagination, func(key []byte, value []byte) error {
		var jobPosting types.JobPosting
		if err := k.cdc.Unmarshal(value, &jobPosting); err != nil {
			return err
		}

		jobPostings = append(jobPostings, jobPosting)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllJobPostingResponse{JobPosting: jobPostings, Pagination: pageRes}, nil
}

func (k Keeper) JobPosting(ctx context.Context, req *types.QueryGetJobPostingRequest) (*types.QueryGetJobPostingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetJobPosting(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetJobPostingResponse{JobPosting: val}, nil
}
