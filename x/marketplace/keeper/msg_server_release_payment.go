package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"skillchain/x/marketplace/types"
	profiletypes "skillchain/x/profile/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ReleasePayment(goCtx context.Context, msg *types.MsgReleasePayment) (*types.MsgReleasePaymentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the milestone
	milestone, found := k.GetMilestone(ctx, msg.MilestoneId)
	if !found {
		return nil, fmt.Errorf("milestone not found: %s", msg.MilestoneId)
	}

	// Check if milestone is already paid
	if milestone.IsPaid {
		return nil, fmt.Errorf("milestone payment already released")
	}

	// Check if milestone is completed
	if !milestone.IsCompleted {
		return nil, fmt.Errorf("milestone must be completed before payment release")
	}

	// Get the project
	project, found := k.GetProject(ctx, milestone.ProjectId)
	if !found {
		return nil, fmt.Errorf("project not found: %s", milestone.ProjectId)
	}

	// Verify that the message creator is the project client
	if msg.Creator != project.ClientAddress {
		return nil, fmt.Errorf("only project client can release payments")
	}

	// Validate rating (1-5 scale)
	rating, err := strconv.ParseUint(msg.Rating, 10, 64)
	if err != nil || rating < 1 || rating > 5 {
		return nil, fmt.Errorf("rating must be between 1 and 5, got: %s", msg.Rating)
	}

	// Use milestone amount directly (now uint64)
	milestoneAmount := milestone.Amount

	// Transfer vUSD from escrow (marketplace module) to freelancer
	freelancerAddr, err := sdk.AccAddressFromBech32(project.FreelancerAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid freelancer address: %w", err)
	}

	paymentCoin := sdk.NewInt64Coin("uvusd", int64(milestoneAmount))
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, freelancerAddr, sdk.NewCoins(paymentCoin))
	if err != nil {
		return nil, fmt.Errorf("failed to transfer payment to freelancer: %w", err)
	}

	// Update milestone status
	milestone.IsPaid = true
	milestone.ApprovedAt = time.Now().Unix()
	k.SetMilestone(ctx, milestone)

	// Update project payment tracking
	currentPaidAmount := project.PaidAmount
	currentEscrowAmount := project.EscrowAmount

	newPaidAmount := currentPaidAmount + milestoneAmount
	newEscrowAmount := currentEscrowAmount - milestoneAmount

	project.PaidAmount = newPaidAmount
	project.EscrowAmount = newEscrowAmount

	// Check if project is fully paid
	totalAmount := project.TotalAmount
	if newPaidAmount >= totalAmount {
		project.Status = "completed"
		project.ActualEndDate = time.Now().Unix()
	}

	k.SetProject(ctx, project)

	// Calculate reputation boost based on rating
	var reputationBoost uint64
	switch rating {
	case 5: // Excellent
		reputationBoost = 25
	case 4: // Very Good
		reputationBoost = 20
	case 3: // Good
		reputationBoost = 15
	case 2: // Fair
		reputationBoost = 10
	case 1: // Poor
		reputationBoost = 5
	}

	// Apply reputation boost to freelancer
	// Find profile by address (need to search through all profiles)
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

		// Add reputation boost directly
		newReputation := currentReputation + reputationBoost
		freelancerProfile.ReputationScore = newReputation
		freelancerProfile.UpdatedAt = uint64(time.Now().Unix())

		// Save updated profile through profileKeeper using correct index
		k.profileKeeper.SetUserProfile(goCtx, *freelancerProfile)
	}

	// Update freelancer's overall reputation calculation
	err = k.profileKeeper.UpdateUserReputation(goCtx, project.FreelancerAddress)
	if err != nil {
		// Log error but don't fail the payment
		ctx.Logger().Error("Failed to update freelancer reputation after payment", "error", err)
	}

	// Emit payment released event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"payment_released",
			sdk.NewAttribute("milestone_id", msg.MilestoneId),
			sdk.NewAttribute("project_id", milestone.ProjectId),
			sdk.NewAttribute("freelancer", project.FreelancerAddress),
			sdk.NewAttribute("client", project.ClientAddress),
			sdk.NewAttribute("amount", fmt.Sprintf("%d", milestoneAmount)),
			sdk.NewAttribute("currency", "uvusd"),
			sdk.NewAttribute("rating", msg.Rating),
			sdk.NewAttribute("reputation_boost", fmt.Sprintf("%d", reputationBoost)),
			sdk.NewAttribute("project_status", project.Status),
		),
	)

	// Emit project completion event if fully paid
	if project.Status == "completed" {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"project_completed",
				sdk.NewAttribute("project_id", project.Index),
				sdk.NewAttribute("client", project.ClientAddress),
				sdk.NewAttribute("freelancer", project.FreelancerAddress),
				sdk.NewAttribute("total_amount", fmt.Sprintf("%d", project.TotalAmount)),
				sdk.NewAttribute("final_rating", msg.Rating),
				sdk.NewAttribute("feedback", msg.Feedback),
			),
		)
	}

	return &types.MsgReleasePaymentResponse{}, nil
}
