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

func createNPlatformMetric(keeper keeper.Keeper, ctx context.Context, n int) []types.PlatformMetric {
	items := make([]types.PlatformMetric, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetPlatformMetric(ctx, items[i])
	}
	return items
}

func TestPlatformMetricGet(t *testing.T) {
	keeper, ctx := keepertest.AnalyticsKeeper(t)
	items := createNPlatformMetric(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPlatformMetric(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPlatformMetricRemove(t *testing.T) {
	keeper, ctx := keepertest.AnalyticsKeeper(t)
	items := createNPlatformMetric(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePlatformMetric(ctx,
			item.Index,
		)
		_, found := keeper.GetPlatformMetric(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestPlatformMetricGetAll(t *testing.T) {
	keeper, ctx := keepertest.AnalyticsKeeper(t)
	items := createNPlatformMetric(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPlatformMetric(ctx)),
	)
}
