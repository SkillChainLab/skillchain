package types_test

import (
	"testing"

	"skillchain/x/marketplace/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated jobPosting",
			genState: &types.GenesisState{
				JobPostingList: []types.JobPosting{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated proposal",
			genState: &types.GenesisState{
				ProposalList: []types.Proposal{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated project",
			genState: &types.GenesisState{
				ProjectList: []types.Project{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated milestone",
			genState: &types.GenesisState{
				MilestoneList: []types.Milestone{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
