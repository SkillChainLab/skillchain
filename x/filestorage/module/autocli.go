package filestorage

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "skillchain/api/skillchain/filestorage"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "FileRecordAll",
					Use:       "list-file-record",
					Short:     "List all file-record",
				},
				{
					RpcMethod:      "FileRecord",
					Use:            "show-file-record [id]",
					Short:          "Shows a file-record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "FilePermissionAll",
					Use:       "list-file-permission",
					Short:     "List all file-permission",
				},
				{
					RpcMethod:      "FilePermission",
					Use:            "show-file-permission [id]",
					Short:          "Shows a file-permission",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateFileRecord",
					Use:            "create-file-record [index] [owner] [filename] [fileHash] [fileSize] [contentType] [uploadDate] [ipfsHash] [metadata] [isPublic]",
					Short:          "Create a new file-record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "owner"}, {ProtoField: "filename"}, {ProtoField: "fileHash"}, {ProtoField: "fileSize"}, {ProtoField: "contentType"}, {ProtoField: "uploadDate"}, {ProtoField: "ipfsHash"}, {ProtoField: "metadata"}, {ProtoField: "isPublic"}},
				},
				{
					RpcMethod:      "UpdateFileRecord",
					Use:            "update-file-record [index] [owner] [filename] [fileHash] [fileSize] [contentType] [uploadDate] [ipfsHash] [metadata] [isPublic]",
					Short:          "Update file-record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "owner"}, {ProtoField: "filename"}, {ProtoField: "fileHash"}, {ProtoField: "fileSize"}, {ProtoField: "contentType"}, {ProtoField: "uploadDate"}, {ProtoField: "ipfsHash"}, {ProtoField: "metadata"}, {ProtoField: "isPublic"}},
				},
				{
					RpcMethod:      "DeleteFileRecord",
					Use:            "delete-file-record [index]",
					Short:          "Delete file-record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateFilePermission",
					Use:            "create-file-permission [index] [fileId] [userAddress] [permissionLevel] [grantedBy] [grantedAt] [expiresAt]",
					Short:          "Create a new file-permission",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "fileId"}, {ProtoField: "userAddress"}, {ProtoField: "permissionLevel"}, {ProtoField: "grantedBy"}, {ProtoField: "grantedAt"}, {ProtoField: "expiresAt"}},
				},
				{
					RpcMethod:      "UpdateFilePermission",
					Use:            "update-file-permission [index] [fileId] [userAddress] [permissionLevel] [grantedBy] [grantedAt] [expiresAt]",
					Short:          "Update file-permission",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "fileId"}, {ProtoField: "userAddress"}, {ProtoField: "permissionLevel"}, {ProtoField: "grantedBy"}, {ProtoField: "grantedAt"}, {ProtoField: "expiresAt"}},
				},
				{
					RpcMethod:      "DeleteFilePermission",
					Use:            "delete-file-permission [index]",
					Short:          "Delete file-permission",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
