package types_test

import (
	"testing"

	"skillchain/x/profile/types"

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

				UserProfileList: []types.UserProfile{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				UserSkillList: []types.UserSkill{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				SkillEndorsementList: []types.SkillEndorsement{
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
			desc: "duplicated userProfile",
			genState: &types.GenesisState{
				UserProfileList: []types.UserProfile{
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
			desc: "duplicated userSkill",
			genState: &types.GenesisState{
				UserSkillList: []types.UserSkill{
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
			desc: "duplicated skillEndorsement",
			genState: &types.GenesisState{
				SkillEndorsementList: []types.SkillEndorsement{
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
