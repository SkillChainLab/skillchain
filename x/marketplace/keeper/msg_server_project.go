package keeper

import (
	"context"
	"strconv"

	"skillchain/x/marketplace/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateProject(goCtx context.Context, msg *types.MsgCreateProject) (*types.MsgCreateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetProject(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// Convert string fields to proper types
	totalAmount, err := strconv.ParseUint(msg.TotalAmount, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid total amount")
	}

	paidAmount, err := strconv.ParseUint(msg.PaidAmount, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid paid amount")
	}

	escrowAmount, err := strconv.ParseUint(msg.EscrowAmount, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid escrow amount")
	}

	startDate, err := strconv.ParseInt(msg.StartDate, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid start date")
	}

	expectedEndDate, err := strconv.ParseInt(msg.ExpectedEndDate, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid expected end date")
	}

	actualEndDate, err := strconv.ParseInt(msg.ActualEndDate, 10, 64)
	if err != nil {
		// Default to 0 if not provided or invalid
		actualEndDate = 0
	}

	var project = types.Project{
		Creator:           msg.Creator,
		Index:             msg.Index,
		JobPostingId:      msg.JobPostingId,
		ProposalId:        msg.ProposalId,
		ClientAddress:     msg.ClientAddress,
		FreelancerAddress: msg.FreelancerAddress,
		TotalAmount:       totalAmount,
		PaidAmount:        paidAmount,
		EscrowAmount:      escrowAmount,
		Status:            msg.Status,
		StartDate:         startDate,
		ExpectedEndDate:   expectedEndDate,
		ActualEndDate:     actualEndDate,
	}

	k.SetProject(
		ctx,
		project,
	)
	return &types.MsgCreateProjectResponse{}, nil
}

func (k msgServer) UpdateProject(goCtx context.Context, msg *types.MsgUpdateProject) (*types.MsgUpdateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetProject(
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
	totalAmount, err := strconv.ParseUint(msg.TotalAmount, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid total amount")
	}

	paidAmount, err := strconv.ParseUint(msg.PaidAmount, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid paid amount")
	}

	escrowAmount, err := strconv.ParseUint(msg.EscrowAmount, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid escrow amount")
	}

	startDate, err := strconv.ParseInt(msg.StartDate, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid start date")
	}

	expectedEndDate, err := strconv.ParseInt(msg.ExpectedEndDate, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid expected end date")
	}

	actualEndDate, err := strconv.ParseInt(msg.ActualEndDate, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid actual end date")
	}

	var project = types.Project{
		Creator:           msg.Creator,
		Index:             msg.Index,
		JobPostingId:      msg.JobPostingId,
		ProposalId:        msg.ProposalId,
		ClientAddress:     msg.ClientAddress,
		FreelancerAddress: msg.FreelancerAddress,
		TotalAmount:       totalAmount,
		PaidAmount:        paidAmount,
		EscrowAmount:      escrowAmount,
		Status:            msg.Status,
		StartDate:         startDate,
		ExpectedEndDate:   expectedEndDate,
		ActualEndDate:     actualEndDate,
	}

	k.SetProject(ctx, project)

	return &types.MsgUpdateProjectResponse{}, nil
}

func (k msgServer) DeleteProject(goCtx context.Context, msg *types.MsgDeleteProject) (*types.MsgDeleteProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetProject(
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

	k.RemoveProject(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteProjectResponse{}, nil
}
