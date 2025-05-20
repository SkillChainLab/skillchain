package keeper

import (
	"context"
	"strconv"
	"strings"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

// FilterJobs filters jobs based on various criteria
func (k Keeper) FilterJobs(ctx context.Context, req *types.QueryFilterJobsRequest) (*types.QueryFilterJobsResponse, error) {
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

		// Apply filters
		if req.Title != "" && !strings.Contains(strings.ToLower(job.Title), strings.ToLower(req.Title)) {
			return nil
		}

		if req.Creator != "" && job.Creator != req.Creator {
			return nil
		}

		// Parse and compare budget
		if req.MinBudget != "" || req.MaxBudget != "" {
			jobBudget, err := strconv.ParseInt(strings.TrimSuffix(job.Budget, "stake"), 10, 64)
			if err != nil {
				return err
			}

			if req.MinBudget != "" {
				minBudget, err := strconv.ParseInt(strings.TrimSuffix(req.MinBudget, "stake"), 10, 64)
				if err != nil {
					return err
				}
				if jobBudget < minBudget {
					return nil
				}
			}

			if req.MaxBudget != "" {
				maxBudget, err := strconv.ParseInt(strings.TrimSuffix(req.MaxBudget, "stake"), 10, 64)
				if err != nil {
					return err
				}
				if jobBudget > maxBudget {
					return nil
				}
			}
		}

		jobs = append(jobs, job)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryFilterJobsResponse{
		Jobs:       jobs,
		Pagination: pageRes,
	}, nil
}
