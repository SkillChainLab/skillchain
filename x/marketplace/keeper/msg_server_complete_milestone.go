package keeper

import (
	"context"
	"fmt"
	"time"

	"skillchain/x/marketplace/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CompleteMilestone(goCtx context.Context, msg *types.MsgCompleteMilestone) (*types.MsgCompleteMilestoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the milestone
	milestone, found := k.GetMilestone(ctx, msg.MilestoneId)
	if !found {
		return nil, fmt.Errorf("milestone not found: %s", msg.MilestoneId)
	}

	// Check if milestone is already completed
	if milestone.IsCompleted {
		return nil, fmt.Errorf("milestone already completed")
	}

	// Check if milestone is already paid (shouldn't be possible, but safety check)
	if milestone.IsPaid {
		return nil, fmt.Errorf("milestone already paid")
	}

	// Get the project to verify freelancer
	project, found := k.GetProject(ctx, milestone.ProjectId)
	if !found {
		return nil, fmt.Errorf("project not found: %s", milestone.ProjectId)
	}

	// Verify that the message creator is the project freelancer
	if msg.Creator != project.FreelancerAddress {
		return nil, fmt.Errorf("only project freelancer can complete milestones")
	}

	// Check project is still active
	if project.Status != "active" {
		return nil, fmt.Errorf("project status is %s, cannot complete milestones", project.Status)
	}

	// Verify freelancer reputation is still adequate (safety check)
	currentReputation := k.profileKeeper.CalculateUserReputation(goCtx, project.FreelancerAddress)
	if currentReputation < 50 { // Lower threshold for ongoing work
		return nil, fmt.Errorf("freelancer reputation %d too low to complete work", currentReputation)
	}

	// Validate delivery notes (minimum content requirement)
	if len(msg.DeliveryNotes) < 10 {
		return nil, fmt.Errorf("delivery notes must be at least 10 characters long")
	}

	// Update milestone status
	milestone.IsCompleted = true
	milestone.SubmittedAt = time.Now().Unix()
	milestone.Status = "pending_approval"

	// Add delivery notes to description
	if milestone.Description != "" {
		milestone.Description = fmt.Sprintf("%s\n\n[DELIVERY NOTES]: %s", milestone.Description, msg.DeliveryNotes)
	} else {
		milestone.Description = fmt.Sprintf("[DELIVERY NOTES]: %s", msg.DeliveryNotes)
	}

	k.SetMilestone(ctx, milestone)

	// Check if this is the last milestone for the project
	allMilestones := k.GetAllMilestone(ctx)
	projectMilestones := make([]types.Milestone, 0)
	completedMilestones := 0

	for _, m := range allMilestones {
		if m.ProjectId == milestone.ProjectId {
			projectMilestones = append(projectMilestones, m)
			if m.IsCompleted {
				completedMilestones++
			}
		}
	}

	// Update project status if all milestones completed
	if len(projectMilestones) > 0 && completedMilestones == len(projectMilestones) {
		project.Status = "pending_final_approval"
		k.SetProject(ctx, project)
	}

	// Emit milestone completion event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"milestone_completed",
			sdk.NewAttribute("milestone_id", msg.MilestoneId),
			sdk.NewAttribute("project_id", milestone.ProjectId),
			sdk.NewAttribute("freelancer", project.FreelancerAddress),
			sdk.NewAttribute("client", project.ClientAddress),
			sdk.NewAttribute("amount", fmt.Sprintf("%d", milestone.Amount)),
			sdk.NewAttribute("submitted_at", fmt.Sprintf("%d", milestone.SubmittedAt)),
			sdk.NewAttribute("freelancer_reputation", fmt.Sprintf("%d", currentReputation)),
			sdk.NewAttribute("delivery_notes_length", fmt.Sprintf("%d", len(msg.DeliveryNotes))),
		),
	)

	// Emit notification event for client
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"milestone_awaiting_approval",
			sdk.NewAttribute("milestone_id", msg.MilestoneId),
			sdk.NewAttribute("project_id", milestone.ProjectId),
			sdk.NewAttribute("client", project.ClientAddress),
			sdk.NewAttribute("milestone_title", milestone.Title),
			sdk.NewAttribute("amount", fmt.Sprintf("%d", milestone.Amount)),
			sdk.NewAttribute("due_date", fmt.Sprintf("%d", milestone.DueDate)),
		),
	)

	return &types.MsgCompleteMilestoneResponse{}, nil
}
