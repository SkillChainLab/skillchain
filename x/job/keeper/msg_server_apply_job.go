package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

func (k msgServer) ApplyJob(goCtx context.Context, msg *types.MsgApplyJob) (*types.MsgApplyJobResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger().Info("ApplyJob function called", "creator", msg.Creator, "job_id", msg.JobId)

	if msg.Creator == "" || msg.CoverLetter == "" || msg.JobId == "" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid request: creator, cover letter and job ID cannot be empty")
	}

	jobID, err := strconv.ParseUint(msg.JobId, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid job ID: %s", msg.JobId)
	}

	job, found := k.GetJob(ctx, jobID)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrJobNotFound, "job %d not found", jobID)
	}

	// Validate that the job exists and has required fields
	if job.Id == 0 || job.Title == "" || job.Creator == "" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid job data")
	}

	if err := k.AppendApplication(goCtx, jobID, msg.Creator, msg.CoverLetter); err != nil {
		return nil, err
	}

	return &types.MsgApplyJobResponse{}, nil
}
