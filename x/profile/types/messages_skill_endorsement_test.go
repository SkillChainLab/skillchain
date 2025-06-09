package types

import (
	"testing"

	"skillchain/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateSkillEndorsement_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateSkillEndorsement
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateSkillEndorsement{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateSkillEndorsement{
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

func TestMsgUpdateSkillEndorsement_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateSkillEndorsement
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateSkillEndorsement{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateSkillEndorsement{
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

func TestMsgDeleteSkillEndorsement_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteSkillEndorsement
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteSkillEndorsement{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteSkillEndorsement{
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
