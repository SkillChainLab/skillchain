package types

import (
	"testing"

	"skillchain/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateFilePermission_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateFilePermission
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateFilePermission{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateFilePermission{
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

func TestMsgUpdateFilePermission_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateFilePermission
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateFilePermission{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateFilePermission{
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

func TestMsgDeleteFilePermission_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteFilePermission
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteFilePermission{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteFilePermission{
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
