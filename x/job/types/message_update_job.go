package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"fmt"
)

const TypeMsgUpdateJob = "update_job"

var _ sdk.Msg = &MsgUpdateJob{}

func NewMsgUpdateJob(creator string, id uint64, title string, description string, budget string) *MsgUpdateJob {
	return &MsgUpdateJob{
		Creator:     creator,
		Id:          id,
		Title:       title,
		Description: description,
		Budget:      budget,
	}
}

func (msg *MsgUpdateJob) Route() string {
	return RouterKey
}

func (msg *MsgUpdateJob) Type() string {
	return TypeMsgUpdateJob
}

func (msg *MsgUpdateJob) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateJob) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateJob) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return fmt.Errorf("invalid creator address: %w", err)
	}
	return nil
} 