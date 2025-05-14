package keeper

import (
	"github.com/SkillChainLab/skillchain/x/job/types"
)

var _ types.QueryServer = Keeper{}
