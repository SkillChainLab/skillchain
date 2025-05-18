package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgReviewApplication{}

func NewMsgReviewApplication(creator string, jobId uint64, applicant string, status string) *MsgReviewApplication {
	return &MsgReviewApplication{
		Creator:   creator,
		JobId:     jobId,
		Applicant: applicant,
		Status:    status,
	}
}

func (msg *MsgReviewApplication) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
