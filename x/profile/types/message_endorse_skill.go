package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEndorseSkill{}

func NewMsgEndorseSkill(creator string, targetUser string, skillName string, endorsementType string, comment string) *MsgEndorseSkill {
	return &MsgEndorseSkill{
		Creator:         creator,
		TargetUser:      targetUser,
		SkillName:       skillName,
		EndorsementType: endorsementType,
		Comment:         comment,
	}
}

func (msg *MsgEndorseSkill) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
