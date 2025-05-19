package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

func (k msgServer) UpdateJob(goCtx context.Context, msg *types.MsgUpdateJob) (*types.MsgUpdateJobResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check authorization
	if err := k.Keeper.CheckJobUpdateAuthorization(ctx, msg.Id, msg.Creator); err != nil {
		return nil, err
	}

	// Get the job
	job, found := k.Keeper.GetJob(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrNotFound, "job %d not found", msg.Id)
	}

	// Update job fields
	if msg.Title != "" {
		job.Title = msg.Title
	}
	if msg.Description != "" {
		job.Description = msg.Description
	}
	if msg.Budget != "" {
		job.Budget = msg.Budget
	}

	// Save updated job
	k.Keeper.SetJob(ctx, job)

	// Get all applications for this job
	applications := k.Keeper.GetJobApplications(ctx, job.Id)

	// Create notifications for all applicants
	var applicants []string
	for _, app := range applications {
		applicants = append(applicants, app.Applicant)
	}

	err := k.Keeper.CreateJobUpdateNotification(ctx, fmt.Sprint(job.Id), job.Creator, applicants)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "failed to create notifications: %s", err.Error())
	}

	return &types.MsgUpdateJobResponse{}, nil
} 