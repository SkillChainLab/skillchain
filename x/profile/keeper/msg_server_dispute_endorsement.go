package keeper

import (
	"context"
	"fmt"
	"time"

	"skillchain/x/profile/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DisputeEndorsement(goCtx context.Context, msg *types.MsgDisputeEndorsement) (*types.MsgDisputeEndorsementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the disputed endorsement
	endorsement, found := k.GetSkillEndorsement(ctx, msg.EndorsementId)
	if !found {
		return nil, fmt.Errorf("endorsement not found: %s", msg.EndorsementId)
	}

	// Verify the disputer is authorized (could be the target user who hired the endorser)
	// In this case, the target user (who received the endorsement) is disputing
	// because the endorser (freelancer) failed to deliver the job properly
	if msg.Creator != endorsement.TargetUser {
		return nil, fmt.Errorf("only the target user can dispute this endorsement")
	}

	// Check if there are staked tokens to slash
	if endorsement.SkillTokensStaked == 0 {
		return nil, fmt.Errorf("no staked tokens to slash for this endorsement")
	}

	// Transfer staked tokens to the disputer (client who was wronged)
	clientAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("invalid client address: %v", err)
	}

	// Calculate slash amount (could be partial or full)
	slashAmount := endorsement.SkillTokensStaked // Full slash for now
	slashCoin := sdk.NewInt64Coin("uskill", int64(slashAmount))

	// Transfer from profile module to client
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, clientAddr, sdk.NewCoins(slashCoin))
	if err != nil {
		return nil, fmt.Errorf("failed to transfer slashed tokens: %v", err)
	}

	// Apply reputation penalty to the endorser (freelancer who failed)
	endorserProfile, found := k.GetUserProfile(ctx, endorsement.Endorser)
	if found {
		// Apply -10 reputation penalty
		if endorserProfile.ReputationScore >= 10 {
			endorserProfile.ReputationScore -= 10
		} else {
			endorserProfile.ReputationScore = 0 // Don't go below 0
		}
		endorserProfile.UpdatedAt = uint64(time.Now().Unix())
		k.SetUserProfile(ctx, endorserProfile)
	}

	// Mark endorsement as disputed/slashed
	endorsement.SkillTokensStaked = 0 // Tokens have been slashed
	endorsement.Comment = fmt.Sprintf("[DISPUTED] %s | Original: %s", msg.DisputeReason, endorsement.Comment)
	k.SetSkillEndorsement(ctx, endorsement)

	// CRITICAL: Recalculate endorser's reputation (includes lost staking bonus)
	err = k.UpdateUserReputation(ctx, endorsement.Endorser)
	if err != nil {
		// Log error but don't fail the transaction
		ctx.Logger().Error("Failed to update endorser reputation after dispute", "error", err)
	}

	// Recalculate target user's reputation (since endorsement value changed)
	err = k.UpdateUserReputation(ctx, endorsement.TargetUser)
	if err != nil {
		// Log error but don't fail the transaction
		ctx.Logger().Error("Failed to update target user reputation after dispute", "error", err)
	}

	// Emit dispute event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"endorsement_disputed",
			sdk.NewAttribute("disputer", msg.Creator),
			sdk.NewAttribute("endorsement_id", msg.EndorsementId),
			sdk.NewAttribute("slashed_tokens", fmt.Sprintf("%d", slashAmount)),
			sdk.NewAttribute("endorser_penalized", endorsement.Endorser),
			sdk.NewAttribute("reputation_penalty", "10"),
			sdk.NewAttribute("dispute_reason", msg.DisputeReason),
		),
	)

	return &types.MsgDisputeEndorsementResponse{}, nil
}
