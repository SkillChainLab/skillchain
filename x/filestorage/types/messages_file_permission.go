package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateFilePermission{}

func NewMsgCreateFilePermission(
	creator string,
	index string,
	fileId string,
	userAddress string,
	permissionLevel string,
	grantedBy string,
	grantedAt uint64,
	expiresAt uint64,

) *MsgCreateFilePermission {
	return &MsgCreateFilePermission{
		Creator:         creator,
		Index:           index,
		FileId:          fileId,
		UserAddress:     userAddress,
		PermissionLevel: permissionLevel,
		GrantedBy:       grantedBy,
		GrantedAt:       grantedAt,
		ExpiresAt:       expiresAt,
	}
}

func (msg *MsgCreateFilePermission) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateFilePermission{}

func NewMsgUpdateFilePermission(
	creator string,
	index string,
	fileId string,
	userAddress string,
	permissionLevel string,
	grantedBy string,
	grantedAt uint64,
	expiresAt uint64,

) *MsgUpdateFilePermission {
	return &MsgUpdateFilePermission{
		Creator:         creator,
		Index:           index,
		FileId:          fileId,
		UserAddress:     userAddress,
		PermissionLevel: permissionLevel,
		GrantedBy:       grantedBy,
		GrantedAt:       grantedAt,
		ExpiresAt:       expiresAt,
	}
}

func (msg *MsgUpdateFilePermission) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteFilePermission{}

func NewMsgDeleteFilePermission(
	creator string,
	index string,

) *MsgDeleteFilePermission {
	return &MsgDeleteFilePermission{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteFilePermission) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
