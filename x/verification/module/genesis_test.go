package verification_test

import (
	"testing"

	keepertest "github.com/SkillChainLab/skillchain/testutil/keeper"
	"github.com/SkillChainLab/skillchain/testutil/nullify"
	verification "github.com/SkillChainLab/skillchain/x/verification/module"
	"github.com/SkillChainLab/skillchain/x/verification/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VerificationKeeper(t)
	verification.InitGenesis(ctx, k, genesisState)
	got := verification.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
