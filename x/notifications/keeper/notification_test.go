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

func createNNotification(keeper keeper.Keeper, ctx context.Context, n int) []types.Notification {
	items := make([]types.Notification, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetNotification(ctx, items[i])
	}
	return items
}

func TestNotificationGet(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	items := createNNotification(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNotification(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNotificationRemove(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	items := createNNotification(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNotification(ctx,
			item.Index,
		)
		_, found := keeper.GetNotification(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestNotificationGetAll(t *testing.T) {
	keeper, ctx := keepertest.NotificationsKeeper(t)
	items := createNNotification(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNotification(ctx)),
	)
}
