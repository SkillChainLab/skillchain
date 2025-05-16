package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/SkillChainLab/skillchain/testutil/keeper"
	"github.com/SkillChainLab/skillchain/x/application/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ApplicationKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
