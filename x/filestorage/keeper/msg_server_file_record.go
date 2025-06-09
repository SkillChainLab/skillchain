package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"skillchain/x/filestorage/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateFileRecord(goCtx context.Context, msg *types.MsgCreateFileRecord) (*types.MsgCreateFileRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Enhanced Input Validation
	if err := k.validateFileRecord(msg); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Check if the value already exists
	_, isFound := k.GetFileRecord(ctx, msg.Index)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "file already exists with this index")
	}

	// Auto-generate upload date if not provided
	uploadDate := msg.UploadDate
	if uploadDate == 0 {
		uploadDate = uint64(time.Now().Unix())
	}

	// Generate unique file ID if index is empty
	fileIndex := msg.Index
	if fileIndex == "" {
		fileIndex = k.generateFileIndex(msg.Creator, msg.Filename, uploadDate)
	}

	// Verify file hash integrity
	if msg.FileHash != "" && !k.isValidFileHash(msg.FileHash) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid file hash format")
	}

	// Create file record with enhanced data
	var fileRecord = types.FileRecord{
		Creator:     msg.Creator,
		Index:       fileIndex,
		Owner:       msg.Owner,
		Filename:    msg.Filename,
		FileHash:    msg.FileHash,
		FileSize:    msg.FileSize,
		ContentType: msg.ContentType,
		UploadDate:  uploadDate,
		IpfsHash:    msg.IpfsHash,
		Metadata:    msg.Metadata,
		IsPublic:    msg.IsPublic,
	}

	k.SetFileRecord(ctx, fileRecord)

	// Emit event for file creation
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"file_created",
			sdk.NewAttribute("creator", msg.Creator),
			sdk.NewAttribute("file_index", fileIndex),
			sdk.NewAttribute("filename", msg.Filename),
			sdk.NewAttribute("file_size", fmt.Sprintf("%d", msg.FileSize)),
			sdk.NewAttribute("is_public", fmt.Sprintf("%t", msg.IsPublic)),
		),
	)

	return &types.MsgCreateFileRecordResponse{}, nil
}

// validateFileRecord performs comprehensive validation on file record
func (k msgServer) validateFileRecord(msg *types.MsgCreateFileRecord) error {
	// Check filename
	if msg.Filename == "" {
		return fmt.Errorf("filename cannot be empty")
	}
	if len(msg.Filename) > 255 {
		return fmt.Errorf("filename too long (max 255 characters)")
	}
	
	// Check file size limits (max 100GB)
	if msg.FileSize > 100*1024*1024*1024 {
		return fmt.Errorf("file size exceeds maximum limit (100GB)")
	}
	
	// Validate content type
	if msg.ContentType != "" && !k.isValidContentType(msg.ContentType) {
		return fmt.Errorf("invalid content type")
	}
	
	// Validate IPFS hash if provided
	if msg.IpfsHash != "" && !k.isValidIPFSHash(msg.IpfsHash) {
		return fmt.Errorf("invalid IPFS hash format")
	}
	
	// Owner validation
	if msg.Owner == "" {
		return fmt.Errorf("owner address cannot be empty")
	}
	
	return nil
}

// generateFileIndex creates a unique file identifier
func (k msgServer) generateFileIndex(creator, filename string, uploadDate uint64) string {
	data := fmt.Sprintf("%s:%s:%d", creator, filename, uploadDate)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])[:16] // Use first 16 chars for readability
}

// isValidFileHash checks if file hash is valid (SHA256)
func (k msgServer) isValidFileHash(hash string) bool {
	if len(hash) != 64 {
		return false
	}
	_, err := hex.DecodeString(hash)
	return err == nil
}

// isValidContentType validates MIME types
func (k msgServer) isValidContentType(contentType string) bool {
	validTypes := []string{
		"text/", "image/", "video/", "audio/", "application/",
		"multipart/", "message/", "model/", "font/",
	}
	
	for _, validType := range validTypes {
		if strings.HasPrefix(contentType, validType) {
			return true
		}
	}
	return false
}

// isValidIPFSHash validates IPFS hash format
func (k msgServer) isValidIPFSHash(hash string) bool {
	// Basic IPFS hash validation (starts with Qm and proper length)
	if strings.HasPrefix(hash, "Qm") && len(hash) == 46 {
		return true
	}
	// Also support newer CIDv1 format
	if strings.HasPrefix(hash, "bafy") && len(hash) >= 59 {
		return true
	}
	return false
}

func (k msgServer) UpdateFileRecord(goCtx context.Context, msg *types.MsgUpdateFileRecord) (*types.MsgUpdateFileRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetFileRecord(ctx, msg.Index)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "file record not found")
	}

	// Enhanced authorization check - allow owner or creator to update
	if msg.Creator != valFound.Creator && msg.Creator != valFound.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only file owner or creator can update")
	}

	// Validate update data
	if err := k.validateFileUpdateData(msg); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Preserve original creation data but allow updates to specific fields
	var fileRecord = types.FileRecord{
		Creator:     valFound.Creator,     // Preserve original creator
		Index:       msg.Index,
		Owner:       msg.Owner,           // Allow ownership transfer
		Filename:    msg.Filename,        // Allow filename updates
		FileHash:    msg.FileHash,        // Allow hash updates (for file versions)
		FileSize:    msg.FileSize,        // Allow size updates
		ContentType: msg.ContentType,     // Allow content type updates
		UploadDate:  valFound.UploadDate, // Preserve original upload date
		IpfsHash:    msg.IpfsHash,        // Allow IPFS hash updates
		Metadata:    msg.Metadata,        // Allow metadata updates
		IsPublic:    msg.IsPublic,        // Allow visibility changes
	}

	k.SetFileRecord(ctx, fileRecord)

	// Emit update event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"file_updated",
			sdk.NewAttribute("updater", msg.Creator),
			sdk.NewAttribute("file_index", msg.Index),
			sdk.NewAttribute("new_owner", msg.Owner),
			sdk.NewAttribute("is_public", fmt.Sprintf("%t", msg.IsPublic)),
		),
	)

	return &types.MsgUpdateFileRecordResponse{}, nil
}

func (k msgServer) DeleteFileRecord(goCtx context.Context, msg *types.MsgDeleteFileRecord) (*types.MsgDeleteFileRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetFileRecord(ctx, msg.Index)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "file record not found")
	}

	// Enhanced authorization - only creator or owner can delete
	if msg.Creator != valFound.Creator && msg.Creator != valFound.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only file owner or creator can delete")
	}

	// Check if file has active permissions that need cleanup
	permissions := k.GetFilePermissions(ctx, msg.Index)
	if len(permissions) > 0 {
		// Remove all associated permissions first
		for _, permission := range permissions {
			k.RemoveFilePermission(ctx, permission.Index)
		}
	}

	// Remove the file record
	k.RemoveFileRecord(ctx, msg.Index)

	// Emit deletion event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"file_deleted",
			sdk.NewAttribute("deleter", msg.Creator),
			sdk.NewAttribute("file_index", msg.Index),
			sdk.NewAttribute("filename", valFound.Filename),
			sdk.NewAttribute("permissions_cleaned", fmt.Sprintf("%d", len(permissions))),
		),
	)

	return &types.MsgDeleteFileRecordResponse{}, nil
}

// validateFileUpdateData validates update-specific data
func (k msgServer) validateFileUpdateData(msg *types.MsgUpdateFileRecord) error {
	// Basic validation similar to create
	if msg.Filename == "" {
		return fmt.Errorf("filename cannot be empty")
	}
	if len(msg.Filename) > 255 {
		return fmt.Errorf("filename too long (max 255 characters)")
	}
	if msg.FileSize > 100*1024*1024*1024 {
		return fmt.Errorf("file size exceeds maximum limit (100GB)")
	}
	if msg.ContentType != "" && !k.isValidContentType(msg.ContentType) {
		return fmt.Errorf("invalid content type")
	}
	if msg.IpfsHash != "" && !k.isValidIPFSHash(msg.IpfsHash) {
		return fmt.Errorf("invalid IPFS hash format")
	}
	if msg.Owner == "" {
		return fmt.Errorf("owner address cannot be empty")
	}
	return nil
}

// GetFilePermissions retrieves all permissions for a specific file
func (k msgServer) GetFilePermissions(ctx sdk.Context, fileId string) []types.FilePermission {
	var permissions []types.FilePermission
	
	// Get all file permissions and filter by file ID
	allPermissions := k.GetAllFilePermission(ctx)
	for _, permission := range allPermissions {
		if permission.FileId == fileId {
			permissions = append(permissions, permission)
		}
	}
	
	return permissions
}
