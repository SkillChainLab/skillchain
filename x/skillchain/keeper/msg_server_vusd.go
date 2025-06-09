package keeper

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/math"

	"skillchain/x/skillchain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ConvertSkillToVUSD converts SKILL tokens to vUSD
func (k msgServer) ConvertSkillToVUSD(goCtx context.Context, msg *types.MsgConvertSkillToVUSD) (*types.MsgConvertSkillToVUSDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate parameters
	params := k.GetParams(ctx)
	if !params.VusdEnabled {
		return nil, fmt.Errorf("vUSD functionality is disabled")
	}

	// Validate amount
	if msg.Amount.Denom != "uskill" {
		return nil, fmt.Errorf("invalid denom: expected uskill, got %s", msg.Amount.Denom)
	}

	skillAmount := msg.Amount.Amount
	if skillAmount.LTE(math.ZeroInt()) {
		return nil, fmt.Errorf("amount must be positive")
	}

	// Calculate vUSD amount to mint
	vUSDAmount, err := k.CalculateVUSDAmount(goCtx, skillAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate vUSD amount: %w", err)
	}

	// Simple 1:1 conversion based on price - no collateral ratio limitations
	// 1 SKILL ($0.50) = 0.5 vUSD directly

	// Get or create user position
	userPosition, found := k.GetUserVUSDPosition(goCtx, msg.Creator)
	if !found {
		userPosition = types.UserVUSDPosition{
			Address:         msg.Creator,
			SkillCollateral: math.ZeroInt(),
			VusdDebt:        math.ZeroInt(),
			CollateralRatio: "0",
			CreatedAt:       time.Now().Unix(),
			UpdatedAt:       time.Now().Unix(),
		}
	}

	// Update user position
	newSkillCollateral := userPosition.SkillCollateral.Add(skillAmount)
	newVUSDDebt := userPosition.VusdDebt.Add(vUSDAmount)

	// Calculate new collateral ratio for tracking purposes only
	newRatio, err := k.CalculateCollateralRatio(goCtx, newSkillCollateral, newVUSDDebt)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate collateral ratio: %w", err)
	}

	// Update user position
	userPosition.SkillCollateral = newSkillCollateral
	userPosition.VusdDebt = newVUSDDebt
	userPosition.CollateralRatio = newRatio
	userPosition.UpdatedAt = time.Now().Unix()

	// Transfer SKILL from user to module account
	skillCoin := sdk.NewCoin("uskill", skillAmount)
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("invalid creator address: %w", err)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, types.ModuleName, sdk.NewCoins(skillCoin))
	if err != nil {
		return nil, fmt.Errorf("failed to transfer SKILL: %w", err)
	}

	// Mint vUSD tokens
	vUSDCoin := sdk.NewCoin(VUSDDenom, vUSDAmount)
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(vUSDCoin))
	if err != nil {
		return nil, fmt.Errorf("failed to mint vUSD: %w", err)
	}

	// Send vUSD to user
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAddr, sdk.NewCoins(vUSDCoin))
	if err != nil {
		return nil, fmt.Errorf("failed to send vUSD: %w", err)
	}

	// Update treasury
	treasury, found := k.GetVUSDTreasury(goCtx)
	if !found {
		k.InitializeVUSDTreasury(goCtx)
		treasury, _ = k.GetVUSDTreasury(goCtx)
	}

	treasury.TotalSkillLocked = treasury.TotalSkillLocked.Add(skillAmount)
	treasury.TotalVusdSupply = treasury.TotalVusdSupply.Add(vUSDAmount)
	treasury.LastUpdate = time.Now().Unix()

	// Save state
	k.SetUserVUSDPosition(goCtx, userPosition)
	k.SetVUSDTreasury(goCtx, treasury)

	return &types.MsgConvertSkillToVUSDResponse{
		VusdMinted:      vUSDCoin,
		CollateralRatio: newRatio,
	}, nil
}

// ConvertVUSDToSkill converts vUSD tokens back to SKILL
func (k msgServer) ConvertVUSDToSkill(goCtx context.Context, msg *types.MsgConvertVUSDToSkill) (*types.MsgConvertVUSDToSkillResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate parameters
	params := k.GetParams(ctx)
	if !params.VusdEnabled {
		return nil, fmt.Errorf("vUSD functionality is disabled")
	}

	// Validate amount
	if msg.Amount.Denom != VUSDDenom {
		return nil, fmt.Errorf("invalid denom: expected %s, got %s", VUSDDenom, msg.Amount.Denom)
	}

	vUSDAmount := msg.Amount.Amount
	if vUSDAmount.LTE(math.ZeroInt()) {
		return nil, fmt.Errorf("amount must be positive")
	}

	// Calculate SKILL amount to return (simple price conversion)
	skillAmount, err := k.CalculateSkillAmount(goCtx, vUSDAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate SKILL amount: %w", err)
	}

	// Check if treasury has enough SKILL to release
	treasury, found := k.GetVUSDTreasury(goCtx)
	if !found {
		return nil, fmt.Errorf("treasury not found")
	}

	var skillSource string = "treasury_locked"
	var treasuryReserveAddress string = "skill12lqeht2us0jh6temhkf574vaugag9fcqqascw3" // Hardcoded for testing

	// Try to use locked SKILL first
	if treasury.TotalSkillLocked.GTE(skillAmount) {
		// Use locked SKILL from treasury
		skillSource = "treasury_locked"
	} else {
		// Check if treasury reserve can cover the difference
		// TODO: Use params.TreasuryReserveAddress when proto field is working

		if treasuryReserveAddress == "" {
			return nil, fmt.Errorf("insufficient SKILL in treasury: have %s, need %s, and no reserve configured",
				treasury.TotalSkillLocked.String(), skillAmount.String())
		}

		// Check treasury reserve balance
		treasuryReserveAddr, err := sdk.AccAddressFromBech32(treasuryReserveAddress)
		if err != nil {
			return nil, fmt.Errorf("invalid treasury reserve address: %w", err)
		}

		reserveBalance := k.bankKeeper.GetBalance(ctx, treasuryReserveAddr, "uskill")
		needed := skillAmount.Sub(treasury.TotalSkillLocked)

		if reserveBalance.Amount.LT(needed) {
			return nil, fmt.Errorf("insufficient SKILL: treasury has %s locked, reserve has %s, need total %s",
				treasury.TotalSkillLocked.String(), reserveBalance.Amount.String(), skillAmount.String())
		}

		skillSource = "treasury_reserve"
	}

	// Burn vUSD from user
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("invalid creator address: %w", err)
	}

	vUSDCoin := sdk.NewCoin(VUSDDenom, vUSDAmount)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, types.ModuleName, sdk.NewCoins(vUSDCoin))
	if err != nil {
		return nil, fmt.Errorf("failed to transfer vUSD: %w", err)
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(vUSDCoin))
	if err != nil {
		return nil, fmt.Errorf("failed to burn vUSD: %w", err)
	}

	// Release SKILL from treasury (module already has it)
	skillCoin := sdk.NewCoin("uskill", skillAmount)

	if skillSource == "treasury_locked" {
		// Use locked SKILL from module account
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAddr, sdk.NewCoins(skillCoin))
		if err != nil {
			return nil, fmt.Errorf("failed to send SKILL from treasury: %w", err)
		}

		// Update treasury - reduce locked SKILL
		treasury.TotalSkillLocked = treasury.TotalSkillLocked.Sub(skillAmount)

	} else {
		// Use reserve SKILL + remaining locked SKILL
		treasuryReserveAddr, _ := sdk.AccAddressFromBech32(treasuryReserveAddress)

		if treasury.TotalSkillLocked.GT(math.ZeroInt()) {
			// First, use all remaining locked SKILL
			lockedCoin := sdk.NewCoin("uskill", treasury.TotalSkillLocked)
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAddr, sdk.NewCoins(lockedCoin))
			if err != nil {
				return nil, fmt.Errorf("failed to send locked SKILL: %w", err)
			}
		}

		// Then, get remaining from reserve
		needed := skillAmount.Sub(treasury.TotalSkillLocked)
		if needed.GT(math.ZeroInt()) {
			reserveCoin := sdk.NewCoin("uskill", needed)
			err = k.bankKeeper.SendCoins(ctx, treasuryReserveAddr, creatorAddr, sdk.NewCoins(reserveCoin))
			if err != nil {
				return nil, fmt.Errorf("failed to send SKILL from reserve: %w", err)
			}
		}

		// Update treasury - all locked SKILL is used
		treasury.TotalSkillLocked = math.ZeroInt()
	}

	// Update treasury - reduce locked SKILL
	treasury.TotalSkillLocked = treasury.TotalSkillLocked.Sub(skillAmount)
	treasury.TotalVusdSupply = treasury.TotalVusdSupply.Sub(vUSDAmount)
	treasury.LastUpdate = time.Now().Unix()
	k.SetVUSDTreasury(goCtx, treasury)

	return &types.MsgConvertVUSDToSkillResponse{
		SkillReleased:   skillCoin,
		CollateralRatio: "100", // Always 100% for free conversion
	}, nil
}

// UpdateVUSDPrice updates the SKILL/USD price (authority only)
func (k msgServer) UpdateVUSDPrice(goCtx context.Context, msg *types.MsgUpdateVUSDPrice) (*types.MsgUpdateVUSDPriceResponse, error) {
	// Check authority
	params := k.GetParams(goCtx)
	if msg.Authority != params.PriceUpdateAuthority {
		return nil, fmt.Errorf("unauthorized: %s is not the price update authority", msg.Authority)
	}

	// Update params with new price
	params.VusdMockPrice = msg.NewPrice
	k.SetParams(goCtx, params)

	// Update treasury
	treasury, found := k.GetVUSDTreasury(goCtx)
	if !found {
		k.InitializeVUSDTreasury(goCtx)
		treasury, _ = k.GetVUSDTreasury(goCtx)
	}

	treasury.CurrentSkillPrice = msg.NewPrice
	treasury.LastUpdate = msg.Timestamp
	k.SetVUSDTreasury(goCtx, treasury)

	return &types.MsgUpdateVUSDPriceResponse{}, nil
}
