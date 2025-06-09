package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateUserActivity{}

func NewMsgCreateUserActivity(
	creator string,
	index string,
	userAddress string,
	activityType string,
	action string,
	resourceId string,
	timestamp uint64,
	ipAddress string,
	userAgent string,
	metadata string,

) *MsgCreateUserActivity {
	return &MsgCreateUserActivity{
		Creator:      creator,
		Index:        index,
		UserAddress:  userAddress,
		ActivityType: activityType,
		Action:       action,
		ResourceId:   resourceId,
		Timestamp:    timestamp,
		IpAddress:    ipAddress,
		UserAgent:    userAgent,
		Metadata:     metadata,
	}
}

func (msg *MsgCreateUserActivity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateUserActivity{}

func NewMsgUpdateUserActivity(
	creator string,
	index string,
	userAddress string,
	activityType string,
	action string,
	resourceId string,
	timestamp uint64,
	ipAddress string,
	userAgent string,
	metadata string,

) *MsgUpdateUserActivity {
	return &MsgUpdateUserActivity{
		Creator:      creator,
		Index:        index,
		UserAddress:  userAddress,
		ActivityType: activityType,
		Action:       action,
		ResourceId:   resourceId,
		Timestamp:    timestamp,
		IpAddress:    ipAddress,
		UserAgent:    userAgent,
		Metadata:     metadata,
	}
}

func (msg *MsgUpdateUserActivity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteUserActivity{}

func NewMsgDeleteUserActivity(
	creator string,
	index string,

) *MsgDeleteUserActivity {
	return &MsgDeleteUserActivity{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteUserActivity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
