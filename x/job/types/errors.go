package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/job module sentinel errors
var (
	ErrInvalidSigner        = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample               = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrJobNotFound          = sdkerrors.Register(ModuleName, 1102, "job not found")
	ErrApplicationNotFound  = sdkerrors.Register(ModuleName, 1103, "application not found")
)
