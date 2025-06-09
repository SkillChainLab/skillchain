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

func (k Keeper) MilestoneAll(ctx context.Context, req *types.QueryAllMilestoneRequest) (*types.QueryAllMilestoneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var milestones []types.Milestone

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	milestoneStore := prefix.NewStore(store, types.KeyPrefix(types.MilestoneKeyPrefix))

	pageRes, err := query.Paginate(milestoneStore, req.Pagination, func(key []byte, value []byte) error {
		var milestone types.Milestone
		if err := k.cdc.Unmarshal(value, &milestone); err != nil {
			return err
		}

		milestones = append(milestones, milestone)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMilestoneResponse{Milestone: milestones, Pagination: pageRes}, nil
}

func (k Keeper) Milestone(ctx context.Context, req *types.QueryGetMilestoneRequest) (*types.QueryGetMilestoneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetMilestone(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMilestoneResponse{Milestone: val}, nil
}
