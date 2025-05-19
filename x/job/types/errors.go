package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/job module sentinel errors
var (
	ErrInvalidSigner        = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample               = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrJobNotFound          = sdkerrors.Register(ModuleName, 1, "job not found")
	ErrJobAlreadyExists     = sdkerrors.Register(ModuleName, 2, "job already exists")
	ErrUnauthorized         = sdkerrors.Register(ModuleName, 3, "unauthorized")
	ErrApplicationNotFound  = sdkerrors.Register(ModuleName, 4, "application not found")
	ErrNotificationCreation = sdkerrors.Register(ModuleName, 5, "failed to create notification")
)
