package types

import (
	"testing"

	"skillchain/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateFileRecord_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateFileRecord
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateFileRecord{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateFileRecord{
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

func TestMsgUpdateFileRecord_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateFileRecord
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateFileRecord{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateFileRecord{
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

func TestMsgDeleteFileRecord_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteFileRecord
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteFileRecord{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteFileRecord{
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
