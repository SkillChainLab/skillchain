package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

func (k msgServer) MarkNotificationAsRead(goCtx context.Context, msg *types.MsgMarkNotificationAsRead) (*types.MsgMarkNotificationAsReadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve the notification by ID
	notification, err := k.GetNotificationByID(ctx, msg.NotificationId)
	if err != nil {
		return nil, errors.Wrapf(err, "notification not found: %s", err.Error())
	}

	// Mark the notification as read
	notification.IsRead = true

	// Save the updated notification
	if err := k.CreateNotification(ctx, notification); err != nil {
		return nil, errors.Wrapf(err, "failed to update notification: %s", err.Error())
	}

	return &types.MsgMarkNotificationAsReadResponse{}, nil
}
