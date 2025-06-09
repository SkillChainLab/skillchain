package keeper

import (
	"context"
	"strconv"
	"time"

	"skillchain/x/marketplace/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateJobPosting(goCtx context.Context, msg *types.MsgCreateJobPosting) (*types.MsgCreateJobPostingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetJobPosting(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// Convert string fields to proper types
	budgetAmount, err := strconv.ParseUint(msg.BudgetAmount, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid budget amount")
	}

	deadline, err := strconv.ParseInt(msg.Deadline, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid deadline")
	}

	isActive, err := strconv.ParseBool(msg.IsActive)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid isActive value")
	}

	createdAt, err := strconv.ParseInt(msg.CreatedAt, 10, 64)
	if err != nil {
		// Default to current time if not provided or invalid
		createdAt = time.Now().Unix()
	}

	var jobPosting = types.JobPosting{
		Creator:         msg.Creator,
		Index:           msg.Index,
		ClientAddress:   msg.ClientAddress,
		Title:           msg.Title,
		Description:     msg.Description,
		SkillsRequired:  msg.SkillsRequired,
		BudgetAmount:    budgetAmount,
		PaymentCurrency: msg.PaymentCurrency,
		Deadline:        deadline,
		IsActive:        isActive,
		CreatedAt:       createdAt,
	}

	// Debug: Log before setting job posting
	k.Logger().Info("Setting job posting", "index", msg.Index, "title", msg.Title)
	
	k.SetJobPosting(
		ctx,
		jobPosting,
	)
	
	// Debug: Verify it was set
	_, found := k.GetJobPosting(ctx, msg.Index)
	k.Logger().Info("Job posting verification", "index", msg.Index, "found", found)
	
	return &types.MsgCreateJobPostingResponse{}, nil
}

func (k msgServer) UpdateJobPosting(goCtx context.Context, msg *types.MsgUpdateJobPosting) (*types.MsgUpdateJobPostingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetJobPosting(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Convert string fields to proper types
	budgetAmount, err := strconv.ParseUint(msg.BudgetAmount, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid budget amount")
	}

	deadline, err := strconv.ParseInt(msg.Deadline, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid deadline")
	}

	isActive, err := strconv.ParseBool(msg.IsActive)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid isActive value")
	}

	createdAt, err := strconv.ParseInt(msg.CreatedAt, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid createdAt")
	}

	var jobPosting = types.JobPosting{
		Creator:         msg.Creator,
		Index:           msg.Index,
		ClientAddress:   msg.ClientAddress,
		Title:           msg.Title,
		Description:     msg.Description,
		SkillsRequired:  msg.SkillsRequired,
		BudgetAmount:    budgetAmount,
		PaymentCurrency: msg.PaymentCurrency,
		Deadline:        deadline,
		IsActive:        isActive,
		CreatedAt:       createdAt,
	}

	k.SetJobPosting(ctx, jobPosting)

	return &types.MsgUpdateJobPostingResponse{}, nil
}

func (k msgServer) DeleteJobPosting(goCtx context.Context, msg *types.MsgDeleteJobPosting) (*types.MsgDeleteJobPostingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetJobPosting(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveJobPosting(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteJobPostingResponse{}, nil
}
