package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateJobPosting{}

func NewMsgCreateJobPosting(
	creator string,
	index string,
	clientAddress string,
	title string,
	description string,
	skillsRequired string,
	budgetAmount string,
	paymentCurrency string,
	deadline string,
	isActive string,
	createdAt string,

) *MsgCreateJobPosting {
	return &MsgCreateJobPosting{
		Creator:         creator,
		Index:           index,
		ClientAddress:   clientAddress,
		Title:           title,
		Description:     description,
		SkillsRequired:  skillsRequired,
		BudgetAmount:    budgetAmount,
		PaymentCurrency: paymentCurrency,
		Deadline:        deadline,
		IsActive:        isActive,
		CreatedAt:       createdAt,
	}
}

func (msg *MsgCreateJobPosting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateJobPosting{}

func NewMsgUpdateJobPosting(
	creator string,
	index string,
	clientAddress string,
	title string,
	description string,
	skillsRequired string,
	budgetAmount string,
	paymentCurrency string,
	deadline string,
	isActive string,
	createdAt string,

) *MsgUpdateJobPosting {
	return &MsgUpdateJobPosting{
		Creator:         creator,
		Index:           index,
		ClientAddress:   clientAddress,
		Title:           title,
		Description:     description,
		SkillsRequired:  skillsRequired,
		BudgetAmount:    budgetAmount,
		PaymentCurrency: paymentCurrency,
		Deadline:        deadline,
		IsActive:        isActive,
		CreatedAt:       createdAt,
	}
}

func (msg *MsgUpdateJobPosting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteJobPosting{}

func NewMsgDeleteJobPosting(
	creator string,
	index string,

) *MsgDeleteJobPosting {
	return &MsgDeleteJobPosting{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteJobPosting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
