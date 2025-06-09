package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"skillchain/x/marketplace/keeper"
	"skillchain/x/marketplace/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the jobPosting
	for _, elem := range genState.JobPostingList {
		k.SetJobPosting(ctx, elem)
	}
	// Set all the proposal
	for _, elem := range genState.ProposalList {
		k.SetProposal(ctx, elem)
	}
	// Set all the project
	for _, elem := range genState.ProjectList {
		k.SetProject(ctx, elem)
	}
	// Set all the milestone
	for _, elem := range genState.MilestoneList {
		k.SetMilestone(ctx, elem)
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

	genesis.JobPostingList = k.GetAllJobPosting(ctx)
	genesis.ProposalList = k.GetAllProposal(ctx)
	genesis.ProjectList = k.GetAllProject(ctx)
	genesis.MilestoneList = k.GetAllMilestone(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
