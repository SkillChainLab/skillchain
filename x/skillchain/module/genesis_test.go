package skillchain_test

import (
	"testing"

	keepertest "skillchain/testutil/keeper"
	"skillchain/testutil/nullify"
	skillchain "skillchain/x/skillchain/module"
	"skillchain/x/skillchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SkillchainKeeper(t)
	skillchain.InitGenesis(ctx, k, genesisState)
	got := skillchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
