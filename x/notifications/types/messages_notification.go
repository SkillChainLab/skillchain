package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateNotification{}

func NewMsgCreateNotification(
	creator string,
	index string,
	userAddress string,
	notificationType string,
	title string,
	message string,
	data string,
	isRead bool,
	createdAt uint64,
	priority string,
	sourceModule string,
	sourceAction string,

) *MsgCreateNotification {
	return &MsgCreateNotification{
		Creator:          creator,
		Index:            index,
		UserAddress:      userAddress,
		NotificationType: notificationType,
		Title:            title,
		Message:          message,
		Data:             data,
		IsRead:           isRead,
		CreatedAt:        createdAt,
		Priority:         priority,
		SourceModule:     sourceModule,
		SourceAction:     sourceAction,
	}
}

func (msg *MsgCreateNotification) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateNotification{}

func NewMsgUpdateNotification(
	creator string,
	index string,
	userAddress string,
	notificationType string,
	title string,
	message string,
	data string,
	isRead bool,
	createdAt uint64,
	priority string,
	sourceModule string,
	sourceAction string,

) *MsgUpdateNotification {
	return &MsgUpdateNotification{
		Creator:          creator,
		Index:            index,
		UserAddress:      userAddress,
		NotificationType: notificationType,
		Title:            title,
		Message:          message,
		Data:             data,
		IsRead:           isRead,
		CreatedAt:        createdAt,
		Priority:         priority,
		SourceModule:     sourceModule,
		SourceAction:     sourceAction,
	}
}

func (msg *MsgUpdateNotification) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteNotification{}

func NewMsgDeleteNotification(
	creator string,
	index string,

) *MsgDeleteNotification {
	return &MsgDeleteNotification{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteNotification) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
