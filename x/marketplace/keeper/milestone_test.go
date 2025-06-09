package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "skillchain/testutil/keeper"
	"skillchain/testutil/nullify"
	"skillchain/x/marketplace/keeper"
	"skillchain/x/marketplace/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMilestone(keeper keeper.Keeper, ctx context.Context, n int) []types.Milestone {
	items := make([]types.Milestone, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetMilestone(ctx, items[i])
	}
	return items
}

func TestMilestoneGet(t *testing.T) {
	keeper, ctx := keepertest.MarketplaceKeeper(t)
	items := createNMilestone(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMilestone(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMilestoneRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketplaceKeeper(t)
	items := createNMilestone(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMilestone(ctx,
			item.Index,
		)
		_, found := keeper.GetMilestone(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestMilestoneGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketplaceKeeper(t)
	items := createNMilestone(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMilestone(ctx)),
	)
}
