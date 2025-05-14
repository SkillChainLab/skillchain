package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/SkillChainLab/skillchain/x/profile/types"
)

func (k msgServer) DeleteProfile(goCtx context.Context, msg *types.MsgDeleteProfile) (*types.MsgDeleteProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	profile, found := k.GetProfile(ctx, msg.Username)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "profile not found")
	}

	if msg.Creator != profile.Creator {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "you are not the owner of this profile")
	}

	k.RemoveProfile(ctx, msg.Username)

	return &types.MsgDeleteProfileResponse{}, nil
}
