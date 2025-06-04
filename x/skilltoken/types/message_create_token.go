package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateToken{}

func NewMsgCreateToken(creator string, symbol string, name string, decimals string, mintable string, burnable string, transferable string) *MsgCreateToken {
	return &MsgCreateToken{
		Creator:      creator,
		Symbol:       symbol,
		Name:         name,
		Decimals:     decimals,
		Mintable:     mintable,
		Burnable:     burnable,
		Transferable: transferable,
	}
}

func (msg *MsgCreateToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
