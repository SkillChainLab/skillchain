package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

func (k msgServer) ReviewApplication(goCtx context.Context, msg *types.MsgReviewApplication) (*types.MsgReviewApplicationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check authorization
	if err := k.Keeper.CheckJobReviewAuthorization(ctx, msg.JobId, msg.Creator); err != nil {
		return nil, err
	}

	// Get the application
	application, found := k.Keeper.GetJobApplication(ctx, msg.JobId, msg.Applicant)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrNotFound, "application not found for job %d and applicant %s", msg.JobId, msg.Applicant)
	}

	// Update application status
	application.Status = msg.Status
	k.Keeper.SetJobApplication(ctx, application)

	// Create notification for the applicant
	err := k.Keeper.CreateReviewNotification(ctx, fmt.Sprint(msg.JobId), fmt.Sprintf("%d:%s", msg.JobId, msg.Applicant), msg.Applicant, msg.Creator, msg.Status)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "failed to create notification: %s", err.Error())
	}

	return &types.MsgReviewApplicationResponse{}, nil
}
