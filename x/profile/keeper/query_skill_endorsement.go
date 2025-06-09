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

func (k Keeper) SkillEndorsementAll(ctx context.Context, req *types.QueryAllSkillEndorsementRequest) (*types.QueryAllSkillEndorsementResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var skillEndorsements []types.SkillEndorsement

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	skillEndorsementStore := prefix.NewStore(store, types.KeyPrefix(types.SkillEndorsementKeyPrefix))

	pageRes, err := query.Paginate(skillEndorsementStore, req.Pagination, func(key []byte, value []byte) error {
		var skillEndorsement types.SkillEndorsement
		if err := k.cdc.Unmarshal(value, &skillEndorsement); err != nil {
			return err
		}

		skillEndorsements = append(skillEndorsements, skillEndorsement)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSkillEndorsementResponse{SkillEndorsement: skillEndorsements, Pagination: pageRes}, nil
}

func (k Keeper) SkillEndorsement(ctx context.Context, req *types.QueryGetSkillEndorsementRequest) (*types.QueryGetSkillEndorsementResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetSkillEndorsement(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSkillEndorsementResponse{SkillEndorsement: val}, nil
}
