package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgApproveVerificationRequest(institutionAddress, requestId string) *MsgApproveVerificationRequest {
	return &MsgApproveVerificationRequest{
		InstitutionAddress: institutionAddress,
		RequestId:          requestId,
	}
}

func NewMsgRejectVerificationRequest(institutionAddress, requestId, reason string) *MsgRejectVerificationRequest {
	return &MsgRejectVerificationRequest{
		InstitutionAddress: institutionAddress,
		RequestId:          requestId,
		Reason:             reason,
	}
}

func NewMsgCreateVerifiedInstitution(creator, address, name, website string, verificationCategories []string, verificationLevel uint32) *MsgCreateVerifiedInstitution {
	return &MsgCreateVerifiedInstitution{
		Creator:                creator,
		Address:                address,
		Name:                   name,
		Website:                website,
		VerificationCategories: verificationCategories,
		VerificationLevel:      verificationLevel,
	}
}

func NewMsgCreateVerificationRequest(userAddress, institutionAddress string, skills []string, evidence string) *MsgCreateVerificationRequest {
	return &MsgCreateVerificationRequest{
		UserAddress:        userAddress,
		InstitutionAddress: institutionAddress,
		Skills:             skills,
		Evidence:           evidence,
	}
}

func (msg *MsgApproveVerificationRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.InstitutionAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid institution address: %s", err)
	}
	if msg.RequestId == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("request id cannot be empty")
	}
	return nil
}

func (msg *MsgRejectVerificationRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.InstitutionAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid institution address: %s", err)
	}
	if msg.RequestId == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("request id cannot be empty")
	}
	if msg.Reason == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("reason cannot be empty")
	}
	return nil
}

func (msg *MsgCreateVerifiedInstitution) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid creator address: %s", err)
	}
	if msg.Address == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("address cannot be empty")
	}
	if msg.Name == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("name cannot be empty")
	}
	if msg.Website == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("website cannot be empty")
	}
	if len(msg.VerificationCategories) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("verification categories cannot be empty")
	}
	return nil
}

func (msg *MsgCreateVerificationRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.UserAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid user address: %s", err)
	}
	if msg.InstitutionAddress == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("institution address cannot be empty")
	}
	if len(msg.Skills) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("skills cannot be empty")
	}
	if msg.Evidence == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("evidence cannot be empty")
	}
	return nil
}
