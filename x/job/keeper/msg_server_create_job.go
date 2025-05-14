package keeper

import (
	"context"

	"github.com/SkillChainLab/skillchain/x/job/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateJob(goCtx context.Context, msg *types.MsgCreateJob) (*types.MsgCreateJobResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// İş ilanı nesnesi
	job := types.Job{
		Creator:     msg.Creator,
		Id:          k.AppendJob(ctx, msg.Title, msg.Description, msg.Budget, msg.Creator),
		Title:       msg.Title,
		Description: msg.Description,
		Budget:      msg.Budget,
	}

	return &types.MsgCreateJobResponse{Id: job.Id}, nil
}

