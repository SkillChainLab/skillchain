package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"skillchain/x/filestorage/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateFilePermission(goCtx context.Context, msg *types.MsgCreateFilePermission) (*types.MsgCreateFilePermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Verify that the file exists
	fileRecord, fileExists := k.GetFileRecord(ctx, msg.FileId)
	if !fileExists {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "file does not exist")
	}

	// Authorization check - only file owner or creator can grant permissions
	if msg.Creator != fileRecord.Owner && msg.Creator != fileRecord.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only file owner or creator can grant permissions")
	}

	// Validate permission data
	if err := k.validatePermissionData(msg); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Generate permission index if not provided
	permissionIndex := msg.Index
	if permissionIndex == "" {
		permissionIndex = k.generatePermissionIndex(msg.FileId, msg.UserAddress, msg.GrantedBy)
	}

	// Check if permission already exists for this user and file
	if k.hasExistingPermission(ctx, msg.FileId, msg.UserAddress) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "permission already exists for this user and file")
	}

	// Auto-set timestamp if not provided
	grantedAt := msg.GrantedAt
	if grantedAt == 0 {
		grantedAt = uint64(time.Now().Unix())
	}

	// Validate expiration date
	if msg.ExpiresAt != 0 && msg.ExpiresAt <= grantedAt {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "expiration date must be after grant date")
	}

	var filePermission = types.FilePermission{
		Creator:         msg.Creator,
		Index:           permissionIndex,
		FileId:          msg.FileId,
		UserAddress:     msg.UserAddress,
		PermissionLevel: msg.PermissionLevel,
		GrantedBy:       msg.GrantedBy,
		GrantedAt:       grantedAt,
		ExpiresAt:       msg.ExpiresAt,
	}

	k.SetFilePermission(ctx, filePermission)

	// Emit permission granted event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"permission_granted",
			sdk.NewAttribute("granter", msg.Creator),
			sdk.NewAttribute("grantee", msg.UserAddress),
			sdk.NewAttribute("file_id", msg.FileId),
			sdk.NewAttribute("permission_level", msg.PermissionLevel),
			sdk.NewAttribute("expires_at", fmt.Sprintf("%d", msg.ExpiresAt)),
		),
	)

	return &types.MsgCreateFilePermissionResponse{}, nil
}

func (k msgServer) UpdateFilePermission(goCtx context.Context, msg *types.MsgUpdateFilePermission) (*types.MsgUpdateFilePermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the permission exists
	valFound, isFound := k.GetFilePermission(ctx, msg.Index)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "permission not found")
	}

	// Verify file still exists
	fileRecord, fileExists := k.GetFileRecord(ctx, valFound.FileId)
	if !fileExists {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "associated file no longer exists")
	}

	// Authorization check - only file owner, creator, or permission granter can modify
	if msg.Creator != fileRecord.Owner && msg.Creator != fileRecord.Creator && msg.Creator != valFound.GrantedBy {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "insufficient permissions to modify")
	}

	// Validate updated permission data
	if err := k.validatePermissionUpdateData(msg); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	var filePermission = types.FilePermission{
		Creator:         valFound.Creator,         // Preserve original creator
		Index:           msg.Index,
		FileId:          valFound.FileId,          // Don't allow file ID changes
		UserAddress:     valFound.UserAddress,    // Don't allow user changes
		PermissionLevel: msg.PermissionLevel,     // Allow permission level changes
		GrantedBy:       valFound.GrantedBy,      // Preserve original granter
		GrantedAt:       valFound.GrantedAt,      // Preserve original grant time
		ExpiresAt:       msg.ExpiresAt,           // Allow expiration updates
	}

	k.SetFilePermission(ctx, filePermission)

	// Emit permission updated event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"permission_updated",
			sdk.NewAttribute("updater", msg.Creator),
			sdk.NewAttribute("grantee", valFound.UserAddress),
			sdk.NewAttribute("file_id", valFound.FileId),
			sdk.NewAttribute("new_permission_level", msg.PermissionLevel),
			sdk.NewAttribute("new_expires_at", fmt.Sprintf("%d", msg.ExpiresAt)),
		),
	)

	return &types.MsgUpdateFilePermissionResponse{}, nil
}

func (k msgServer) DeleteFilePermission(goCtx context.Context, msg *types.MsgDeleteFilePermission) (*types.MsgDeleteFilePermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the permission exists
	valFound, isFound := k.GetFilePermission(ctx, msg.Index)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "permission not found")
	}

	// Get file record for authorization
	fileRecord, fileExists := k.GetFileRecord(ctx, valFound.FileId)
	
	// Authorization check - allow deletion by file owner, creator, granter, or grantee
	canDelete := msg.Creator == valFound.Creator ||
		msg.Creator == valFound.GrantedBy ||
		msg.Creator == valFound.UserAddress ||
		(fileExists && (msg.Creator == fileRecord.Owner || msg.Creator == fileRecord.Creator))

	if !canDelete {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "insufficient permissions to revoke")
	}

	k.RemoveFilePermission(ctx, msg.Index)

	// Emit permission revoked event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"permission_revoked",
			sdk.NewAttribute("revoker", msg.Creator),
			sdk.NewAttribute("grantee", valFound.UserAddress),
			sdk.NewAttribute("file_id", valFound.FileId),
			sdk.NewAttribute("permission_level", valFound.PermissionLevel),
		),
	)

	return &types.MsgDeleteFilePermissionResponse{}, nil
}

// validatePermissionData validates permission creation data
func (k msgServer) validatePermissionData(msg *types.MsgCreateFilePermission) error {
	if msg.FileId == "" {
		return fmt.Errorf("file ID cannot be empty")
	}
	if msg.UserAddress == "" {
		return fmt.Errorf("user address cannot be empty")
	}
	if msg.GrantedBy == "" {
		return fmt.Errorf("granter address cannot be empty")
	}
	
	// Validate permission level
	validLevels := []string{"read", "write", "admin", "owner"}
	isValidLevel := false
	for _, level := range validLevels {
		if msg.PermissionLevel == level {
			isValidLevel = true
			break
		}
	}
	if !isValidLevel {
		return fmt.Errorf("invalid permission level. Must be one of: read, write, admin, owner")
	}
	
	return nil
}

// validatePermissionUpdateData validates permission update data
func (k msgServer) validatePermissionUpdateData(msg *types.MsgUpdateFilePermission) error {
	// Validate permission level
	validLevels := []string{"read", "write", "admin", "owner"}
	isValidLevel := false
	for _, level := range validLevels {
		if msg.PermissionLevel == level {
			isValidLevel = true
			break
		}
	}
	if !isValidLevel {
		return fmt.Errorf("invalid permission level. Must be one of: read, write, admin, owner")
	}
	
	return nil
}

// hasExistingPermission checks if a permission already exists for user and file
func (k msgServer) hasExistingPermission(ctx sdk.Context, fileId, userAddress string) bool {
	allPermissions := k.GetAllFilePermission(ctx)
	for _, permission := range allPermissions {
		if permission.FileId == fileId && permission.UserAddress == userAddress {
			// Check if permission is not expired
			if permission.ExpiresAt == 0 || permission.ExpiresAt > uint64(time.Now().Unix()) {
				return true
			}
		}
	}
	return false
}

// generatePermissionIndex creates a unique permission identifier
func (k msgServer) generatePermissionIndex(fileId, userAddress, grantedBy string) string {
	data := fmt.Sprintf("%s:%s:%s:%d", fileId, userAddress, grantedBy, time.Now().Unix())
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])[:16]
}
