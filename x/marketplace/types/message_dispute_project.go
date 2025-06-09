package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDisputeProject{}

func NewMsgDisputeProject(creator string, projectId string, reason string, evidence string) *MsgDisputeProject {
	return &MsgDisputeProject{
		Creator:   creator,
		ProjectId: projectId,
		Reason:    reason,
		Evidence:  evidence,
	}
}

func (msg *MsgDisputeProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
