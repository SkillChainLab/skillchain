package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/SkillChainLab/skillchain/x/profile/types"
)

func (k msgServer) UpdateProfile(goCtx context.Context, msg *types.MsgUpdateProfile) (*types.MsgUpdateProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	profile, found := k.GetProfile(ctx, msg.Username)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "profile not found")
	}

	if msg.Creator != profile.Creator {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "you are not the owner of this profile")
	}

	// Güncelleme işlemi
	profile.Bio = msg.Bio
	// İleride username de güncellenecekse burada dikkatli olman gerekir.

	k.SetProfile(ctx, profile)

	return &types.MsgUpdateProfileResponse{}, nil
}
