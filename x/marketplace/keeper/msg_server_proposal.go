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

func (k msgServer) CreateProposal(goCtx context.Context, msg *types.MsgCreateProposal) (*types.MsgCreateProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetProposal(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// Convert string fields to proper types
	proposedAmount, err := strconv.ParseUint(msg.ProposedAmount, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid proposed amount")
	}

	freelancerReputation, err := strconv.ParseUint(msg.FreelancerReputation, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid freelancer reputation")
	}

	isAccepted, err := strconv.ParseBool(msg.IsAccepted)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid isAccepted value")
	}

	createdAt, err := strconv.ParseInt(msg.CreatedAt, 10, 64)
	if err != nil {
		// Default to current time if not provided or invalid
		createdAt = time.Now().Unix()
	}

	var proposal = types.Proposal{
		Creator:              msg.Creator,
		Index:                msg.Index,
		JobPostingId:         msg.JobPostingId,
		FreelancerAddress:    msg.FreelancerAddress,
		ProposedAmount:       proposedAmount,
		ProposedCurrency:     msg.ProposedCurrency,
		ProposedTimeline:     msg.ProposedTimeline,
		CoverLetter:          msg.CoverLetter,
		FreelancerReputation: freelancerReputation,
		IsAccepted:           isAccepted,
		CreatedAt:            createdAt,
	}

	k.SetProposal(
		ctx,
		proposal,
	)
	return &types.MsgCreateProposalResponse{}, nil
}

func (k msgServer) UpdateProposal(goCtx context.Context, msg *types.MsgUpdateProposal) (*types.MsgUpdateProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetProposal(
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
	proposedAmount, err := strconv.ParseUint(msg.ProposedAmount, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid proposed amount")
	}

	freelancerReputation, err := strconv.ParseUint(msg.FreelancerReputation, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid freelancer reputation")
	}

	isAccepted, err := strconv.ParseBool(msg.IsAccepted)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid isAccepted value")
	}

	createdAt, err := strconv.ParseInt(msg.CreatedAt, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid createdAt")
	}

	var proposal = types.Proposal{
		Creator:              msg.Creator,
		Index:                msg.Index,
		JobPostingId:         msg.JobPostingId,
		FreelancerAddress:    msg.FreelancerAddress,
		ProposedAmount:       proposedAmount,
		ProposedCurrency:     msg.ProposedCurrency,
		ProposedTimeline:     msg.ProposedTimeline,
		CoverLetter:          msg.CoverLetter,
		FreelancerReputation: freelancerReputation,
		IsAccepted:           isAccepted,
		CreatedAt:            createdAt,
	}

	k.SetProposal(ctx, proposal)

	return &types.MsgUpdateProposalResponse{}, nil
}

func (k msgServer) DeleteProposal(goCtx context.Context, msg *types.MsgDeleteProposal) (*types.MsgDeleteProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetProposal(
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

	k.RemoveProposal(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteProposalResponse{}, nil
}
