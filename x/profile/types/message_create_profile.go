package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateProfile{}

func NewMsgCreateProfile(creator string, username string, bio string, skills []string, experiences []*Experience, website string, github string, linkedin string, twitter string, avatar string, location string, email string) *MsgCreateProfile {
	return &MsgCreateProfile{
		Creator:     creator,
		Username:    username,
		Bio:         bio,
		Skills:      skills,
		Experiences: experiences,
		Website:     website,
		Github:      github,
		Linkedin:    linkedin,
		Twitter:     twitter,
		Avatar:      avatar,
		Location:    location,
		Email:       email,
	}
}

func (msg *MsgCreateProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Username == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "username cannot be empty")
	}
	if msg.Email == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "email cannot be empty")
	}
	return nil
}
