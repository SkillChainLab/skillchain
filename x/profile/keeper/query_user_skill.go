package keeper

import (
	"context"

	"skillchain/x/profile/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserSkillAll(ctx context.Context, req *types.QueryAllUserSkillRequest) (*types.QueryAllUserSkillResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userSkills []types.UserSkill

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	userSkillStore := prefix.NewStore(store, types.KeyPrefix(types.UserSkillKeyPrefix))

	pageRes, err := query.Paginate(userSkillStore, req.Pagination, func(key []byte, value []byte) error {
		var userSkill types.UserSkill
		if err := k.cdc.Unmarshal(value, &userSkill); err != nil {
			return err
		}

		userSkills = append(userSkills, userSkill)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserSkillResponse{UserSkill: userSkills, Pagination: pageRes}, nil
}

func (k Keeper) UserSkill(ctx context.Context, req *types.QueryGetUserSkillRequest) (*types.QueryGetUserSkillResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetUserSkill(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUserSkillResponse{UserSkill: val}, nil
}
