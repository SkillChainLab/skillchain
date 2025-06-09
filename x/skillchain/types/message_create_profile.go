package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateProfile{}

func NewMsgCreateProfile(creator string, displayName string, bio string, location string, website string, github string, linkedin string, twitter string) *MsgCreateProfile {
	return &MsgCreateProfile{
		Creator:     creator,
		DisplayName: displayName,
		Bio:         bio,
		Location:    location,
		Website:     website,
		Github:      github,
		Linkedin:    linkedin,
		Twitter:     twitter,
	}
}

func (msg *MsgCreateProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
