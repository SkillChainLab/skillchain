package analytics_test

import (
	"testing"

	keepertest "skillchain/testutil/keeper"
	"skillchain/testutil/nullify"
	analytics "skillchain/x/analytics/module"
	"skillchain/x/analytics/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PlatformMetricList: []types.PlatformMetric{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		UserActivityList: []types.UserActivity{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		RevenueRecordList: []types.RevenueRecord{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AnalyticsKeeper(t)
	analytics.InitGenesis(ctx, k, genesisState)
	got := analytics.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PlatformMetricList, got.PlatformMetricList)
	require.ElementsMatch(t, genesisState.UserActivityList, got.UserActivityList)
	require.ElementsMatch(t, genesisState.RevenueRecordList, got.RevenueRecordList)
	// this line is used by starport scaffolding # genesis/test/assert
}
