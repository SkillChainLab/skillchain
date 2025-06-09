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

func createNSkillEndorsement(keeper keeper.Keeper, ctx context.Context, n int) []types.SkillEndorsement {
	items := make([]types.SkillEndorsement, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetSkillEndorsement(ctx, items[i])
	}
	return items
}

func TestSkillEndorsementGet(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNSkillEndorsement(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSkillEndorsement(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSkillEndorsementRemove(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNSkillEndorsement(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSkillEndorsement(ctx,
			item.Index,
		)
		_, found := keeper.GetSkillEndorsement(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestSkillEndorsementGetAll(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNSkillEndorsement(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSkillEndorsement(ctx)),
	)
}
