package profile_test

import (
	"testing"

	keepertest "skillchain/testutil/keeper"
	"skillchain/testutil/nullify"
	profile "skillchain/x/profile/module"
	"skillchain/x/profile/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ProfileKeeper(t)
	profile.InitGenesis(ctx, k, genesisState)
	got := profile.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.UserProfileList, got.UserProfileList)
	require.ElementsMatch(t, genesisState.UserSkillList, got.UserSkillList)
	require.ElementsMatch(t, genesisState.SkillEndorsementList, got.SkillEndorsementList)
	// this line is used by starport scaffolding # genesis/test/assert
}
