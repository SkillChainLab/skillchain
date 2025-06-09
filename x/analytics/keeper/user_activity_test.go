package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "skillchain/testutil/keeper"
	"skillchain/testutil/nullify"
	"skillchain/x/analytics/keeper"
	"skillchain/x/analytics/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUserActivity(keeper keeper.Keeper, ctx context.Context, n int) []types.UserActivity {
	items := make([]types.UserActivity, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetUserActivity(ctx, items[i])
	}
	return items
}

func TestUserActivityGet(t *testing.T) {
	keeper, ctx := keepertest.AnalyticsKeeper(t)
	items := createNUserActivity(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUserActivity(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUserActivityRemove(t *testing.T) {
	keeper, ctx := keepertest.AnalyticsKeeper(t)
	items := createNUserActivity(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserActivity(ctx,
			item.Index,
		)
		_, found := keeper.GetUserActivity(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestUserActivityGetAll(t *testing.T) {
	keeper, ctx := keepertest.AnalyticsKeeper(t)
	items := createNUserActivity(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUserActivity(ctx)),
	)
}
