package keeper

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"skillchain/x/notifications/types"
)

var _ types.QueryServer = Keeper{}

// GetUserNotifications returns all notifications for a specific user
func (k Keeper) GetUserNotifications(ctx sdk.Context, userAddress string) []types.Notification {
	var userNotifications []types.Notification
	allNotifications := k.GetAllNotification(ctx)
	
	for _, notification := range allNotifications {
		if notification.UserAddress == userAddress {
			userNotifications = append(userNotifications, notification)
		}
	}
	
	return userNotifications
}

// GetUnreadNotifications returns unread notifications for a user
func (k Keeper) GetUnreadNotifications(ctx sdk.Context, userAddress string) []types.Notification {
	var unreadNotifications []types.Notification
	userNotifications := k.GetUserNotifications(ctx, userAddress)
	
	for _, notification := range userNotifications {
		if !notification.IsRead {
			unreadNotifications = append(unreadNotifications, notification)
		}
	}
	
	return unreadNotifications
}

// GetNotificationsByPriority returns notifications filtered by priority level
func (k Keeper) GetNotificationsByPriority(ctx sdk.Context, userAddress, priority string) []types.Notification {
	var priorityNotifications []types.Notification
	userNotifications := k.GetUserNotifications(ctx, userAddress)
	
	for _, notification := range userNotifications {
		if strings.EqualFold(notification.Priority, priority) {
			priorityNotifications = append(priorityNotifications, notification)
		}
	}
	
	return priorityNotifications
}

// GetNotificationsByType returns notifications filtered by type
func (k Keeper) GetNotificationsByType(ctx sdk.Context, userAddress, notificationType string) []types.Notification {
	var typeNotifications []types.Notification
	userNotifications := k.GetUserNotifications(ctx, userAddress)
	
	for _, notification := range userNotifications {
		if strings.EqualFold(notification.NotificationType, notificationType) {
			typeNotifications = append(typeNotifications, notification)
		}
	}
	
	return typeNotifications
}

// GetNotificationsByDateRange returns notifications within a date range
func (k Keeper) GetNotificationsByDateRange(ctx sdk.Context, userAddress string, startTime, endTime uint64) []types.Notification {
	var rangeNotifications []types.Notification
	userNotifications := k.GetUserNotifications(ctx, userAddress)
	
	for _, notification := range userNotifications {
		if notification.CreatedAt >= startTime && notification.CreatedAt <= endTime {
			rangeNotifications = append(rangeNotifications, notification)
		}
	}
	
	return rangeNotifications
}

// GetRecentNotifications returns recent notifications (last 24 hours)
func (k Keeper) GetRecentNotifications(ctx sdk.Context, userAddress string) []types.Notification {
	currentTime := uint64(time.Now().Unix())
	oneDayAgo := currentTime - 86400 // 24 hours in seconds
	
	return k.GetNotificationsByDateRange(ctx, userAddress, oneDayAgo, currentTime)
}

// GetNotificationsBySource returns notifications from a specific source module
func (k Keeper) GetNotificationsBySource(ctx sdk.Context, userAddress, sourceModule string) []types.Notification {
	var sourceNotifications []types.Notification
	userNotifications := k.GetUserNotifications(ctx, userAddress)
	
	for _, notification := range userNotifications {
		if strings.EqualFold(notification.SourceModule, sourceModule) {
			sourceNotifications = append(sourceNotifications, notification)
		}
	}
	
	return sourceNotifications
}

// CountUnreadNotifications returns the count of unread notifications for a user
func (k Keeper) CountUnreadNotifications(ctx sdk.Context, userAddress string) int {
	unreadNotifications := k.GetUnreadNotifications(ctx, userAddress)
	return len(unreadNotifications)
}

// CountNotificationsByPriority returns counts grouped by priority
func (k Keeper) CountNotificationsByPriority(ctx sdk.Context, userAddress string) map[string]int {
	counts := make(map[string]int)
	userNotifications := k.GetUserNotifications(ctx, userAddress)
	
	for _, notification := range userNotifications {
		if !notification.IsRead {
			counts[notification.Priority]++
		}
	}
	
	return counts
}

// MarkAllAsRead marks all notifications as read for a user
func (k Keeper) MarkAllAsRead(ctx sdk.Context, userAddress string) int {
	count := 0
	userNotifications := k.GetUserNotifications(ctx, userAddress)
	
	for _, notification := range userNotifications {
		if !notification.IsRead {
			notification.IsRead = true
			k.SetNotification(ctx, notification)
			count++
		}
	}
	
	// Emit batch read event
	if count > 0 {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"notifications_bulk_read",
				sdk.NewAttribute("user_address", userAddress),
				sdk.NewAttribute("count", string(rune(count))),
			),
		)
	}
	
	return count
}

// DeleteOldNotifications removes notifications older than specified days
func (k Keeper) DeleteOldNotifications(ctx sdk.Context, userAddress string, daysOld int) int {
	count := 0
	currentTime := uint64(time.Now().Unix())
	cutoffTime := currentTime - uint64(daysOld*86400) // Convert days to seconds
	
	userNotifications := k.GetUserNotifications(ctx, userAddress)
	
	for _, notification := range userNotifications {
		if notification.CreatedAt < cutoffTime && notification.IsRead {
			k.RemoveNotification(ctx, notification.Index)
			count++
		}
	}
	
	// Emit cleanup event
	if count > 0 {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"notifications_bulk_cleanup",
				sdk.NewAttribute("user_address", userAddress),
				sdk.NewAttribute("count", string(rune(count))),
			),
		)
	}
	
	return count
}

// GetNotificationStatistics returns comprehensive notification statistics
func (k Keeper) GetNotificationStatistics(ctx sdk.Context, userAddress string) map[string]interface{} {
	userNotifications := k.GetUserNotifications(ctx, userAddress)
	
	stats := make(map[string]interface{})
	stats["total"] = len(userNotifications)
	
	// Count by read status
	unreadCount := 0
	readCount := 0
	
	// Count by priority
	priorityCounts := make(map[string]int)
	
	// Count by type
	typeCounts := make(map[string]int)
	
	// Count by source
	sourceCounts := make(map[string]int)
	
	for _, notification := range userNotifications {
		if notification.IsRead {
			readCount++
		} else {
			unreadCount++
		}
		
		priorityCounts[notification.Priority]++
		typeCounts[notification.NotificationType]++
		sourceCounts[notification.SourceModule]++
	}
	
	stats["unread"] = unreadCount
	stats["read"] = readCount
	stats["priorities"] = priorityCounts
	stats["types"] = typeCounts
	stats["sources"] = sourceCounts
	
	return stats
}

// CanUserReceiveNotification checks if user can receive a specific notification type
func (k Keeper) CanUserReceiveNotification(ctx sdk.Context, userAddress, notificationType, priority string) bool {
	// Get user settings
	allSettings := k.GetAllNotificationSettings(ctx)
	var userSettings *types.NotificationSettings
	
	for _, setting := range allSettings {
		if setting.UserAddress == userAddress {
			userSettings = &setting
			break
		}
	}
	
	// If no settings, allow by default
	if userSettings == nil {
		return true
	}
	
	// Check if notification type is allowed
	if len(userSettings.NotificationTypes) > 0 {
		typeAllowed := false
		for _, allowedType := range userSettings.NotificationTypes {
			if strings.EqualFold(allowedType, notificationType) {
				typeAllowed = true
				break
			}
		}
		if !typeAllowed {
			return false
		}
	}
	
	// Always allow critical/urgent notifications
	if strings.EqualFold(priority, "critical") || strings.EqualFold(priority, "urgent") {
		return true
	}
	
	// Check frequency setting
	if strings.EqualFold(userSettings.Frequency, "never") {
		return false
	}
	
	return true
}

// GetUserNotificationSettings returns notification settings for a user
func (k Keeper) GetUserNotificationSettings(ctx sdk.Context, userAddress string) *types.NotificationSettings {
	allSettings := k.GetAllNotificationSettings(ctx)
	for _, setting := range allSettings {
		if setting.UserAddress == userAddress {
			return &setting
		}
	}
	return nil
}
