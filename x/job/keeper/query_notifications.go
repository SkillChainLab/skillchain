package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

// GetNotifications returns all notifications for a recipient
func (k Keeper) GetNotifications(ctx context.Context, req *types.QueryGetNotificationsRequest) (*types.QueryGetNotificationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	notifications, pageRes, err := k.ListNotifications(ctx, req.Recipient, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Convert []Notification to []*Notification
	notificationPtrs := make([]*types.Notification, len(notifications))
	for i := range notifications {
		notificationPtrs[i] = &notifications[i]
	}

	return &types.QueryGetNotificationsResponse{
		Notifications: notificationPtrs,
		Pagination:    pageRes,
	}, nil
}

// GetNotification returns a specific notification by ID
func (k Keeper) GetNotification(ctx context.Context, req *types.QueryGetNotificationRequest) (*types.QueryGetNotificationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	notification, err := k.GetNotificationByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetNotificationResponse{
		Notification: &notification,
	}, nil
} 