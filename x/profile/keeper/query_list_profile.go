package keeper

import (
	"context"

	"github.com/SkillChainLab/skillchain/x/profile/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListProfile(goCtx context.Context, req *types.QueryListProfileRequest) (*types.QueryListProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := k.getStore(ctx)

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var profiles []*types.Profile
	for ; iterator.Valid(); iterator.Next() {
		var profile types.Profile
		k.cdc.MustUnmarshal(iterator.Value(), &profile)
		profiles = append(profiles, &profile)
	}

	return &types.QueryListProfileResponse{Profiles: profiles}, nil
}
