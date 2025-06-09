package notifications_test

import (
	"testing"

	keepertest "skillchain/testutil/keeper"
	"skillchain/testutil/nullify"
	notifications "skillchain/x/notifications/module"
	"skillchain/x/notifications/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		NotificationList: []types.Notification{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		NotificationSettingsList: []types.NotificationSettings{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NotificationsKeeper(t)
	notifications.InitGenesis(ctx, k, genesisState)
	got := notifications.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.NotificationList, got.NotificationList)
	require.ElementsMatch(t, genesisState.NotificationSettingsList, got.NotificationSettingsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
