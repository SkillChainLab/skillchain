package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "skillchain/testutil/keeper"
	"skillchain/x/filestorage/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.FilestorageKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
