package keeper

import (
	"context"

	"github.com/SkillChainLab/skillchain/x/job/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateJob(goCtx context.Context, msg *types.MsgCreateJob) (*types.MsgCreateJobResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create job and get its ID
	jobID := k.AppendJob(ctx, msg.Title, msg.Description, msg.Budget, msg.Creator)

	return &types.MsgCreateJobResponse{Id: jobID}, nil
}
