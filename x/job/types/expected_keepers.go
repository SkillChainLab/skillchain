package types

import (
	"context"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// JobKeeper defines the expected interface for the Job module.
type JobKeeper interface {
	GetJob(ctx context.Context, id uint64) (Job, bool)
	SetJob(ctx context.Context, job Job)
	AppendJob(ctx context.Context, title, desc, budget, creator string) uint64
	GetJobCount(ctx context.Context) uint64
	SetJobCount(ctx context.Context, count uint64)
	AppendApplication(ctx context.Context, jobID uint64, applicant string, coverLetter string) error
	GetApplicationCount(ctx context.Context) uint64
	SetApplicationCount(ctx context.Context, count uint64)
	GetAuthority() string
	Logger() log.Logger
}

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}
