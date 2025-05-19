package keeper

import (
	"context"
	"strings"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

// SearchJobs searches for jobs by title or description
func (k Keeper) SearchJobs(ctx context.Context, req *types.QuerySearchJobsRequest) (*types.QuerySearchJobsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.JobKeyPrefix))

	var jobs []types.Job
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var job types.Job
		if err := k.cdc.Unmarshal(value, &job); err != nil {
			return err
		}

		// Check if the job matches the search term
		searchTerm := strings.ToLower(req.SearchTerm)
		title := strings.ToLower(job.Title)
		description := strings.ToLower(job.Description)

		if strings.Contains(title, searchTerm) || strings.Contains(description, searchTerm) {
			jobs = append(jobs, job)
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySearchJobsResponse{
		Jobs:       jobs,
		Pagination: pageRes,
	}, nil
} 