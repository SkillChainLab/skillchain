package types

import (
	"testing"

	"skillchain/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateRevenueRecord_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateRevenueRecord
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateRevenueRecord{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateRevenueRecord{
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

func TestMsgUpdateRevenueRecord_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateRevenueRecord
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateRevenueRecord{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateRevenueRecord{
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

func TestMsgDeleteRevenueRecord_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteRevenueRecord
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteRevenueRecord{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteRevenueRecord{
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
