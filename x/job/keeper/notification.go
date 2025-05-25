package keeper

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

const (
	NotificationKeyPrefix = "Notification/value/"
)

// CreateNotification creates a new notification or updates an existing one
func (k Keeper) CreateNotification(ctx context.Context, notification types.Notification) error {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(NotificationKeyPrefix))

	// Check if the notification already exists
	existingNotification, err := k.GetNotificationByID(ctx, notification.Id)
	if err == nil {
		// Update the existing notification
		existingNotification.IsRead = notification.IsRead
		notification = existingNotification
	}

	// Set notification
	b := k.cdc.MustMarshal(&notification)
	store.Set(types.KeyPrefix(notification.Id), b)

	return nil
}

// GetNotificationByID returns a notification by ID
func (k Keeper) GetNotificationByID(ctx context.Context, id string) (types.Notification, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(NotificationKeyPrefix))

	var notification types.Notification
	b := store.Get(types.KeyPrefix(id))
	if b == nil {
		return notification, status.Error(codes.NotFound, "notification not found")
	}

	k.cdc.MustUnmarshal(b, &notification)
	return notification, nil
}

// ListNotifications returns all notifications for a recipient
func (k Keeper) ListNotifications(ctx context.Context, recipient string, req *query.PageRequest) ([]types.Notification, *query.PageResponse, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(NotificationKeyPrefix))

	var notifications []types.Notification
	pageRes, err := query.Paginate(store, req, func(key []byte, value []byte) error {
		var notification types.Notification
		if err := k.cdc.Unmarshal(value, &notification); err != nil {
			return err
		}

		if notification.Recipient == recipient {
			notifications = append(notifications, notification)
		}

		return nil
	})

	if err != nil {
		return nil, nil, status.Error(codes.Internal, err.Error())
	}

	return notifications, pageRes, nil
}

// MarkNotificationAsRead marks a notification as read
func (k Keeper) MarkNotificationAsRead(ctx context.Context, id string) error {
	notification, err := k.GetNotificationByID(ctx, id)
	if err != nil {
		return err
	}

	notification.IsRead = true
	return k.CreateNotification(ctx, notification)
}

// CreateApplicationNotification creates a notification for a new job application
func (k Keeper) CreateApplicationNotification(ctx context.Context, jobId string, applicationId string, applicant string, jobCreator string) error {
	now := time.Now()
	notification := types.Notification{
		Id:            fmt.Sprintf("app_%s_%s", jobId, applicationId),
		Recipient:     jobCreator,
		Sender:        applicant,
		Type:          "APPLICATION_RECEIVED",
		Content:       fmt.Sprintf("New application received for job #%s", jobId),
		IsRead:        false,
		CreatedAt:     &now,
		JobId:         jobId,
		ApplicationId: applicationId,
	}

	return k.CreateNotification(ctx, notification)
}

// CreateReviewNotification creates a notification for an application review
func (k Keeper) CreateReviewNotification(ctx context.Context, jobId string, applicationId string, applicant string, reviewer string, status string) error {
	now := time.Now()
	notification := types.Notification{
		Id:            fmt.Sprintf("review_%s_%s", jobId, applicationId),
		Recipient:     applicant,
		Sender:        reviewer,
		Type:          "APPLICATION_REVIEWED",
		Content:       fmt.Sprintf("Your application for job #%s has been %s", jobId, status),
		IsRead:        false,
		CreatedAt:     &now,
		JobId:         jobId,
		ApplicationId: applicationId,
	}

	return k.CreateNotification(ctx, notification)
}

// CreateJobUpdateNotification creates a notification for a job update
func (k Keeper) CreateJobUpdateNotification(ctx context.Context, jobId string, jobCreator string, applicants []string) error {
	now := time.Now()
	for _, applicant := range applicants {
		notification := types.Notification{
			Id:        fmt.Sprintf("update_%s_%s", jobId, applicant),
			Recipient: applicant,
			Sender:    jobCreator,
			Type:      "JOB_UPDATED",
			Content:   fmt.Sprintf("Job #%s has been updated", jobId),
			IsRead:    false,
			CreatedAt: &now,
			JobId:     jobId,
		}

		if err := k.CreateNotification(ctx, notification); err != nil {
			return err
		}
	}

	return nil
}
