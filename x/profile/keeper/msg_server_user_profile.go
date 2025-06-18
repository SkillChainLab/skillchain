package keeper

import (
	"context"

	"skillchain/x/profile/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateUserProfile(goCtx context.Context, msg *types.MsgCreateUserProfile) (*types.MsgCreateUserProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetUserProfile(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// Set default avatar if empty
	avatar := msg.Avatar
	if avatar == "" {
		avatar = "QmZSpwmV3dfwVDVcaJmdga3VVW15SEgXQDY1wiEj8gzpqc" // Default IPFS hash
	}

	var userProfile = types.UserProfile{
		Creator:         msg.Creator,
		Index:           msg.Index,
		Owner:           msg.Owner,
		DisplayName:     msg.DisplayName,
		Bio:             msg.Bio,
		Location:        msg.Location,
		Website:         msg.Website,
		Github:          msg.Github,
		Linkedin:        msg.Linkedin,
		Twitter:         msg.Twitter,
		Avatar:          avatar,
		ReputationScore: msg.ReputationScore,
		CreatedAt:       msg.CreatedAt,
		UpdatedAt:       msg.UpdatedAt,
	}

	k.SetUserProfile(
		ctx,
		userProfile,
	)
	return &types.MsgCreateUserProfileResponse{}, nil
}

func (k msgServer) UpdateUserProfile(goCtx context.Context, msg *types.MsgUpdateUserProfile) (*types.MsgUpdateUserProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUserProfile(
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

	var userProfile = types.UserProfile{
		Creator:         msg.Creator,
		Index:           msg.Index,
		Owner:           msg.Owner,
		DisplayName:     msg.DisplayName,
		Bio:             msg.Bio,
		Location:        msg.Location,
		Website:         msg.Website,
		Github:          msg.Github,
		Linkedin:        msg.Linkedin,
		Twitter:         msg.Twitter,
		Avatar:          msg.Avatar, // Use whatever the user sends, even if empty
		ReputationScore: msg.ReputationScore,
		CreatedAt:       msg.CreatedAt,
		UpdatedAt:       msg.UpdatedAt,
	}

	k.SetUserProfile(ctx, userProfile)

	return &types.MsgUpdateUserProfileResponse{}, nil
}

func (k msgServer) DeleteUserProfile(goCtx context.Context, msg *types.MsgDeleteUserProfile) (*types.MsgDeleteUserProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUserProfile(
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

	k.RemoveUserProfile(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteUserProfileResponse{}, nil
}
