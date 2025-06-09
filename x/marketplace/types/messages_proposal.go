package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateProposal{}

func NewMsgCreateProposal(
	creator string,
	index string,
	jobPostingId string,
	freelancerAddress string,
	proposedAmount string,
	proposedCurrency string,
	proposedTimeline string,
	coverLetter string,
	freelancerReputation string,
	isAccepted string,
	createdAt string,

) *MsgCreateProposal {
	return &MsgCreateProposal{
		Creator:              creator,
		Index:                index,
		JobPostingId:         jobPostingId,
		FreelancerAddress:    freelancerAddress,
		ProposedAmount:       proposedAmount,
		ProposedCurrency:     proposedCurrency,
		ProposedTimeline:     proposedTimeline,
		CoverLetter:          coverLetter,
		FreelancerReputation: freelancerReputation,
		IsAccepted:           isAccepted,
		CreatedAt:            createdAt,
	}
}

func (msg *MsgCreateProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateProposal{}

func NewMsgUpdateProposal(
	creator string,
	index string,
	jobPostingId string,
	freelancerAddress string,
	proposedAmount string,
	proposedCurrency string,
	proposedTimeline string,
	coverLetter string,
	freelancerReputation string,
	isAccepted string,
	createdAt string,

) *MsgUpdateProposal {
	return &MsgUpdateProposal{
		Creator:              creator,
		Index:                index,
		JobPostingId:         jobPostingId,
		FreelancerAddress:    freelancerAddress,
		ProposedAmount:       proposedAmount,
		ProposedCurrency:     proposedCurrency,
		ProposedTimeline:     proposedTimeline,
		CoverLetter:          coverLetter,
		FreelancerReputation: freelancerReputation,
		IsAccepted:           isAccepted,
		CreatedAt:            createdAt,
	}
}

func (msg *MsgUpdateProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteProposal{}

func NewMsgDeleteProposal(
	creator string,
	index string,

) *MsgDeleteProposal {
	return &MsgDeleteProposal{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
