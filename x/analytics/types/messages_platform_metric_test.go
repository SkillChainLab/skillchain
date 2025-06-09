package types

import (
	"testing"

	"skillchain/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreatePlatformMetric_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreatePlatformMetric
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreatePlatformMetric{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreatePlatformMetric{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdatePlatformMetric_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdatePlatformMetric
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdatePlatformMetric{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdatePlatformMetric{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeletePlatformMetric_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeletePlatformMetric
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeletePlatformMetric{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeletePlatformMetric{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
