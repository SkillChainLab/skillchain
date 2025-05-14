package keeper

import (
	"github.com/SkillChainLab/skillchain/x/profile/types"
)

var _ types.QueryServer = Keeper{}
