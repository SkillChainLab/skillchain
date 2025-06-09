package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "skillchain/testutil/keeper"
	"skillchain/testutil/nullify"
	"skillchain/x/notifications/keeper"
	"skillchain/x/notifications/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNotificationSettings(keeper keeper.Keeper, ctx context.Context, n int) []types.NotificationSettings {
	items := make([]types.NotificationSettings, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetNotificationSettings(ctx, items[i])
	}
	return items
}

func TestNotificationSettingsGet(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	items := createNNotificationSettings(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNotificationSettings(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNotificationSettingsRemove(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	items := createNNotificationSettings(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNotificationSettings(ctx,
			item.Index,
		)
		_, found := keeper.GetNotificationSettings(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestNotificationSettingsGetAll(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	items := createNNotificationSettings(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNotificationSettings(ctx)),
	)
}
