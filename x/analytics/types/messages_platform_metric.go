package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePlatformMetric{}

func NewMsgCreatePlatformMetric(
	creator string,
	index string,
	metricName string,
	metricValue uint64,
	metricType string,
	period string,
	timestamp uint64,
	metadata string,

) *MsgCreatePlatformMetric {
	return &MsgCreatePlatformMetric{
		Creator:     creator,
		Index:       index,
		MetricName:  metricName,
		MetricValue: metricValue,
		MetricType:  metricType,
		Period:      period,
		Timestamp:   timestamp,
		Metadata:    metadata,
	}
}

func (msg *MsgCreatePlatformMetric) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePlatformMetric{}

func NewMsgUpdatePlatformMetric(
	creator string,
	index string,
	metricName string,
	metricValue uint64,
	metricType string,
	period string,
	timestamp uint64,
	metadata string,

) *MsgUpdatePlatformMetric {
	return &MsgUpdatePlatformMetric{
		Creator:     creator,
		Index:       index,
		MetricName:  metricName,
		MetricValue: metricValue,
		MetricType:  metricType,
		Period:      period,
		Timestamp:   timestamp,
		Metadata:    metadata,
	}
}

func (msg *MsgUpdatePlatformMetric) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePlatformMetric{}

func NewMsgDeletePlatformMetric(
	creator string,
	index string,

) *MsgDeletePlatformMetric {
	return &MsgDeletePlatformMetric{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeletePlatformMetric) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
