package types

import (
	sdkerrors "cosmossdk.io/errors"
)

var (
	ErrProfileNotFound = sdkerrors.New(ModuleName, 1100, "profile not found")
	ErrInvalidRequest  = sdkerrors.New(ModuleName, 1101, "invalid request")
)
