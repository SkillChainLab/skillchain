package keeper

import (
	"context"

	"skillchain/x/profile/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSkillEndorsement(goCtx context.Context, msg *types.MsgCreateSkillEndorsement) (*types.MsgCreateSkillEndorsementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetSkillEndorsement(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var skillEndorsement = types.SkillEndorsement{
		Creator:           msg.Creator,
		Index:             msg.Index,
		Endorser:          msg.Endorser,
		TargetUser:        msg.TargetUser,
		SkillName:         msg.SkillName,
		EndorsementType:   msg.EndorsementType,
		Comment:           msg.Comment,
		CreatedAt:         msg.CreatedAt,
		SkillTokensStaked: msg.SkillTokensStaked,
	}

	k.SetSkillEndorsement(
		ctx,
		skillEndorsement,
	)
	return &types.MsgCreateSkillEndorsementResponse{}, nil
}

func (k msgServer) UpdateSkillEndorsement(goCtx context.Context, msg *types.MsgUpdateSkillEndorsement) (*types.MsgUpdateSkillEndorsementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSkillEndorsement(
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

	var skillEndorsement = types.SkillEndorsement{
		Creator:           msg.Creator,
		Index:             msg.Index,
		Endorser:          msg.Endorser,
		TargetUser:        msg.TargetUser,
		SkillName:         msg.SkillName,
		EndorsementType:   msg.EndorsementType,
		Comment:           msg.Comment,
		CreatedAt:         msg.CreatedAt,
		SkillTokensStaked: msg.SkillTokensStaked,
	}

	k.SetSkillEndorsement(ctx, skillEndorsement)

	return &types.MsgUpdateSkillEndorsementResponse{}, nil
}

func (k msgServer) DeleteSkillEndorsement(goCtx context.Context, msg *types.MsgDeleteSkillEndorsement) (*types.MsgDeleteSkillEndorsementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSkillEndorsement(
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

	k.RemoveSkillEndorsement(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteSkillEndorsementResponse{}, nil
}
