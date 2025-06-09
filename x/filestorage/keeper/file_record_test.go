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

func createNFileRecord(keeper keeper.Keeper, ctx context.Context, n int) []types.FileRecord {
	items := make([]types.FileRecord, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetFileRecord(ctx, items[i])
	}
	return items
}

func TestFileRecordGet(t *testing.T) {
	keeper, ctx := keepertest.FilestorageKeeper(t)
	items := createNFileRecord(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFileRecord(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFileRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.FilestorageKeeper(t)
	items := createNFileRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFileRecord(ctx,
			item.Index,
		)
		_, found := keeper.GetFileRecord(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestFileRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.FilestorageKeeper(t)
	items := createNFileRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFileRecord(ctx)),
	)
}
