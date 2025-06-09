package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateRevenueRecord{}

func NewMsgCreateRevenueRecord(
	creator string,
	index string,
	transactionType string,
	amount uint64,
	currency string,
	fromAddress string,
	toAddress string,
	timestamp uint64,
	feeAmount uint64,
	projectId string,
	platformFee uint64,

) *MsgCreateRevenueRecord {
	return &MsgCreateRevenueRecord{
		Creator:         creator,
		Index:           index,
		TransactionType: transactionType,
		Amount:          amount,
		Currency:        currency,
		FromAddress:     fromAddress,
		ToAddress:       toAddress,
		Timestamp:       timestamp,
		FeeAmount:       feeAmount,
		ProjectId:       projectId,
		PlatformFee:     platformFee,
	}
}

func (msg *MsgCreateRevenueRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRevenueRecord{}

func NewMsgUpdateRevenueRecord(
	creator string,
	index string,
	transactionType string,
	amount uint64,
	currency string,
	fromAddress string,
	toAddress string,
	timestamp uint64,
	feeAmount uint64,
	projectId string,
	platformFee uint64,

) *MsgUpdateRevenueRecord {
	return &MsgUpdateRevenueRecord{
		Creator:         creator,
		Index:           index,
		TransactionType: transactionType,
		Amount:          amount,
		Currency:        currency,
		FromAddress:     fromAddress,
		ToAddress:       toAddress,
		Timestamp:       timestamp,
		FeeAmount:       feeAmount,
		ProjectId:       projectId,
		PlatformFee:     platformFee,
	}
}

func (msg *MsgUpdateRevenueRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteRevenueRecord{}

func NewMsgDeleteRevenueRecord(
	creator string,
	index string,

) *MsgDeleteRevenueRecord {
	return &MsgDeleteRevenueRecord{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteRevenueRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
