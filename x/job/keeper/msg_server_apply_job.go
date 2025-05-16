package keeper

import (
	"context"
	"strconv"
	"github.com/SkillChainLab/skillchain/x/job/types"
)

func (k msgServer) ApplyJob(goCtx context.Context, msg *types.MsgApplyJob) (*types.MsgApplyJobResponse, error) {


	jobID, err := strconv.ParseUint(msg.JobId, 10, 64)
	if err != nil {
		return nil, err
	}

	err = k.AppendApplication(goCtx, jobID, msg.Creator, msg.CoverLetter)
	if err != nil {
		return nil, err
	}

	return &types.MsgApplyJobResponse{
		Id: jobID,
	}, nil
}
