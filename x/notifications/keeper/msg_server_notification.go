package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"skillchain/x/notifications/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateNotification(goCtx context.Context, msg *types.MsgCreateNotification) (*types.MsgCreateNotificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Enhanced Input Validation
	if err := k.validateNotificationData(msg); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Generate unique notification index if not provided
	notificationIndex := msg.Index
	if notificationIndex == "" {
		notificationIndex = k.generateNotificationIndex(msg.UserAddress, msg.Title, msg.CreatedAt)
	}

	// Check if notification already exists
	_, isFound := k.GetNotification(ctx, notificationIndex)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "notification already exists with this index")
	}

	// Auto-generate creation timestamp if not provided
	createdAt := msg.CreatedAt
	if createdAt == 0 {
		createdAt = uint64(time.Now().Unix())
	}

	// Set default priority if empty
	priority := msg.Priority
	if priority == "" {
		priority = "medium"
	}

	// Validate and normalize priority
	priority = k.normalizePriority(priority)

	// Check user notification preferences
	if !k.shouldDeliverNotification(ctx, msg.UserAddress, msg.NotificationType, priority) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "notification blocked by user preferences")
	}

	var notification = types.Notification{
		Creator:          msg.Creator,
		Index:            notificationIndex,
		UserAddress:      msg.UserAddress,
		NotificationType: msg.NotificationType,
		Title:            msg.Title,
		Message:          msg.Message,
		Data:             msg.Data,
		IsRead:           false, // Always start as unread
		CreatedAt:        createdAt,
		Priority:         priority,
		SourceModule:     msg.SourceModule,
		SourceAction:     msg.SourceAction,
	}

	k.SetNotification(ctx, notification)

	// Emit notification created event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"notification_created",
			sdk.NewAttribute("creator", msg.Creator),
			sdk.NewAttribute("user_address", msg.UserAddress),
			sdk.NewAttribute("notification_index", notificationIndex),
			sdk.NewAttribute("notification_type", msg.NotificationType),
			sdk.NewAttribute("priority", priority),
			sdk.NewAttribute("source_module", msg.SourceModule),
		),
	)

	// Auto-increment unread counter
	k.incrementUnreadCount(ctx, msg.UserAddress)

	return &types.MsgCreateNotificationResponse{}, nil
}

func (k msgServer) UpdateNotification(goCtx context.Context, msg *types.MsgUpdateNotification) (*types.MsgUpdateNotificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetNotification(ctx, msg.Index)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "notification not found")
	}

	// Enhanced authorization - allow user or creator to update
	if msg.Creator != valFound.Creator && msg.Creator != valFound.UserAddress {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only notification creator or recipient can update")
	}

	// Validate update data
	if err := k.validateNotificationUpdateData(msg); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Track read status change
	wasRead := valFound.IsRead
	isNowRead := msg.IsRead

	var notification = types.Notification{
		Creator:          valFound.Creator,        // Preserve original creator
		Index:            msg.Index,
		UserAddress:      valFound.UserAddress,   // Don't allow user change
		NotificationType: msg.NotificationType,   // Allow type updates
		Title:            msg.Title,              // Allow title updates
		Message:          msg.Message,            // Allow message updates
		Data:             msg.Data,               // Allow data updates
		IsRead:           msg.IsRead,             // Allow read status updates
		CreatedAt:        valFound.CreatedAt,     // Preserve creation time
		Priority:         k.normalizePriority(msg.Priority), // Normalize priority
		SourceModule:     valFound.SourceModule,  // Preserve source
		SourceAction:     valFound.SourceAction,  // Preserve action
	}

	k.SetNotification(ctx, notification)

	// Handle read status change
	if !wasRead && isNowRead {
		k.decrementUnreadCount(ctx, valFound.UserAddress)
		
		// Emit notification read event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"notification_read",
				sdk.NewAttribute("user_address", valFound.UserAddress),
				sdk.NewAttribute("notification_index", msg.Index),
				sdk.NewAttribute("priority", notification.Priority),
			),
		)
	}

	// Emit notification updated event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"notification_updated",
			sdk.NewAttribute("updater", msg.Creator),
			sdk.NewAttribute("notification_index", msg.Index),
			sdk.NewAttribute("user_address", valFound.UserAddress),
		),
	)

	return &types.MsgUpdateNotificationResponse{}, nil
}

func (k msgServer) DeleteNotification(goCtx context.Context, msg *types.MsgDeleteNotification) (*types.MsgDeleteNotificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetNotification(ctx, msg.Index)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "notification not found")
	}

	// Enhanced authorization - allow user or creator to delete
	if msg.Creator != valFound.Creator && msg.Creator != valFound.UserAddress {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only notification creator or recipient can delete")
	}

	// If notification was unread, decrement counter
	if !valFound.IsRead {
		k.decrementUnreadCount(ctx, valFound.UserAddress)
	}

	k.RemoveNotification(ctx, msg.Index)

	// Emit notification deleted event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"notification_deleted",
			sdk.NewAttribute("deleter", msg.Creator),
			sdk.NewAttribute("notification_index", msg.Index),
			sdk.NewAttribute("user_address", valFound.UserAddress),
			sdk.NewAttribute("was_read", fmt.Sprintf("%t", valFound.IsRead)),
		),
	)

	return &types.MsgDeleteNotificationResponse{}, nil
}

// validateNotificationData performs comprehensive validation
func (k msgServer) validateNotificationData(msg *types.MsgCreateNotification) error {
	if msg.UserAddress == "" {
		return fmt.Errorf("user address cannot be empty")
	}
	if msg.Title == "" {
		return fmt.Errorf("notification title cannot be empty")
	}
	if len(msg.Title) > 200 {
		return fmt.Errorf("notification title too long (max 200 characters)")
	}
	if msg.Message == "" {
		return fmt.Errorf("notification message cannot be empty")
	}
	if len(msg.Message) > 2000 {
		return fmt.Errorf("notification message too long (max 2000 characters)")
	}
	if msg.NotificationType == "" {
		return fmt.Errorf("notification type cannot be empty")
	}
	if !k.isValidNotificationType(msg.NotificationType) {
		return fmt.Errorf("invalid notification type")
	}
	if msg.Priority != "" && !k.isValidPriority(msg.Priority) {
		return fmt.Errorf("invalid priority level")
	}
	return nil
}

// validateNotificationUpdateData validates update-specific data
func (k msgServer) validateNotificationUpdateData(msg *types.MsgUpdateNotification) error {
	if msg.Title == "" {
		return fmt.Errorf("notification title cannot be empty")
	}
	if len(msg.Title) > 200 {
		return fmt.Errorf("notification title too long (max 200 characters)")
	}
	if msg.Message == "" {
		return fmt.Errorf("notification message cannot be empty")
	}
	if len(msg.Message) > 2000 {
		return fmt.Errorf("notification message too long (max 2000 characters)")
	}
	if msg.NotificationType == "" {
		return fmt.Errorf("notification type cannot be empty")
	}
	if !k.isValidNotificationType(msg.NotificationType) {
		return fmt.Errorf("invalid notification type")
	}
	if msg.Priority != "" && !k.isValidPriority(msg.Priority) {
		return fmt.Errorf("invalid priority level")
	}
	return nil
}

// generateNotificationIndex creates a unique notification identifier
func (k msgServer) generateNotificationIndex(userAddress, title string, createdAt uint64) string {
	if createdAt == 0 {
		createdAt = uint64(time.Now().Unix())
	}
	data := fmt.Sprintf("%s:%s:%d", userAddress, title, createdAt)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])[:16]
}

// isValidNotificationType validates notification types
func (k msgServer) isValidNotificationType(notificationType string) bool {
	validTypes := []string{
		"system", "user", "marketing", "security", "transaction", 
		"marketplace", "filestorage", "profile", "analytics", 
		"reward", "warning", "error", "info", "success",
	}
	
	for _, validType := range validTypes {
		if strings.EqualFold(notificationType, validType) {
			return true
		}
	}
	return false
}

// isValidPriority validates priority levels
func (k msgServer) isValidPriority(priority string) bool {
	validPriorities := []string{"low", "medium", "high", "urgent", "critical"}
	
	for _, validPriority := range validPriorities {
		if strings.EqualFold(priority, validPriority) {
			return true
		}
	}
	return false
}

// normalizePriority converts priority to lowercase
func (k msgServer) normalizePriority(priority string) string {
	if priority == "" {
		return "medium"
	}
	return strings.ToLower(priority)
}

// shouldDeliverNotification checks user preferences and delivery rules
func (k msgServer) shouldDeliverNotification(ctx sdk.Context, userAddress, notificationType, priority string) bool {
	// Get user notification settings
	settings := k.getUserNotificationSettings(ctx, userAddress)
	if settings == nil {
		return true // Deliver if no settings (default allow)
	}

	// Check if this notification type is enabled
	if len(settings.NotificationTypes) > 0 {
		typeAllowed := false
		for _, allowedType := range settings.NotificationTypes {
			if strings.EqualFold(allowedType, notificationType) {
				typeAllowed = true
				break
			}
		}
		if !typeAllowed {
			return false
		}
	}

	// Always deliver critical/urgent notifications
	if priority == "critical" || priority == "urgent" {
		return true
	}

	// Check frequency settings
	if settings.Frequency == "never" {
		return false
	}

	return true
}

// getUserNotificationSettings retrieves user settings
func (k msgServer) getUserNotificationSettings(ctx sdk.Context, userAddress string) *types.NotificationSettings {
	allSettings := k.GetAllNotificationSettings(ctx)
	for _, setting := range allSettings {
		if setting.UserAddress == userAddress {
			return &setting
		}
	}
	return nil
}

// incrementUnreadCount increments user's unread notification counter
func (k msgServer) incrementUnreadCount(ctx sdk.Context, userAddress string) {
	// This would be implemented with a separate counter store in production
	// For now, we'll emit an event for tracking
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"unread_count_increment",
			sdk.NewAttribute("user_address", userAddress),
		),
	)
}

// decrementUnreadCount decrements user's unread notification counter
func (k msgServer) decrementUnreadCount(ctx sdk.Context, userAddress string) {
	// This would be implemented with a separate counter store in production
	// For now, we'll emit an event for tracking
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"unread_count_decrement",
			sdk.NewAttribute("user_address", userAddress),
		),
	)
}
