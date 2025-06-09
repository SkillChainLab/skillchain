package types

import (
	"testing"

	"skillchain/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateNotificationSettings_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateNotificationSettings
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateNotificationSettings{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateNotificationSettings{
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

func TestMsgUpdateNotificationSettings_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateNotificationSettings
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateNotificationSettings{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateNotificationSettings{
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

func TestMsgDeleteNotificationSettings_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteNotificationSettings
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteNotificationSettings{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteNotificationSettings{
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
