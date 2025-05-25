package keeper

import (
	"context"
	// "cosmossdk.io/store/prefix" // unused
	// "github.com/cosmos/cosmos-sdk/types/query" // unused
	"github.com/SkillChainLab/skillchain/x/verification/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// VerifiedInstitution returns a single VerifiedInstitution by address
func (k Keeper) VerifiedInstitution(c context.Context, req *types.QueryVerifiedInstitutionRequest) (*types.QueryVerifiedInstitutionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	institution, found := k.GetVerifiedInstitution(ctx, req.Address)
	if !found {
		return nil, status.Error(codes.NotFound, "institution not found")
	}
	return &types.QueryVerifiedInstitutionResponse{Institution: &institution}, nil
}

// VerifiedInstitutionAll returns all VerifiedInstitutions
func (k Keeper) VerifiedInstitutionAll(c context.Context, req *types.QueryAllVerifiedInstitutionRequest) (*types.QueryAllVerifiedInstitutionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	institutions := k.ListVerifiedInstitutions(ctx)
	var result []*types.VerifiedInstitution
	for _, i := range institutions {
		iCopy := i
		result = append(result, &iCopy)
	}
	return &types.QueryAllVerifiedInstitutionResponse{Institutions: result}, nil
}

// VerificationRequest returns a single VerificationRequest by request_id
func (k Keeper) VerificationRequest(c context.Context, req *types.QueryVerificationRequestRequest) (*types.QueryVerificationRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	request, found := k.GetVerificationRequest(ctx, req.RequestId)
	if !found {
		return nil, status.Error(codes.NotFound, "request not found")
	}
	return &types.QueryVerificationRequestResponse{Request: &request}, nil
}

// VerificationRequestAll returns all VerificationRequests
func (k Keeper) VerificationRequestAll(c context.Context, req *types.QueryAllVerificationRequestRequest) (*types.QueryAllVerificationRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	requests := k.ListVerificationRequests(ctx)
	var result []*types.VerificationRequest
	for _, r := range requests {
		rCopy := r
		result = append(result, &rCopy)
	}
	return &types.QueryAllVerificationRequestResponse{Requests: result}, nil
}
