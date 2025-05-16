package keeper

import (
	"github.com/SkillChainLab/skillchain/x/application/types"
)

var _ types.QueryServer = Keeper{}
