package keeper

import (
	"context"
	"strconv"

	"skillchain/x/marketplace/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMilestone(goCtx context.Context, msg *types.MsgCreateMilestone) (*types.MsgCreateMilestoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetMilestone(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// Convert string fields to proper types
	amount, err := strconv.ParseUint(msg.Amount, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}

	dueDate, err := strconv.ParseInt(msg.DueDate, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid due date")
	}

	isCompleted, err := strconv.ParseBool(msg.IsCompleted)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid isCompleted value")
	}

	isPaid, err := strconv.ParseBool(msg.IsPaid)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid isPaid value")
	}

	submittedAt, err := strconv.ParseInt(msg.SubmittedAt, 10, 64)
	if err != nil {
		// Default to 0 if not provided or invalid
		submittedAt = 0
	}

	approvedAt, err := strconv.ParseInt(msg.ApprovedAt, 10, 64)
	if err != nil {
		// Default to 0 if not provided or invalid
		approvedAt = 0
	}

	var milestone = types.Milestone{
		Creator:     msg.Creator,
		Index:       msg.Index,
		ProjectId:   msg.ProjectId,
		Title:       msg.Title,
		Description: msg.Description,
		Amount:      amount,
		DueDate:     dueDate,
		Status:      msg.Status,
		IsCompleted: isCompleted,
		IsPaid:      isPaid,
		SubmittedAt: submittedAt,
		ApprovedAt:  approvedAt,
	}

	k.SetMilestone(
		ctx,
		milestone,
	)
	return &types.MsgCreateMilestoneResponse{}, nil
}

func (k msgServer) UpdateMilestone(goCtx context.Context, msg *types.MsgUpdateMilestone) (*types.MsgUpdateMilestoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMilestone(
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
	amount, err := strconv.ParseUint(msg.Amount, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}

	dueDate, err := strconv.ParseInt(msg.DueDate, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid due date")
	}

	isCompleted, err := strconv.ParseBool(msg.IsCompleted)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid isCompleted value")
	}

	isPaid, err := strconv.ParseBool(msg.IsPaid)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid isPaid value")
	}

	submittedAt, err := strconv.ParseInt(msg.SubmittedAt, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid submittedAt")
	}

	approvedAt, err := strconv.ParseInt(msg.ApprovedAt, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid approvedAt")
	}

	var milestone = types.Milestone{
		Creator:     msg.Creator,
		Index:       msg.Index,
		ProjectId:   msg.ProjectId,
		Title:       msg.Title,
		Description: msg.Description,
		Amount:      amount,
		DueDate:     dueDate,
		Status:      msg.Status,
		IsCompleted: isCompleted,
		IsPaid:      isPaid,
		SubmittedAt: submittedAt,
		ApprovedAt:  approvedAt,
	}

	k.SetMilestone(ctx, milestone)

	return &types.MsgUpdateMilestoneResponse{}, nil
}

func (k msgServer) DeleteMilestone(goCtx context.Context, msg *types.MsgDeleteMilestone) (*types.MsgDeleteMilestoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMilestone(
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

	k.RemoveMilestone(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteMilestoneResponse{}, nil
}
