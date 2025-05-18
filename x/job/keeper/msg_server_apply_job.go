package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

func (k msgServer) ApplyJob(goCtx context.Context, msg *types.MsgApplyJob) (*types.MsgApplyJobResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Verify job exists first
	job, found := k.GetJob(ctx, msg.JobId)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrJobNotFound, "job %d not found", msg.JobId)
	}

	err := k.AppendApplication(ctx, msg.JobId, msg.Creator, msg.CoverLetter)
	if err != nil {
		return nil, err
	}

	return &types.MsgApplyJobResponse{
		JobId:     msg.JobId,
		Applicant: msg.Creator,
		JobTitle:  job.Title,
	}, nil
}
