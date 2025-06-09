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

func createNProposal(keeper keeper.Keeper, ctx context.Context, n int) []types.Proposal {
	items := make([]types.Proposal, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetProposal(ctx, items[i])
	}
	return items
}

func TestProposalGet(t *testing.T) {
	keeper, ctx := keepertest.MarketplaceKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProposal(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestProposalRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketplaceKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProposal(ctx,
			item.Index,
		)
		_, found := keeper.GetProposal(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestProposalGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketplaceKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProposal(ctx)),
	)
}
