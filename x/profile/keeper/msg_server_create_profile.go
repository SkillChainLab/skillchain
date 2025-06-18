package keeper

import (
	"context"
	"time"

	"skillchain/x/profile/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateProfile(goCtx context.Context, msg *types.MsgCreateProfile) (*types.MsgCreateProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if profile already exists
	_, found := k.GetUserProfile(ctx, msg.Creator)
	if found {
		return nil, types.ErrProfileAlreadyExists
	}

	// Create new user profile
	avatar := msg.Avatar
	if avatar == "" {
		avatar = "QmZSpwmV3dfwVDVcaJmdga3VVW15SEgXQDY1wiEj8gzpqc" // Default IPFS hash
	}

	profile := types.UserProfile{
		Index:           msg.Creator, // Use creator address as profile ID
		Owner:           msg.Creator,
		DisplayName:     msg.DisplayName,
		Bio:             msg.Bio,
		Location:        msg.Location,
		Website:         msg.Website,
		Github:          msg.Github,
		Linkedin:        msg.Linkedin,
		Twitter:         msg.Twitter,
		Avatar:          avatar,
		ReputationScore: 100, // Starting reputation score
		CreatedAt:       uint64(time.Now().Unix()),
		UpdatedAt:       uint64(time.Now().Unix()),
	}

	// Store the profile
	k.SetUserProfile(ctx, profile)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"profile_created",
			sdk.NewAttribute("creator", msg.Creator),
			sdk.NewAttribute("display_name", msg.DisplayName),
		),
	)

	return &types.MsgCreateProfileResponse{}, nil
}
