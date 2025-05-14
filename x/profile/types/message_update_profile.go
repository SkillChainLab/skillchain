package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateProfile{}

func NewMsgUpdateProfile(creator string, username string, bio string) *MsgUpdateProfile {
	return &MsgUpdateProfile{
		Creator:  creator,
		Username: username,
		Bio:      bio,
	}
}

func (msg *MsgUpdateProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
