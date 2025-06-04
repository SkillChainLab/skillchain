package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransferToken{}

func NewMsgTransferToken(creator string, to string, amount string) *MsgTransferToken {
	return &MsgTransferToken{
		Creator: creator,
		To:      to,
		Amount:  amount,
	}
}

func (msg *MsgTransferToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
