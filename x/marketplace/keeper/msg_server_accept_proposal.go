package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"skillchain/x/marketplace/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AcceptProposal(goCtx context.Context, msg *types.MsgAcceptProposal) (*types.MsgAcceptProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the proposal
	proposal, found := k.GetProposal(ctx, msg.ProposalId)
	if !found {
		return nil, fmt.Errorf("proposal not found: %s", msg.ProposalId)
	}

	// Check if proposal is already accepted
	if proposal.IsAccepted {
		return nil, fmt.Errorf("proposal already accepted")
	}

	// Get the job posting
	jobPosting, found := k.GetJobPosting(ctx, proposal.JobPostingId)
	if !found {
		return nil, fmt.Errorf("job posting not found: %s", proposal.JobPostingId)
	}

	// Verify that the message creator is the job posting client
	if msg.Creator != jobPosting.ClientAddress {
		return nil, fmt.Errorf("only job posting client can accept proposals")
	}

	// Check if job posting is still active
	if !jobPosting.IsActive {
		return nil, fmt.Errorf("job posting is no longer active")
	}

	// Verify freelancer reputation meets minimum requirements
	currentReputation := k.profileKeeper.CalculateUserReputation(goCtx, proposal.FreelancerAddress)
	if currentReputation < 100 { // Minimum reputation requirement
		return nil, fmt.Errorf("freelancer reputation %d below minimum 100", currentReputation)
	}

	// Validate payment currency (must be vUSD for now)
	if proposal.ProposedCurrency != "uvusd" {
		return nil, fmt.Errorf("only vUSD payments supported, got %s", proposal.ProposedCurrency)
	}

	// Use proposed amount directly (now uint64)
	proposedAmount := proposal.ProposedAmount

	// Check client has sufficient vUSD balance for escrow
	clientAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, fmt.Errorf("invalid client address: %w", err)
	}

	clientBalance := k.bankKeeper.GetBalance(ctx, clientAddr, "uvusd")
	requiredAmount := sdk.NewInt64Coin("uvusd", int64(proposedAmount))

	if clientBalance.IsLT(requiredAmount) {
		return nil, fmt.Errorf("insufficient vUSD balance: have %s, need %s",
			clientBalance.String(), requiredAmount.String())
	}

	// Create project ID
	projectId := fmt.Sprintf("%s-%s-%d", proposal.JobPostingId, proposal.FreelancerAddress, time.Now().Unix())

	// Transfer payment to escrow (marketplace module account)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, clientAddr, types.ModuleName, sdk.NewCoins(requiredAmount))
	if err != nil {
		return nil, fmt.Errorf("failed to transfer funds to escrow: %w", err)
	}

	// Mark proposal as accepted
	proposal.IsAccepted = true
	k.SetProposal(ctx, proposal)

	// Create project
	project := types.Project{
		Index:             projectId,
		JobPostingId:      proposal.JobPostingId,
		ProposalId:        msg.ProposalId,
		ClientAddress:     jobPosting.ClientAddress,
		FreelancerAddress: proposal.FreelancerAddress,
		TotalAmount:       proposedAmount,
		PaidAmount:        0,
		EscrowAmount:      proposedAmount,
		Status:            "active",
		StartDate:         time.Now().Unix(),
		ExpectedEndDate:   0, // Will be set based on proposed timeline
		ActualEndDate:     0,
		Creator:           msg.Creator,
	}

	// Parse proposed timeline and set expected end date
	if timelineDays, err := strconv.ParseInt(proposal.ProposedTimeline, 10, 64); err == nil {
		project.ExpectedEndDate = time.Now().Unix() + (timelineDays * 24 * 3600) // Convert days to seconds
	}

	k.SetProject(ctx, project)

	// Deactivate job posting (job filled)
	jobPosting.IsActive = false
	k.SetJobPosting(ctx, jobPosting)

	// Emit project started event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"project_started",
			sdk.NewAttribute("project_id", projectId),
			sdk.NewAttribute("client", jobPosting.ClientAddress),
			sdk.NewAttribute("freelancer", proposal.FreelancerAddress),
			sdk.NewAttribute("amount", fmt.Sprintf("%d", proposal.ProposedAmount)),
			sdk.NewAttribute("currency", proposal.ProposedCurrency),
			sdk.NewAttribute("freelancer_reputation", fmt.Sprintf("%d", currentReputation)),
		),
	)

	return &types.MsgAcceptProposalResponse{}, nil
}
