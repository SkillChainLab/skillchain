package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgApplyJob{}

func NewMsgApplyJob(creator string, jobId string, coverLetter string) *MsgApplyJob {
	return &MsgApplyJob{
		Creator:     creator,
		JobId:       jobId,
		CoverLetter: coverLetter,
	}
}

func (msg *MsgApplyJob) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
