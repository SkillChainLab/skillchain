package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

func (k msgServer) ApplyJob(goCtx context.Context, msg *types.MsgApplyJob) (*types.MsgApplyJobResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check authorization
	if err := k.Keeper.CheckApplicationAuthorization(ctx, msg.JobId, msg.Creator); err != nil {
		return nil, err
	}

	// Check if job exists
	job, found := k.Keeper.GetJob(ctx, msg.JobId)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrNotFound, "job %d not found", msg.JobId)
	}

	// Create application
	application := types.Application{
		JobId:          msg.JobId,
		Applicant:      msg.Creator,
		CoverLetter:    msg.CoverLetter,
		Status:         "PENDING",
		JobTitle:       job.Title,
		JobDescription: job.Description,
		JobBudget:      job.Budget,
	}

	// Save application
	k.Keeper.SetJobApplication(ctx, application)

	// Create notification for job creator
	err := k.Keeper.CreateApplicationNotification(ctx, fmt.Sprint(msg.JobId), fmt.Sprintf("%d:%s", msg.JobId, msg.Creator), msg.Creator, job.Creator)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "failed to create notification: %s", err.Error())
	}

	return &types.MsgApplyJobResponse{
		JobId:     msg.JobId,
		Applicant: msg.Creator,
		JobTitle:  job.Title,
	}, nil
}
