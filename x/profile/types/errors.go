package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/profile module sentinel errors
var (
	ErrInvalidSigner        = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample               = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrProfileAlreadyExists = sdkerrors.Register(ModuleName, 1102, "profile already exists for this user")
)
