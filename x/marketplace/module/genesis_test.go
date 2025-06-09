package marketplace_test

import (
	"testing"

	keepertest "skillchain/testutil/keeper"
	"skillchain/testutil/nullify"
	marketplace "skillchain/x/marketplace/module"
	"skillchain/x/marketplace/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		JobPostingList: []types.JobPosting{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ProposalList: []types.Proposal{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ProjectList: []types.Project{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		MilestoneList: []types.Milestone{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MarketplaceKeeper(t)
	marketplace.InitGenesis(ctx, k, genesisState)
	got := marketplace.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.JobPostingList, got.JobPostingList)
	require.ElementsMatch(t, genesisState.ProposalList, got.ProposalList)
	require.ElementsMatch(t, genesisState.ProjectList, got.ProjectList)
	require.ElementsMatch(t, genesisState.MilestoneList, got.MilestoneList)
	// this line is used by starport scaffolding # genesis/test/assert
}
