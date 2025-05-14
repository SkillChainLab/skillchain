package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors" // ✅ bu sarma işlemi için
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors" // ✅ bu eski sabitler için

	"github.com/SkillChainLab/skillchain/x/profile/types"
)

func (k msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if k.GetAuthority() != req.Authority {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid authority; expected %s, got %s", k.GetAuthority(), req.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.SetParams(ctx, req.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}
