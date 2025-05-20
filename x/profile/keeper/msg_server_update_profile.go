package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/SkillChainLab/skillchain/x/profile/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateProfile(goCtx context.Context, msg *types.MsgUpdateProfile) (*types.MsgUpdateProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if profile exists
	profile, found := k.GetProfile(ctx, msg.Username)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrProfileNotFound, "profile %s not found", msg.Username)
	}

	// Check if the creator is the same
	if profile.Creator != msg.Creator {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "only the creator can update the profile")
	}

	// Update profile fields
	profile.Bio = msg.Bio
	profile.Skills = msg.Skills
	profile.Experiences = msg.Experiences
	profile.Website = msg.Website
	profile.Github = msg.Github
	profile.Linkedin = msg.Linkedin
	profile.Twitter = msg.Twitter
	profile.Avatar = msg.Avatar
	profile.Location = msg.Location
	profile.Email = msg.Email

	// Save the updated profile
	k.SetProfile(ctx, profile)

	return &types.MsgUpdateProfileResponse{
		Username: profile.Username,
	}, nil
}
