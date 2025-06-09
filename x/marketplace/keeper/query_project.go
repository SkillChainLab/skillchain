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

func (k Keeper) ProjectAll(ctx context.Context, req *types.QueryAllProjectRequest) (*types.QueryAllProjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var projects []types.Project

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	projectStore := prefix.NewStore(store, types.KeyPrefix(types.ProjectKeyPrefix))

	pageRes, err := query.Paginate(projectStore, req.Pagination, func(key []byte, value []byte) error {
		var project types.Project
		if err := k.cdc.Unmarshal(value, &project); err != nil {
			return err
		}

		projects = append(projects, project)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProjectResponse{Project: projects, Pagination: pageRes}, nil
}

func (k Keeper) Project(ctx context.Context, req *types.QueryGetProjectRequest) (*types.QueryGetProjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetProject(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProjectResponse{Project: val}, nil
}
