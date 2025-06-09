package keeper

import (
	"context"
	"fmt"
	"time"

	"skillchain/x/profile/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) WithdrawStakedTokens(goCtx context.Context, msg *types.MsgWithdrawStakedTokens) (*types.MsgWithdrawStakedTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Find user profile by address (need to search through all profiles)
	allProfiles := k.GetAllUserProfile(ctx)
	var userProfile *types.UserProfile

	for _, profile := range allProfiles {
		if profile.Owner == msg.Creator {
			userProfile = &profile
			break
		}
	}

	if userProfile == nil {
		return nil, fmt.Errorf("user profile not found for address: %s", msg.Creator)
	}

	// Calculate total staked tokens for this user and skill
	totalStaked := uint64(0)
	var userEndorsements []types.SkillEndorsement

	allEndorsements := k.GetAllSkillEndorsement(ctx)
	for _, endorsement := range allEndorsements {
		// Find endorsements where this user is the endorser for this specific skill
		if endorsement.Endorser == msg.Creator && endorsement.SkillName == msg.SkillName && endorsement.SkillTokensStaked > 0 {
			totalStaked += endorsement.SkillTokensStaked
			userEndorsements = append(userEndorsements, endorsement)
		}
	}

	// Check if user has any staked tokens for this skill
	if totalStaked == 0 {
		return nil, fmt.Errorf("no staked tokens found for skill '%s'. Either you haven't staked tokens for this skill or they were already slashed in a dispute", msg.SkillName)
	}

	// Transfer staked tokens back to user
	userAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("invalid user address: %v", err)
	}

	withdrawCoin := sdk.NewInt64Coin("uskill", int64(totalStaked))
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddr, sdk.NewCoins(withdrawCoin))
	if err != nil {
		return nil, fmt.Errorf("failed to transfer tokens back to user: %v", err)
	}

	// Update all endorsements to remove staked tokens
	for _, endorsement := range userEndorsements {
		endorsement.SkillTokensStaked = 0
		endorsement.Comment = fmt.Sprintf("[TOKENS WITHDRAWN] %s", endorsement.Comment)
		k.SetSkillEndorsement(ctx, endorsement)
	}

	// Recalculate user's reputation (personal staking bonus will be lost)
	err = k.UpdateUserReputation(ctx, msg.Creator)
	if err != nil {
		// Log error but don't fail the transaction
		ctx.Logger().Error("Failed to update user reputation after token withdrawal", "error", err)
	}

	// Also recalculate reputation for all target users affected by these endorsements
	targetUsers := make(map[string]bool)
	for _, endorsement := range userEndorsements {
		if !targetUsers[endorsement.TargetUser] {
			targetUsers[endorsement.TargetUser] = true
			err = k.UpdateUserReputation(ctx, endorsement.TargetUser)
			if err != nil {
				ctx.Logger().Error("Failed to update target user reputation after withdrawal", "error", err)
			}
		}
	}

	// Update user profile timestamp
	userProfile.UpdatedAt = uint64(time.Now().Unix())
	k.SetUserProfile(ctx, *userProfile)

	// Emit withdrawal event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"staked_tokens_withdrawn",
			sdk.NewAttribute("user", msg.Creator),
			sdk.NewAttribute("skill_name", msg.SkillName),
			sdk.NewAttribute("withdrawn_tokens", fmt.Sprintf("%d", totalStaked)),
			sdk.NewAttribute("endorsements_affected", fmt.Sprintf("%d", len(userEndorsements))),
		),
	)

	return &types.MsgWithdrawStakedTokensResponse{}, nil
}
