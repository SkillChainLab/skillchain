package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSkillEndorsement{}

func NewMsgCreateSkillEndorsement(
	creator string,
	index string,
	endorser string,
	targetUser string,
	skillName string,
	endorsementType string,
	comment string,
	createdAt uint64,
	skillTokensStaked uint64,

) *MsgCreateSkillEndorsement {
	return &MsgCreateSkillEndorsement{
		Creator:           creator,
		Index:             index,
		Endorser:          endorser,
		TargetUser:        targetUser,
		SkillName:         skillName,
		EndorsementType:   endorsementType,
		Comment:           comment,
		CreatedAt:         createdAt,
		SkillTokensStaked: skillTokensStaked,
	}
}

func (msg *MsgCreateSkillEndorsement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSkillEndorsement{}

func NewMsgUpdateSkillEndorsement(
	creator string,
	index string,
	endorser string,
	targetUser string,
	skillName string,
	endorsementType string,
	comment string,
	createdAt uint64,
	skillTokensStaked uint64,

) *MsgUpdateSkillEndorsement {
	return &MsgUpdateSkillEndorsement{
		Creator:           creator,
		Index:             index,
		Endorser:          endorser,
		TargetUser:        targetUser,
		SkillName:         skillName,
		EndorsementType:   endorsementType,
		Comment:           comment,
		CreatedAt:         createdAt,
		SkillTokensStaked: skillTokensStaked,
	}
}

func (msg *MsgUpdateSkillEndorsement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSkillEndorsement{}

func NewMsgDeleteSkillEndorsement(
	creator string,
	index string,

) *MsgDeleteSkillEndorsement {
	return &MsgDeleteSkillEndorsement{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteSkillEndorsement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
