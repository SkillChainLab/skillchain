package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/profile module sentinel errors
var (
	ErrProfileNotFound = errorsmod.Register("profile", 1100, "profile not found")
	ErrUnauthorized    = errorsmod.Register("profile", 1101, "unauthorized")
)
