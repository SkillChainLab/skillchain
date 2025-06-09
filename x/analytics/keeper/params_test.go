package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "skillchain/testutil/keeper"
	"skillchain/x/analytics/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AnalyticsKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
