package keeper

import (
	"context"
	"fmt"
	"time"

	"skillchain/x/marketplace/types"
	profiletypes "skillchain/x/profile/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DisputeProject(goCtx context.Context, msg *types.MsgDisputeProject) (*types.MsgDisputeProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the project
	project, found := k.GetProject(ctx, msg.ProjectId)
	if !found {
		return nil, fmt.Errorf("project not found: %s", msg.ProjectId)
	}

	// Verify that the message creator is the project client
	if msg.Creator != project.ClientAddress {
		return nil, fmt.Errorf("only project client can dispute a project")
	}

	// Check if project is still active (can't dispute completed projects)
	if project.Status == "completed" || project.Status == "disputed" {
		return nil, fmt.Errorf("cannot dispute project with status: %s", project.Status)
	}

	// Update project status to disputed
	project.Status = "disputed"
	k.SetProject(ctx, project)

	// Apply reputation penalty to freelancer (-10 points)
	reputationPenalty := uint64(10)

	// Find freelancer profile by address
	allProfiles := k.profileKeeper.GetAllUserProfile(goCtx)
	var freelancerProfile *profiletypes.UserProfile

	for _, profile := range allProfiles {
		if profile.Owner == project.FreelancerAddress {
			freelancerProfile = &profile
			break
		}
	}

	if freelancerProfile != nil {
		// Get current reputation (already uint64)
		currentReputation := freelancerProfile.ReputationScore

		// Apply penalty (ensure it doesn't go below 0)
		var newReputation uint64
		if currentReputation > reputationPenalty {
			newReputation = currentReputation - reputationPenalty
		} else {
			newReputation = 0
		}

		freelancerProfile.ReputationScore = newReputation
		freelancerProfile.UpdatedAt = uint64(time.Now().Unix())
		k.profileKeeper.SetUserProfile(goCtx, *freelancerProfile)
	}

	// Return escrow amount to client as compensation
	if project.EscrowAmount > 0 {
		clientAddr, err := sdk.AccAddressFromBech32(project.ClientAddress)
		if err != nil {
			return nil, fmt.Errorf("invalid client address: %w", err)
		}

		compensationCoin := sdk.NewInt64Coin("uvusd", int64(project.EscrowAmount))
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, clientAddr, sdk.NewCoins(compensationCoin))
		if err != nil {
			return nil, fmt.Errorf("failed to transfer compensation to client: %w", err)
		}

		// Update project escrow (set to 0 since it's returned)
		project.EscrowAmount = 0
		k.SetProject(ctx, project)
	}

	// Emit dispute event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"project_disputed",
			sdk.NewAttribute("project_id", msg.ProjectId),
			sdk.NewAttribute("client", project.ClientAddress),
			sdk.NewAttribute("freelancer", project.FreelancerAddress),
			sdk.NewAttribute("reason", msg.Reason),
			sdk.NewAttribute("evidence", msg.Evidence),
			sdk.NewAttribute("compensation_amount", fmt.Sprintf("%d", project.EscrowAmount)),
			sdk.NewAttribute("reputation_penalty", fmt.Sprintf("%d", reputationPenalty)),
			sdk.NewAttribute("disputed_at", fmt.Sprintf("%d", time.Now().Unix())),
		),
	)

	return &types.MsgDisputeProjectResponse{}, nil
}
