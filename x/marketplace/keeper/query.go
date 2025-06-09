package keeper

import (
	"skillchain/x/marketplace/types"
)

var _ types.QueryServer = Keeper{}
