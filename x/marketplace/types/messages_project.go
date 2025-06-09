package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateProject{}

func NewMsgCreateProject(
	creator string,
	index string,
	jobPostingId string,
	proposalId string,
	clientAddress string,
	freelancerAddress string,
	totalAmount string,
	paidAmount string,
	escrowAmount string,
	status string,
	startDate string,
	expectedEndDate string,
	actualEndDate string,

) *MsgCreateProject {
	return &MsgCreateProject{
		Creator:           creator,
		Index:             index,
		JobPostingId:      jobPostingId,
		ProposalId:        proposalId,
		ClientAddress:     clientAddress,
		FreelancerAddress: freelancerAddress,
		TotalAmount:       totalAmount,
		PaidAmount:        paidAmount,
		EscrowAmount:      escrowAmount,
		Status:            status,
		StartDate:         startDate,
		ExpectedEndDate:   expectedEndDate,
		ActualEndDate:     actualEndDate,
	}
}

func (msg *MsgCreateProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateProject{}

func NewMsgUpdateProject(
	creator string,
	index string,
	jobPostingId string,
	proposalId string,
	clientAddress string,
	freelancerAddress string,
	totalAmount string,
	paidAmount string,
	escrowAmount string,
	status string,
	startDate string,
	expectedEndDate string,
	actualEndDate string,

) *MsgUpdateProject {
	return &MsgUpdateProject{
		Creator:           creator,
		Index:             index,
		JobPostingId:      jobPostingId,
		ProposalId:        proposalId,
		ClientAddress:     clientAddress,
		FreelancerAddress: freelancerAddress,
		TotalAmount:       totalAmount,
		PaidAmount:        paidAmount,
		EscrowAmount:      escrowAmount,
		Status:            status,
		StartDate:         startDate,
		ExpectedEndDate:   expectedEndDate,
		ActualEndDate:     actualEndDate,
	}
}

func (msg *MsgUpdateProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteProject{}

func NewMsgDeleteProject(
	creator string,
	index string,

) *MsgDeleteProject {
	return &MsgDeleteProject{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
