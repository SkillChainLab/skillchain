package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateUserProfile{}

func NewMsgCreateUserProfile(
	creator string,
	index string,
	owner string,
	displayName string,
	bio string,
	location string,
	website string,
	github string,
	linkedin string,
	twitter string,
	reputationScore uint64,
	createdAt uint64,
	updatedAt uint64,

) *MsgCreateUserProfile {
	return &MsgCreateUserProfile{
		Creator:         creator,
		Index:           index,
		Owner:           owner,
		DisplayName:     displayName,
		Bio:             bio,
		Location:        location,
		Website:         website,
		Github:          github,
		Linkedin:        linkedin,
		Twitter:         twitter,
		ReputationScore: reputationScore,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}

func (msg *MsgCreateUserProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateUserProfile{}

func NewMsgUpdateUserProfile(
	creator string,
	index string,
	owner string,
	displayName string,
	bio string,
	location string,
	website string,
	github string,
	linkedin string,
	twitter string,
	reputationScore uint64,
	createdAt uint64,
	updatedAt uint64,

) *MsgUpdateUserProfile {
	return &MsgUpdateUserProfile{
		Creator:         creator,
		Index:           index,
		Owner:           owner,
		DisplayName:     displayName,
		Bio:             bio,
		Location:        location,
		Website:         website,
		Github:          github,
		Linkedin:        linkedin,
		Twitter:         twitter,
		ReputationScore: reputationScore,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}

func (msg *MsgUpdateUserProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteUserProfile{}

func NewMsgDeleteUserProfile(
	creator string,
	index string,

) *MsgDeleteUserProfile {
	return &MsgDeleteUserProfile{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteUserProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
