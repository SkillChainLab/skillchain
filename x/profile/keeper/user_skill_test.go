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

func createNUserSkill(keeper keeper.Keeper, ctx context.Context, n int) []types.UserSkill {
	items := make([]types.UserSkill, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetUserSkill(ctx, items[i])
	}
	return items
}

func TestUserSkillGet(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNUserSkill(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUserSkill(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUserSkillRemove(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNUserSkill(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserSkill(ctx,
			item.Index,
		)
		_, found := keeper.GetUserSkill(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestUserSkillGetAll(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	items := createNUserSkill(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUserSkill(ctx)),
	)
}
