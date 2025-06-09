package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"skillchain/x/filestorage/types"
)

var _ types.QueryServer = Keeper{}

// CheckFileAccess verifies if a user has access to a file
func (k Keeper) CheckFileAccess(ctx sdk.Context, fileId, userAddress string, requiredLevel string) bool {
	// Get file record
	fileRecord, exists := k.GetFileRecord(ctx, fileId)
	if !exists {
		return false
	}
	
	// File is public and only read access required
	if fileRecord.IsPublic && requiredLevel == "read" {
		return true
	}
	
	// Owner and creator have full access
	if userAddress == fileRecord.Owner || userAddress == fileRecord.Creator {
		return true
	}
	
	// Check specific permissions
	return k.hasPermissionLevel(ctx, fileId, userAddress, requiredLevel)
}

// hasPermissionLevel checks if user has specific permission level
func (k Keeper) hasPermissionLevel(ctx sdk.Context, fileId, userAddress, requiredLevel string) bool {
	allPermissions := k.GetAllFilePermission(ctx)
	
	for _, permission := range allPermissions {
		if permission.FileId == fileId && permission.UserAddress == userAddress {
			// Check if permission is not expired
			if permission.ExpiresAt != 0 && permission.ExpiresAt <= uint64(time.Now().Unix()) {
				continue
			}
			
			// Check permission level hierarchy
			return k.isPermissionSufficient(permission.PermissionLevel, requiredLevel)
		}
	}
	
	return false
}

// isPermissionSufficient checks if given permission level satisfies required level
func (k Keeper) isPermissionSufficient(givenLevel, requiredLevel string) bool {
	permissionHierarchy := map[string]int{
		"read":  1,
		"write": 2,
		"admin": 3,
		"owner": 4,
	}
	
	given, okGiven := permissionHierarchy[givenLevel]
	required, okRequired := permissionHierarchy[requiredLevel]
	
	if !okGiven || !okRequired {
		return false
	}
	
	return given >= required
}

// GetUserFiles returns all files owned by or accessible to a user
func (k Keeper) GetUserFiles(ctx sdk.Context, userAddress string) []types.FileRecord {
	var userFiles []types.FileRecord
	allFiles := k.GetAllFileRecord(ctx)
	
	for _, file := range allFiles {
		// Add if user is owner, creator, or has public access
		if file.Owner == userAddress || file.Creator == userAddress || 
		   (file.IsPublic) || k.hasPermissionLevel(ctx, file.Index, userAddress, "read") {
			userFiles = append(userFiles, file)
		}
	}
	
	return userFiles
}

// GetFilesByPermissionLevel returns files where user has specific permission level
func (k Keeper) GetFilesByPermissionLevel(ctx sdk.Context, userAddress, permissionLevel string) []types.FileRecord {
	var files []types.FileRecord
	allFiles := k.GetAllFileRecord(ctx)
	
	for _, file := range allFiles {
		if k.CheckFileAccess(ctx, file.Index, userAddress, permissionLevel) {
			files = append(files, file)
		}
	}
	
	return files
}

// GetActivePermissions returns all non-expired permissions for a file
func (k Keeper) GetActivePermissions(ctx sdk.Context, fileId string) []types.FilePermission {
	var activePermissions []types.FilePermission
	allPermissions := k.GetAllFilePermission(ctx)
	currentTime := uint64(time.Now().Unix())
	
	for _, permission := range allPermissions {
		if permission.FileId == fileId {
			// Include if not expired (ExpiresAt = 0 means no expiration)
			if permission.ExpiresAt == 0 || permission.ExpiresAt > currentTime {
				activePermissions = append(activePermissions, permission)
			}
		}
	}
	
	return activePermissions
}
