package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApplyJob = "apply_job"

var _ sdk.Msg = &MsgApplyJob{}

func NewMsgApplyJob(creator string, jobId uint64, coverLetter string) *MsgApplyJob {
	return &MsgApplyJob{
		Creator:     creator,
		JobId:       jobId,
		CoverLetter: coverLetter,
	}
}

func (msg *MsgApplyJob) Route() string {
	return RouterKey
}

func (msg *MsgApplyJob) Type() string {
	return TypeMsgApplyJob
}

func (msg *MsgApplyJob) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApplyJob) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApplyJob) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.JobId == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "job ID cannot be zero")
	}
	if msg.CoverLetter == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "cover letter cannot be empty")
	}
	return nil
}
