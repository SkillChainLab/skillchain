package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateMilestone{}

func NewMsgCreateMilestone(
	creator string,
	index string,
	projectId string,
	title string,
	description string,
	amount string,
	dueDate string,
	status string,
	isCompleted string,
	isPaid string,
	submittedAt string,
	approvedAt string,

) *MsgCreateMilestone {
	return &MsgCreateMilestone{
		Creator:     creator,
		Index:       index,
		ProjectId:   projectId,
		Title:       title,
		Description: description,
		Amount:      amount,
		DueDate:     dueDate,
		Status:      status,
		IsCompleted: isCompleted,
		IsPaid:      isPaid,
		SubmittedAt: submittedAt,
		ApprovedAt:  approvedAt,
	}
}

func (msg *MsgCreateMilestone) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMilestone{}

func NewMsgUpdateMilestone(
	creator string,
	index string,
	projectId string,
	title string,
	description string,
	amount string,
	dueDate string,
	status string,
	isCompleted string,
	isPaid string,
	submittedAt string,
	approvedAt string,

) *MsgUpdateMilestone {
	return &MsgUpdateMilestone{
		Creator:     creator,
		Index:       index,
		ProjectId:   projectId,
		Title:       title,
		Description: description,
		Amount:      amount,
		DueDate:     dueDate,
		Status:      status,
		IsCompleted: isCompleted,
		IsPaid:      isPaid,
		SubmittedAt: submittedAt,
		ApprovedAt:  approvedAt,
	}
}

func (msg *MsgUpdateMilestone) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMilestone{}

func NewMsgDeleteMilestone(
	creator string,
	index string,

) *MsgDeleteMilestone {
	return &MsgDeleteMilestone{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteMilestone) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
