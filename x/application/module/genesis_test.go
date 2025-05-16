package application_test

import (
	"testing"

	keepertest "github.com/SkillChainLab/skillchain/testutil/keeper"
	"github.com/SkillChainLab/skillchain/testutil/nullify"
	application "github.com/SkillChainLab/skillchain/x/application/module"
	"github.com/SkillChainLab/skillchain/x/application/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ApplicationKeeper(t)
	application.InitGenesis(ctx, k, genesisState)
	got := application.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
