package keeper

import (
	"context"
	"fmt"

	"skillchain/x/skillchain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Parse the creator address
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("invalid creator address: %w", err)
	}

	// Validate the amount
	if !msg.Amount.IsValid() || msg.Amount.IsZero() {
		return nil, sdkerrors.ErrInvalidCoins.Wrap("invalid burn amount")
	}

	// Check if burn is enabled
	params := k.GetParams(ctx)
	if !params.BurnEnabled {
		return nil, sdkerrors.ErrUnauthorized.Wrap("token burning is disabled")
	}

	// Check if the creator has enough balance
	balance := k.bankKeeper.GetBalance(ctx, creator, msg.Amount.Denom)
	if balance.IsLT(msg.Amount) {
		return nil, sdkerrors.ErrInsufficientFunds.Wrap("insufficient balance to burn")
	}

	// Burn tokens by sending them to module account and then burning from there
	coins := sdk.Coins{msg.Amount}

	// Send coins to module account
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, coins)
	if err != nil {
		return nil, fmt.Errorf("failed to send coins to module account: %w", err)
	}

	// Burn coins from module account
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, fmt.Errorf("failed to burn coins: %w", err)
	}

	// Emit an event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"skill_token_burn",
			sdk.NewAttribute("burner", msg.Creator),
			sdk.NewAttribute("amount", msg.Amount.String()),
			sdk.NewAttribute("remaining_supply", k.bankKeeper.GetSupply(ctx, msg.Amount.Denom).Amount.String()),
		),
	)

	return &types.MsgBurnResponse{}, nil
}
