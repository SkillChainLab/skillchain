package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateNotificationSettings{}

func NewMsgCreateNotificationSettings(
	creator string,
	index string,
	userAddress string,
	emailEnabled bool,
	pushEnabled bool,
	smsEnabled bool,
	emailAddress string,
	phoneNumber string,
	notificationTypes []string,
	frequency string,

) *MsgCreateNotificationSettings {
	return &MsgCreateNotificationSettings{
		Creator:           creator,
		Index:             index,
		UserAddress:       userAddress,
		EmailEnabled:      emailEnabled,
		PushEnabled:       pushEnabled,
		SmsEnabled:        smsEnabled,
		EmailAddress:      emailAddress,
		PhoneNumber:       phoneNumber,
		NotificationTypes: notificationTypes,
		Frequency:         frequency,
	}
}

func (msg *MsgCreateNotificationSettings) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateNotificationSettings{}

func NewMsgUpdateNotificationSettings(
	creator string,
	index string,
	userAddress string,
	emailEnabled bool,
	pushEnabled bool,
	smsEnabled bool,
	emailAddress string,
	phoneNumber string,
	notificationTypes []string,
	frequency string,

) *MsgUpdateNotificationSettings {
	return &MsgUpdateNotificationSettings{
		Creator:           creator,
		Index:             index,
		UserAddress:       userAddress,
		EmailEnabled:      emailEnabled,
		PushEnabled:       pushEnabled,
		SmsEnabled:        smsEnabled,
		EmailAddress:      emailAddress,
		PhoneNumber:       phoneNumber,
		NotificationTypes: notificationTypes,
		Frequency:         frequency,
	}
}

func (msg *MsgUpdateNotificationSettings) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteNotificationSettings{}

func NewMsgDeleteNotificationSettings(
	creator string,
	index string,

) *MsgDeleteNotificationSettings {
	return &MsgDeleteNotificationSettings{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteNotificationSettings) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
