package keeper

import (
	"context"
	"fmt"
	"time"

	"skillchain/x/profile/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EndorseSkill(goCtx context.Context, msg *types.MsgEndorseSkill) (*types.MsgEndorseSkillResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate that target user exists
	_, found := k.GetUserProfile(ctx, msg.TargetUser)
	if !found {
		return nil, fmt.Errorf("target user profile not found: %s", msg.TargetUser)
	}

	// Check if user is trying to endorse themselves
	if msg.Creator == msg.TargetUser {
		return nil, fmt.Errorf("users cannot endorse their own skills")
	}

	// Create unique endorsement ID
	endorsementId := fmt.Sprintf("%s-%s-%s", msg.Creator, msg.TargetUser, msg.SkillName)

	// Check if endorsement already exists
	_, found = k.GetSkillEndorsement(ctx, endorsementId)
	if found {
		return nil, fmt.Errorf("endorsement already exists for this skill")
	}

	// Handle token staking if specified
	var stakedTokens uint64 = 0
	if msg.StakeTokens > 0 {
		// Validate user has enough SKILL tokens
		endorserAddr, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
			return nil, fmt.Errorf("invalid endorser address: %v", err)
		}

		// Check balance
		skillDenom := "uskill" // Our native token (micro denomination)
		balance := k.bankKeeper.GetBalance(ctx, endorserAddr, skillDenom)
		requiredAmount := sdk.NewInt64Coin(skillDenom, int64(msg.StakeTokens))

		if balance.IsLT(requiredAmount) {
			return nil, fmt.Errorf("insufficient SKILL tokens: have %s, need %s", balance.String(), requiredAmount.String())
		}

		// Transfer tokens to escrow (profile module account)
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, endorserAddr, types.ModuleName, sdk.NewCoins(requiredAmount))
		if err != nil {
			return nil, fmt.Errorf("failed to stake tokens: %v", err)
		}

		stakedTokens = msg.StakeTokens
	}

	// Create skill endorsement record
	endorsement := types.SkillEndorsement{
		Index:             endorsementId,
		Endorser:          msg.Creator,
		TargetUser:        msg.TargetUser,
		SkillName:         msg.SkillName,
		EndorsementType:   msg.EndorsementType,
		Comment:           msg.Comment,
		CreatedAt:         uint64(time.Now().Unix()),
		SkillTokensStaked: stakedTokens,
	}

	// Store the endorsement
	k.SetSkillEndorsement(ctx, endorsement)

	// Update target user's reputation automatically
	err := k.UpdateUserReputation(ctx, msg.TargetUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update reputation: %v", err)
	}

	// Emit endorsement event with staking info
	eventAttrs := []sdk.Attribute{
		sdk.NewAttribute("endorser", msg.Creator),
		sdk.NewAttribute("target_user", msg.TargetUser),
		sdk.NewAttribute("skill_name", msg.SkillName),
		sdk.NewAttribute("endorsement_type", msg.EndorsementType),
		sdk.NewAttribute("endorsement_id", endorsementId),
		sdk.NewAttribute("staked_tokens", fmt.Sprintf("%d", stakedTokens)),
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent("skill_endorsed", eventAttrs...),
	)

	return &types.MsgEndorseSkillResponse{}, nil
}
