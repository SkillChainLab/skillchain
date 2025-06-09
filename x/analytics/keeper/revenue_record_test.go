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

func createNRevenueRecord(keeper keeper.Keeper, ctx context.Context, n int) []types.RevenueRecord {
	items := make([]types.RevenueRecord, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetRevenueRecord(ctx, items[i])
	}
	return items
}

func TestRevenueRecordGet(t *testing.T) {
	keeper, ctx := keepertest.AnalyticsKeeper(t)
	items := createNRevenueRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRevenueRecord(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRevenueRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.AnalyticsKeeper(t)
	items := createNRevenueRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRevenueRecord(ctx,
			item.Index,
		)
		_, found := keeper.GetRevenueRecord(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestRevenueRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.AnalyticsKeeper(t)
	items := createNRevenueRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRevenueRecord(ctx)),
	)
}
