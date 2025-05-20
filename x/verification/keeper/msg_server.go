package keeper

import (
	"context"

	"github.com/SkillChainLab/skillchain/x/verification/types"
	"github.com/google/uuid"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// CreateVerifiedInstitution handles the MsgCreateVerifiedInstitution message
func (k msgServer) CreateVerifiedInstitution(ctx context.Context, msg *types.MsgCreateVerifiedInstitution) (*types.MsgCreateVerifiedInstitutionResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	institution := types.VerifiedInstitution{
		Address:               msg.Address,
		Name:                  msg.Name,
		Website:               msg.Website,
		AddedBy:               msg.Creator,
		VerificationCategories: msg.VerificationCategories,
		VerificationLevel:     msg.VerificationLevel,
		Status:                "active",
		LastVerificationDate:  uint64(sdkCtx.BlockTime().Unix()),
	}
	k.SetVerifiedInstitution(sdkCtx, institution)
	return &types.MsgCreateVerifiedInstitutionResponse{Address: institution.Address}, nil
}

// CreateVerificationRequest handles the MsgCreateVerificationRequest message
func (k msgServer) CreateVerificationRequest(ctx context.Context, msg *types.MsgCreateVerificationRequest) (*types.MsgCreateVerificationRequestResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	requestID := uuid.New().String()
	request := types.VerificationRequest{
		RequestId:          requestID,
		UserAddress:        msg.UserAddress,
		InstitutionAddress: msg.InstitutionAddress,
		Skills:             msg.Skills,
		Status:             "pending",
		Evidence:           msg.Evidence,
		CreatedAt:          uint64(sdkCtx.BlockTime().Unix()),
		UpdatedAt:          uint64(sdkCtx.BlockTime().Unix()),
	}
	k.SetVerificationRequest(sdkCtx, request)
	return &types.MsgCreateVerificationRequestResponse{RequestId: requestID}, nil
}

func (k msgServer) ApproveVerificationRequest(ctx context.Context, msg *types.MsgApproveVerificationRequest) (*types.MsgApproveVerificationRequestResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Get the verification request
	request, found := k.GetVerificationRequest(sdkCtx, msg.RequestId)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound.Wrapf("verification request with id %s not found", msg.RequestId)
	}

	// Check if the institution is authorized
	if request.InstitutionAddress != msg.InstitutionAddress {
		return nil, sdkerrors.ErrUnauthorized.Wrap("institution not authorized to approve this request")
	}

	// Check if the request is in pending status
	if request.Status != "pending" {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("request is not in pending status, current status: %s", request.Status)
	}

	// Update the request status
	request.Status = "approved"
	request.UpdatedAt = uint64(sdkCtx.BlockTime().Unix())
	k.SetVerificationRequest(sdkCtx, request)

	return &types.MsgApproveVerificationRequestResponse{
		RequestId: request.RequestId,
	}, nil
}

func (k msgServer) RejectVerificationRequest(ctx context.Context, msg *types.MsgRejectVerificationRequest) (*types.MsgRejectVerificationRequestResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Get the verification request
	request, found := k.GetVerificationRequest(sdkCtx, msg.RequestId)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound.Wrapf("verification request with id %s not found", msg.RequestId)
	}

	// Check if the institution is authorized
	if request.InstitutionAddress != msg.InstitutionAddress {
		return nil, sdkerrors.ErrUnauthorized.Wrap("institution not authorized to reject this request")
	}

	// Check if the request is in pending status
	if request.Status != "pending" {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("request is not in pending status, current status: %s", request.Status)
	}

	// Update the request status
	request.Status = "rejected"
	request.UpdatedAt = uint64(sdkCtx.BlockTime().Unix())
	k.SetVerificationRequest(sdkCtx, request)

	return &types.MsgRejectVerificationRequestResponse{
		RequestId: request.RequestId,
	}, nil
}
