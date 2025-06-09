package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgReleasePayment{}

func NewMsgReleasePayment(creator string, milestoneId string, rating string, feedback string) *MsgReleasePayment {
	return &MsgReleasePayment{
		Creator:     creator,
		MilestoneId: milestoneId,
		Rating:      rating,
		Feedback:    feedback,
	}
}

func (msg *MsgReleasePayment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
