package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "skillchain/testutil/keeper"
	"skillchain/testutil/nullify"
	"skillchain/x/filestorage/keeper"
	"skillchain/x/filestorage/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNFilePermission(keeper keeper.Keeper, ctx context.Context, n int) []types.FilePermission {
	items := make([]types.FilePermission, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetFilePermission(ctx, items[i])
	}
	return items
}

func TestFilePermissionGet(t *testing.T) {
	keeper, ctx := keepertest.FilestorageKeeper(t)
	items := createNFilePermission(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFilePermission(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFilePermissionRemove(t *testing.T) {
	keeper, ctx := keepertest.FilestorageKeeper(t)
	items := createNFilePermission(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFilePermission(ctx,
			item.Index,
		)
		_, found := keeper.GetFilePermission(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestFilePermissionGetAll(t *testing.T) {
	keeper, ctx := keepertest.FilestorageKeeper(t)
	items := createNFilePermission(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFilePermission(ctx)),
	)
}
