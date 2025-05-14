package job_test

import (
	"testing"

	keepertest "github.com/SkillChainLab/skillchain/testutil/keeper"
	"github.com/SkillChainLab/skillchain/testutil/nullify"
	job "github.com/SkillChainLab/skillchain/x/job/module"
	"github.com/SkillChainLab/skillchain/x/job/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.JobKeeper(t)
	job.InitGenesis(ctx, k, genesisState)
	got := job.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
