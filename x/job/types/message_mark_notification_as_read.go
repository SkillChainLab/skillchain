package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMarkNotificationAsRead{}

func NewMsgMarkNotificationAsRead(creator string, notificationId string) *MsgMarkNotificationAsRead {
	return &MsgMarkNotificationAsRead{
		Creator:        creator,
		NotificationId: notificationId,
	}
}

func (msg *MsgMarkNotificationAsRead) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
