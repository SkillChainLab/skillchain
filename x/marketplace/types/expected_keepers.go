package types

import (
	"context"

	profiletypes "skillchain/x/profile/types"
	skillchaintypes "skillchain/x/skillchain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ProfileKeeper interface {
	// User profile operations
	GetUserProfile(ctx context.Context, address string) (profiletypes.UserProfile, bool)
	SetUserProfile(ctx context.Context, profile profiletypes.UserProfile)
	GetAllUserProfile(ctx context.Context) []profiletypes.UserProfile
	// Reputation-based operations
	CalculateUserReputation(ctx context.Context, userAddress string) uint64
	UpdateUserReputation(ctx context.Context, userAddress string) error
}

type SkillchainKeeper interface {
	// vUSD operations for payments
	ConvertSkillToVUSD(ctx context.Context, skillAmount sdk.Coin) (sdk.Coin, error)
	ConvertVUSDToSkill(ctx context.Context, vUSDAmount sdk.Coin) (sdk.Coin, error)
	GetUserVUSDPosition(ctx context.Context, address string) (skillchaintypes.UserVUSDPosition, bool)
}

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoins(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}
