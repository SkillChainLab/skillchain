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

func createNJobPosting(keeper keeper.Keeper, ctx context.Context, n int) []types.JobPosting {
	items := make([]types.JobPosting, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetJobPosting(ctx, items[i])
	}
	return items
}

func TestJobPostingGet(t *testing.T) {
	keeper, ctx := keepertest.MarketplaceKeeper(t)
	items := createNJobPosting(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetJobPosting(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestJobPostingRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketplaceKeeper(t)
	items := createNJobPosting(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveJobPosting(ctx,
			item.Index,
		)
		_, found := keeper.GetJobPosting(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestJobPostingGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketplaceKeeper(t)
	items := createNJobPosting(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllJobPosting(ctx)),
	)
}
