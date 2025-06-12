package keeper

import (
	"context"

	"skillchain/x/profile/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateUserSkill(goCtx context.Context, msg *types.MsgCreateUserSkill) (*types.MsgCreateUserSkillResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Authorization check - only the skill owner can create skills for themselves
	if msg.Creator != msg.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only skill owner can create skills for themselves")
	}

	// Check if the value already exists
	_, isFound := k.GetUserSkill(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var userSkill = types.UserSkill{
		Creator:          msg.Creator,
		Index:            msg.Index,
		Owner:            msg.Owner,
		SkillName:        msg.SkillName,
		ProficiencyLevel: msg.ProficiencyLevel,
		YearsExperience:  msg.YearsExperience,
		Verified:         msg.Verified,
		VerifiedBy:       msg.VerifiedBy,
		VerificationDate: msg.VerificationDate,
		EndorsementCount: msg.EndorsementCount,
	}

	k.SetUserSkill(
		ctx,
		userSkill,
	)
	return &types.MsgCreateUserSkillResponse{}, nil
}

func (k msgServer) UpdateUserSkill(goCtx context.Context, msg *types.MsgUpdateUserSkill) (*types.MsgUpdateUserSkillResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUserSkill(
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

	var userSkill = types.UserSkill{
		Creator:          msg.Creator,
		Index:            msg.Index,
		Owner:            msg.Owner,
		SkillName:        msg.SkillName,
		ProficiencyLevel: msg.ProficiencyLevel,
		YearsExperience:  msg.YearsExperience,
		Verified:         msg.Verified,
		VerifiedBy:       msg.VerifiedBy,
		VerificationDate: msg.VerificationDate,
		EndorsementCount: msg.EndorsementCount,
	}

	k.SetUserSkill(ctx, userSkill)

	return &types.MsgUpdateUserSkillResponse{}, nil
}

func (k msgServer) DeleteUserSkill(goCtx context.Context, msg *types.MsgDeleteUserSkill) (*types.MsgDeleteUserSkillResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUserSkill(
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

	k.RemoveUserSkill(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteUserSkillResponse{}, nil
}
