package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors" // ✅ bu sarma işlemi için
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors" // ✅ bu eski sabitler için

	"github.com/SkillChainLab/skillchain/x/profile/types"
)

func (k msgServer) CreateProfile(goCtx context.Context, msg *types.MsgCreateProfile) (*types.MsgCreateProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Mevcut mu kontrol et
	_, found := k.GetProfile(ctx, msg.Username)
	if found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "profile already exists for %s", msg.Username)
	}

	profile := types.Profile{
		Creator:  msg.Creator,
		Username: msg.Username,
		Bio:      msg.Bio,
	}

	k.SetProfile(ctx, profile)

	return &types.MsgCreateProfileResponse{}, nil
}
