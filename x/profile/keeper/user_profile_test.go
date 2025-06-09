package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "skillchain/testutil/keeper"
	"skillchain/testutil/nullify"
	"skillchain/x/profile/keeper"
	"skillchain/x/profile/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUserProfile(keeper keeper.Keeper, ctx context.Context, n int) []types.UserProfile {
	items := make([]types.UserProfile, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetUserProfile(ctx, items[i])
	}
	return items
}

func TestUserProfileGet(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNUserProfile(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUserProfile(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUserProfileRemove(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNUserProfile(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserProfile(ctx,
			item.Index,
		)
		_, found := keeper.GetUserProfile(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestUserProfileGetAll(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNUserProfile(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUserProfile(ctx)),
	)
}
