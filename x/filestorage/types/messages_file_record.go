package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateFileRecord{}

func NewMsgCreateFileRecord(
	creator string,
	index string,
	owner string,
	filename string,
	fileHash string,
	fileSize uint64,
	contentType string,
	uploadDate uint64,
	ipfsHash string,
	metadata string,
	isPublic bool,

) *MsgCreateFileRecord {
	return &MsgCreateFileRecord{
		Creator:     creator,
		Index:       index,
		Owner:       owner,
		Filename:    filename,
		FileHash:    fileHash,
		FileSize:    fileSize,
		ContentType: contentType,
		UploadDate:  uploadDate,
		IpfsHash:    ipfsHash,
		Metadata:    metadata,
		IsPublic:    isPublic,
	}
}

func (msg *MsgCreateFileRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateFileRecord{}

func NewMsgUpdateFileRecord(
	creator string,
	index string,
	owner string,
	filename string,
	fileHash string,
	fileSize uint64,
	contentType string,
	uploadDate uint64,
	ipfsHash string,
	metadata string,
	isPublic bool,

) *MsgUpdateFileRecord {
	return &MsgUpdateFileRecord{
		Creator:     creator,
		Index:       index,
		Owner:       owner,
		Filename:    filename,
		FileHash:    fileHash,
		FileSize:    fileSize,
		ContentType: contentType,
		UploadDate:  uploadDate,
		IpfsHash:    ipfsHash,
		Metadata:    metadata,
		IsPublic:    isPublic,
	}
}

func (msg *MsgUpdateFileRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteFileRecord{}

func NewMsgDeleteFileRecord(
	creator string,
	index string,

) *MsgDeleteFileRecord {
	return &MsgDeleteFileRecord{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteFileRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
