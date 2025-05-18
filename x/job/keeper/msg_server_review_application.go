package keeper

import (
	"context"
	"fmt"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

func (k msgServer) ReviewApplication(goCtx context.Context, msg *types.MsgReviewApplication) (*types.MsgReviewApplicationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate status
	status := strings.ToUpper(msg.Status)
	if status != "APPROVED" && status != "REJECTED" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid status: %s. Must be either APPROVED or REJECTED", msg.Status)
	}

	// Get the job
	job, found := k.GetJob(ctx, msg.JobId)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrJobNotFound, "job %d not found", msg.JobId)
	}

	// Check if the reviewer is the job creator
	if msg.Creator != job.Creator {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "only job creator can review applications")
	}

	// Get the application
	store := k.getStore(ctx)
	key := []byte(fmt.Sprintf("Application/value/%d:%s", msg.JobId, msg.Applicant))
	
	if !store.Has(key) {
		return nil, errorsmod.Wrapf(types.ErrApplicationNotFound, "application not found for job %d and applicant %s", msg.JobId, msg.Applicant)
	}

	var application types.Application
	k.cdc.MustUnmarshal(store.Get(key), &application)

	// Check if the application is already reviewed
	if application.Status != "" && application.Status != "PENDING" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "application already reviewed with status: %s", application.Status)
	}

	// Update application status
	application.Status = status
	bz := k.cdc.MustMarshal(&application)
	store.Set(key, bz)

	return &types.MsgReviewApplicationResponse{}, nil
}
