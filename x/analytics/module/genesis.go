package analytics

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"skillchain/x/analytics/keeper"
	"skillchain/x/analytics/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the platformMetric
	for _, elem := range genState.PlatformMetricList {
		k.SetPlatformMetric(ctx, elem)
	}
	// Set all the userActivity
	for _, elem := range genState.UserActivityList {
		k.SetUserActivity(ctx, elem)
	}
	// Set all the revenueRecord
	for _, elem := range genState.RevenueRecordList {
		k.SetRevenueRecord(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PlatformMetricList = k.GetAllPlatformMetric(ctx)
	genesis.UserActivityList = k.GetAllUserActivity(ctx)
	genesis.RevenueRecordList = k.GetAllRevenueRecord(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
