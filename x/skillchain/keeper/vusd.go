package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"cosmossdk.io/math"

	"skillchain/x/skillchain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// vUSD denomination
const VUSDDenom = "uvusd"

// GetVUSDTreasury returns the global vUSD treasury state
func (k Keeper) GetVUSDTreasury(ctx context.Context) (treasury types.VUSDTreasury, found bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := k.storeService.OpenKVStore(sdkCtx)

	b, err := store.Get([]byte("vusd_treasury"))
	if err != nil || b == nil {
		return treasury, false
	}

	k.cdc.MustUnmarshal(b, &treasury)
	return treasury, true
}

// SetVUSDTreasury sets the global vUSD treasury state
func (k Keeper) SetVUSDTreasury(ctx context.Context, treasury types.VUSDTreasury) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := k.storeService.OpenKVStore(sdkCtx)

	b := k.cdc.MustMarshal(&treasury)
	store.Set([]byte("vusd_treasury"), b)
}

// GetUserVUSDPosition returns a user's vUSD position
func (k Keeper) GetUserVUSDPosition(ctx context.Context, address string) (position types.UserVUSDPosition, found bool) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := k.storeService.OpenKVStore(sdkCtx)

	key := []byte(fmt.Sprintf("vusd_position_%s", address))
	b, err := store.Get(key)
	if err != nil || b == nil {
		return position, false
	}

	k.cdc.MustUnmarshal(b, &position)
	return position, true
}

// SetUserVUSDPosition sets a user's vUSD position
func (k Keeper) SetUserVUSDPosition(ctx context.Context, position types.UserVUSDPosition) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := k.storeService.OpenKVStore(sdkCtx)

	key := []byte(fmt.Sprintf("vusd_position_%s", position.Address))
	b := k.cdc.MustMarshal(&position)
	store.Set(key, b)
}

// InitializeVUSDTreasury initializes the treasury with default values
func (k Keeper) InitializeVUSDTreasury(ctx context.Context) {
	treasury := types.VUSDTreasury{
		TotalSkillLocked:      math.ZeroInt(),
		TotalVusdSupply:       math.ZeroInt(),
		CurrentSkillPrice:     "0.50", // Mock price
		LastUpdate:            time.Now().Unix(),
		GlobalCollateralRatio: "0",
	}
	k.SetVUSDTreasury(ctx, treasury)
}

// CalculateVUSDAmount calculates vUSD amount from SKILL amount
func (k Keeper) CalculateVUSDAmount(ctx context.Context, skillAmount math.Int) (math.Int, error) {
	params := k.GetParams(ctx)

	// Parse mock price
	price, err := strconv.ParseFloat(params.VusdMockPrice, 64)
	if err != nil {
		return math.ZeroInt(), fmt.Errorf("invalid mock price: %s", params.VusdMockPrice)
	}

	// Simple conversion: SKILL_amount * price = vUSD_amount
	// 1 SKILL = $0.50, so 1 SKILL = 0.5 vUSD
	priceInt := math.NewIntFromUint64(uint64(price * 1000000)) // Convert to micro units
	vUSDAmount := skillAmount.Mul(priceInt).Quo(math.NewInt(1000000))

	return vUSDAmount, nil
}

// CalculateSkillAmount calculates SKILL amount from vUSD amount
func (k Keeper) CalculateSkillAmount(ctx context.Context, vUSDAmount math.Int) (math.Int, error) {
	params := k.GetParams(ctx)

	// Parse mock price
	price, err := strconv.ParseFloat(params.VusdMockPrice, 64)
	if err != nil {
		return math.ZeroInt(), fmt.Errorf("invalid mock price: %s", params.VusdMockPrice)
	}

	// Convert vUSD to SKILL: vUSD_amount / price
	priceInt := math.NewIntFromUint64(uint64(price * 1000000)) // Convert to micro units
	skillAmount := vUSDAmount.Mul(math.NewInt(1000000)).Quo(priceInt)

	return skillAmount, nil
}

// CalculateCollateralRatio calculates collateral ratio for a position
func (k Keeper) CalculateCollateralRatio(ctx context.Context, skillCollateral, vUSDDebt math.Int) (string, error) {
	if vUSDDebt.IsZero() {
		return "0", nil
	}

	// Get collateral value in USD
	collateralValueUSD, err := k.CalculateVUSDAmount(ctx, skillCollateral)
	if err != nil {
		return "", err
	}

	// Collateral ratio = (collateral_value_USD / debt) * 100
	// With price 0.50: 1 SKILL = 0.5 vUSD
	// So if we deposit 1 SKILL and mint 0.5 vUSD, ratio = (0.5 / 0.5) * 100 = 100%
	// For 150% ratio, we need to mint less vUSD or have more collateral
	ratio := collateralValueUSD.Mul(math.NewInt(100)).Quo(vUSDDebt)
	return ratio.String(), nil
}

// CheckMinCollateralRatio checks if collateral ratio meets minimum requirement
func (k Keeper) CheckMinCollateralRatio(ctx context.Context, skillCollateral, vUSDDebt math.Int) error {
	params := k.GetParams(ctx)

	ratio, err := k.CalculateCollateralRatio(ctx, skillCollateral, vUSDDebt)
	if err != nil {
		return err
	}

	minRatio, err := strconv.ParseFloat(params.MinCollateralRatio, 64)
	if err != nil {
		return fmt.Errorf("invalid min collateral ratio: %s", params.MinCollateralRatio)
	}

	currentRatio, err := strconv.ParseFloat(ratio, 64)
	if err != nil {
		return fmt.Errorf("invalid current ratio: %s", ratio)
	}

	// Both ratios are in percentage form
	// minRatio is like 1.50 (represents 150%)
	// currentRatio is already calculated as percentage (150, 200, etc.)
	minRatioPercentage := minRatio * 100 // Convert 1.50 to 150

	if currentRatio < minRatioPercentage {
		return fmt.Errorf("collateral ratio %.0f%% below minimum %.0f%%", currentRatio, minRatioPercentage)
	}

	return nil
}

// ConvertSkillToVUSD converts SKILL tokens to vUSD based on current price
func (k Keeper) ConvertSkillToVUSD(ctx context.Context, skillAmount sdk.Coin) (sdk.Coin, error) {
	if skillAmount.Denom != "skill" {
		return sdk.Coin{}, fmt.Errorf("invalid denom: expected 'skill', got '%s'", skillAmount.Denom)
	}

	vUSDAmount, err := k.CalculateVUSDAmount(ctx, skillAmount.Amount)
	if err != nil {
		return sdk.Coin{}, err
	}

	return sdk.NewCoin(VUSDDenom, vUSDAmount), nil
}

// ConvertVUSDToSkill converts vUSD to SKILL tokens based on current price
func (k Keeper) ConvertVUSDToSkill(ctx context.Context, vUSDAmount sdk.Coin) (sdk.Coin, error) {
	if vUSDAmount.Denom != VUSDDenom {
		return sdk.Coin{}, fmt.Errorf("invalid denom: expected '%s', got '%s'", VUSDDenom, vUSDAmount.Denom)
	}

	skillAmount, err := k.CalculateSkillAmount(ctx, vUSDAmount.Amount)
	if err != nil {
		return sdk.Coin{}, err
	}

	return sdk.NewCoin("skill", skillAmount), nil
}
