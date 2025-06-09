package keeper

import (
	"context"
	"fmt"

	"skillchain/x/skillchain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// VUSDTreasury returns the global vUSD treasury state
func (k Keeper) VUSDTreasury(goCtx context.Context, req *types.QueryVUSDTreasuryRequest) (*types.QueryVUSDTreasuryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get VUSD treasury information
	skillBalance := k.bankKeeper.GetBalance(ctx, sdk.AccAddress{}, "uskill")
	vusdSupply := k.bankKeeper.GetSupply(ctx, "vusd")

	return &types.QueryVUSDTreasuryResponse{
		// Treasury: treasury, // TODO: implement treasury tracking
		SkillBalance: skillBalance.Amount.String(),
		VusdSupply:   vusdSupply.Amount.String(),
		ExchangeRate: "1.0", // TODO: implement dynamic exchange rate
	}, nil
}

// UserVUSDPosition returns a user's vUSD position
func (k Keeper) UserVUSDPosition(goCtx context.Context, req *types.QueryUserVUSDPositionRequest) (*types.QueryUserVUSDPositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Parse the address
	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid address: %s", err))
	}

	// Get user's VUSD position (placeholder implementation)
	vusdBalance := k.bankKeeper.GetBalance(ctx, address, "vusd")
	
	return &types.QueryUserVUSDPositionResponse{
		// Position: types.UserVUSDPosition{}, // TODO: implement position tracking
		VusdBalance:     vusdBalance.Amount.String(),
		SkillCollateral: "0",
		HealthFactor:    "0",
		Position:        "healthy", // TODO: implement position calculation
		Exists:          true,
	}, nil

	// TODO: Implement proper position tracking
	return &types.QueryUserVUSDPositionResponse{
		// Position: position, // TODO: implement position tracking
		VusdBalance:     "0",
		SkillCollateral: "0",
		HealthFactor:    "0",
		Position:        "none",
		Exists:          false,
	}, nil
}
