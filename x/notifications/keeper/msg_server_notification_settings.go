package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"time"

	"skillchain/x/notifications/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateNotificationSettings(goCtx context.Context, msg *types.MsgCreateNotificationSettings) (*types.MsgCreateNotificationSettingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Enhanced validation for notification settings
	if err := k.validateNotificationSettingsData(msg); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Generate settings index if not provided
	settingsIndex := msg.Index
	if settingsIndex == "" {
		settingsIndex = k.generateSettingsIndex(msg.UserAddress, msg.Creator)
	}

	// Check if settings already exist for this user
	existingSettings := k.getUserNotificationSettingsByUser(ctx, msg.UserAddress)
	if existingSettings != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "notification settings already exist for this user")
	}

	// Check if index already exists
	_, isFound := k.GetNotificationSettings(ctx, settingsIndex)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "settings already exist with this index")
	}

	// Set smart defaults and normalize data
	frequency := k.normalizeFrequency(msg.Frequency)
	notificationTypes := k.normalizeNotificationTypes(msg.NotificationTypes)

	var notificationSettings = types.NotificationSettings{
		Creator:           msg.Creator,
		Index:             settingsIndex,
		UserAddress:       msg.UserAddress,
		EmailEnabled:      msg.EmailEnabled,
		PushEnabled:       msg.PushEnabled,
		SmsEnabled:        msg.SmsEnabled,
		EmailAddress:      strings.ToLower(strings.TrimSpace(msg.EmailAddress)),
		PhoneNumber:       k.normalizePhoneNumber(msg.PhoneNumber),
		NotificationTypes: notificationTypes,
		Frequency:         frequency,
	}

	k.SetNotificationSettings(ctx, notificationSettings)

	// Emit settings created event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"notification_settings_created",
			sdk.NewAttribute("creator", msg.Creator),
			sdk.NewAttribute("user_address", msg.UserAddress),
			sdk.NewAttribute("settings_index", settingsIndex),
			sdk.NewAttribute("email_enabled", fmt.Sprintf("%t", msg.EmailEnabled)),
			sdk.NewAttribute("push_enabled", fmt.Sprintf("%t", msg.PushEnabled)),
			sdk.NewAttribute("sms_enabled", fmt.Sprintf("%t", msg.SmsEnabled)),
			sdk.NewAttribute("frequency", frequency),
		),
	)

	return &types.MsgCreateNotificationSettingsResponse{}, nil
}

func (k msgServer) UpdateNotificationSettings(goCtx context.Context, msg *types.MsgUpdateNotificationSettings) (*types.MsgUpdateNotificationSettingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the settings exist
	valFound, isFound := k.GetNotificationSettings(ctx, msg.Index)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "notification settings not found")
	}

	// Enhanced authorization - only user or original creator can update
	if msg.Creator != valFound.Creator && msg.Creator != valFound.UserAddress {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only settings creator or user can update")
	}

	// Validate update data
	if err := k.validateNotificationSettingsUpdateData(msg); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Normalize and validate update data
	frequency := k.normalizeFrequency(msg.Frequency)
	notificationTypes := k.normalizeNotificationTypes(msg.NotificationTypes)

	var notificationSettings = types.NotificationSettings{
		Creator:           valFound.Creator,        // Preserve original creator
		Index:             msg.Index,
		UserAddress:       valFound.UserAddress,   // Don't allow user change
		EmailEnabled:      msg.EmailEnabled,       // Allow email toggle
		PushEnabled:       msg.PushEnabled,        // Allow push toggle
		SmsEnabled:        msg.SmsEnabled,         // Allow SMS toggle
		EmailAddress:      strings.ToLower(strings.TrimSpace(msg.EmailAddress)), // Normalize email
		PhoneNumber:       k.normalizePhoneNumber(msg.PhoneNumber),              // Normalize phone
		NotificationTypes: notificationTypes,      // Allow type changes
		Frequency:         frequency,              // Allow frequency changes
	}

	k.SetNotificationSettings(ctx, notificationSettings)

	// Emit settings updated event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"notification_settings_updated",
			sdk.NewAttribute("updater", msg.Creator),
			sdk.NewAttribute("user_address", valFound.UserAddress),
			sdk.NewAttribute("settings_index", msg.Index),
			sdk.NewAttribute("email_enabled", fmt.Sprintf("%t", msg.EmailEnabled)),
			sdk.NewAttribute("push_enabled", fmt.Sprintf("%t", msg.PushEnabled)),
			sdk.NewAttribute("sms_enabled", fmt.Sprintf("%t", msg.SmsEnabled)),
			sdk.NewAttribute("frequency", frequency),
		),
	)

	return &types.MsgUpdateNotificationSettingsResponse{}, nil
}

func (k msgServer) DeleteNotificationSettings(goCtx context.Context, msg *types.MsgDeleteNotificationSettings) (*types.MsgDeleteNotificationSettingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the settings exist
	valFound, isFound := k.GetNotificationSettings(ctx, msg.Index)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "notification settings not found")
	}

	// Enhanced authorization - only user or original creator can delete
	if msg.Creator != valFound.Creator && msg.Creator != valFound.UserAddress {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only settings creator or user can delete")
	}

	k.RemoveNotificationSettings(ctx, msg.Index)

	// Emit settings deleted event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"notification_settings_deleted",
			sdk.NewAttribute("deleter", msg.Creator),
			sdk.NewAttribute("user_address", valFound.UserAddress),
			sdk.NewAttribute("settings_index", msg.Index),
		),
	)

	return &types.MsgDeleteNotificationSettingsResponse{}, nil
}

// validateNotificationSettingsData performs comprehensive validation
func (k msgServer) validateNotificationSettingsData(msg *types.MsgCreateNotificationSettings) error {
	if msg.UserAddress == "" {
		return fmt.Errorf("user address cannot be empty")
	}

	// Validate email if email notifications are enabled
	if msg.EmailEnabled && msg.EmailAddress != "" {
		if !k.isValidEmail(msg.EmailAddress) {
			return fmt.Errorf("invalid email address format")
		}
	}

	// Validate phone if SMS notifications are enabled
	if msg.SmsEnabled && msg.PhoneNumber != "" {
		if !k.isValidPhoneNumber(msg.PhoneNumber) {
			return fmt.Errorf("invalid phone number format")
		}
	}

	// Validate frequency
	if msg.Frequency != "" && !k.isValidFrequency(msg.Frequency) {
		return fmt.Errorf("invalid frequency setting")
	}

	// Validate notification types
	for _, notificationType := range msg.NotificationTypes {
		if !k.isValidNotificationType(notificationType) {
			return fmt.Errorf("invalid notification type: %s", notificationType)
		}
	}

	return nil
}

// validateNotificationSettingsUpdateData validates update-specific data
func (k msgServer) validateNotificationSettingsUpdateData(msg *types.MsgUpdateNotificationSettings) error {
	// Validate email if email notifications are enabled
	if msg.EmailEnabled && msg.EmailAddress != "" {
		if !k.isValidEmail(msg.EmailAddress) {
			return fmt.Errorf("invalid email address format")
		}
	}

	// Validate phone if SMS notifications are enabled
	if msg.SmsEnabled && msg.PhoneNumber != "" {
		if !k.isValidPhoneNumber(msg.PhoneNumber) {
			return fmt.Errorf("invalid phone number format")
		}
	}

	// Validate frequency
	if msg.Frequency != "" && !k.isValidFrequency(msg.Frequency) {
		return fmt.Errorf("invalid frequency setting")
	}

	// Validate notification types
	for _, notificationType := range msg.NotificationTypes {
		if !k.isValidNotificationType(notificationType) {
			return fmt.Errorf("invalid notification type: %s", notificationType)
		}
	}

	return nil
}

// generateSettingsIndex creates a unique settings identifier
func (k msgServer) generateSettingsIndex(userAddress, creator string) string {
	data := fmt.Sprintf("%s:%s:%d", userAddress, creator, time.Now().Unix())
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])[:16]
}

// isValidEmail validates email address format
func (k msgServer) isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// isValidPhoneNumber validates phone number format
func (k msgServer) isValidPhoneNumber(phone string) bool {
	// Basic phone validation - accepts international formats
	phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	cleanPhone := regexp.MustCompile(`[^0-9+]`).ReplaceAllString(phone, "")
	return phoneRegex.MatchString(cleanPhone)
}

// isValidFrequency validates frequency settings
func (k msgServer) isValidFrequency(frequency string) bool {
	validFrequencies := []string{
		"immediate", "hourly", "daily", "weekly", "monthly", "never",
	}
	
	for _, validFreq := range validFrequencies {
		if strings.EqualFold(frequency, validFreq) {
			return true
		}
	}
	return false
}

// normalizeFrequency converts frequency to lowercase with default
func (k msgServer) normalizeFrequency(frequency string) string {
	if frequency == "" {
		return "immediate"
	}
	return strings.ToLower(frequency)
}

// normalizeNotificationTypes ensures valid notification types
func (k msgServer) normalizeNotificationTypes(types []string) []string {
	var normalized []string
	seen := make(map[string]bool)
	
	for _, notificationType := range types {
		lower := strings.ToLower(notificationType)
		if k.isValidNotificationType(lower) && !seen[lower] {
			normalized = append(normalized, lower)
			seen[lower] = true
		}
	}
	
	// Add default types if none specified
	if len(normalized) == 0 {
		normalized = []string{"system", "security", "transaction"}
	}
	
	return normalized
}

// normalizePhoneNumber cleans and formats phone number
func (k msgServer) normalizePhoneNumber(phone string) string {
	if phone == "" {
		return phone
	}
	
	// Remove all non-digit characters except +
	cleaned := regexp.MustCompile(`[^0-9+]`).ReplaceAllString(phone, "")
	
	// Ensure it starts with + for international format
	if len(cleaned) > 0 && !strings.HasPrefix(cleaned, "+") {
		cleaned = "+" + cleaned
	}
	
	return cleaned
}

// getUserNotificationSettingsByUser retrieves settings by user address
func (k msgServer) getUserNotificationSettingsByUser(ctx sdk.Context, userAddress string) *types.NotificationSettings {
	allSettings := k.GetAllNotificationSettings(ctx)
	for _, setting := range allSettings {
		if setting.UserAddress == userAddress {
			return &setting
		}
	}
	return nil
}
