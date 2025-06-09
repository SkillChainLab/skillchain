package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateUserSkill{}

func NewMsgCreateUserSkill(
	creator string,
	index string,
	owner string,
	skillName string,
	proficiencyLevel string,
	yearsExperience uint64,
	verified bool,
	verifiedBy string,
	verificationDate uint64,
	endorsementCount uint64,

) *MsgCreateUserSkill {
	return &MsgCreateUserSkill{
		Creator:          creator,
		Index:            index,
		Owner:            owner,
		SkillName:        skillName,
		ProficiencyLevel: proficiencyLevel,
		YearsExperience:  yearsExperience,
		Verified:         verified,
		VerifiedBy:       verifiedBy,
		VerificationDate: verificationDate,
		EndorsementCount: endorsementCount,
	}
}

func (msg *MsgCreateUserSkill) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateUserSkill{}

func NewMsgUpdateUserSkill(
	creator string,
	index string,
	owner string,
	skillName string,
	proficiencyLevel string,
	yearsExperience uint64,
	verified bool,
	verifiedBy string,
	verificationDate uint64,
	endorsementCount uint64,

) *MsgUpdateUserSkill {
	return &MsgUpdateUserSkill{
		Creator:          creator,
		Index:            index,
		Owner:            owner,
		SkillName:        skillName,
		ProficiencyLevel: proficiencyLevel,
		YearsExperience:  yearsExperience,
		Verified:         verified,
		VerifiedBy:       verifiedBy,
		VerificationDate: verificationDate,
		EndorsementCount: endorsementCount,
	}
}

func (msg *MsgUpdateUserSkill) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteUserSkill{}

func NewMsgDeleteUserSkill(
	creator string,
	index string,

) *MsgDeleteUserSkill {
	return &MsgDeleteUserSkill{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteUserSkill) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
