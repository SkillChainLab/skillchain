package keeper

import (
	"skillchain/x/profile/types"
)

var _ types.QueryServer = Keeper{}
