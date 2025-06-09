package notifications

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "skillchain/api/skillchain/notifications"
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
					RpcMethod: "NotificationAll",
					Use:       "list-notification",
					Short:     "List all notification",
				},
				{
					RpcMethod:      "Notification",
					Use:            "show-notification [id]",
					Short:          "Shows a notification",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "NotificationSettingsAll",
					Use:       "list-notification-settings",
					Short:     "List all notification-settings",
				},
				{
					RpcMethod:      "NotificationSettings",
					Use:            "show-notification-settings [id]",
					Short:          "Shows a notification-settings",
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
					RpcMethod:      "CreateNotification",
					Use:            "create-notification [index] [userAddress] [notificationType] [title] [message] [data] [isRead] [createdAt] [priority] [sourceModule] [sourceAction]",
					Short:          "Create a new notification",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "userAddress"}, {ProtoField: "notificationType"}, {ProtoField: "title"}, {ProtoField: "message"}, {ProtoField: "data"}, {ProtoField: "isRead"}, {ProtoField: "createdAt"}, {ProtoField: "priority"}, {ProtoField: "sourceModule"}, {ProtoField: "sourceAction"}},
				},
				{
					RpcMethod:      "UpdateNotification",
					Use:            "update-notification [index] [userAddress] [notificationType] [title] [message] [data] [isRead] [createdAt] [priority] [sourceModule] [sourceAction]",
					Short:          "Update notification",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "userAddress"}, {ProtoField: "notificationType"}, {ProtoField: "title"}, {ProtoField: "message"}, {ProtoField: "data"}, {ProtoField: "isRead"}, {ProtoField: "createdAt"}, {ProtoField: "priority"}, {ProtoField: "sourceModule"}, {ProtoField: "sourceAction"}},
				},
				{
					RpcMethod:      "DeleteNotification",
					Use:            "delete-notification [index]",
					Short:          "Delete notification",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateNotificationSettings",
					Use:            "create-notification-settings [index] [userAddress] [emailEnabled] [pushEnabled] [smsEnabled] [emailAddress] [phoneNumber] [notificationTypes] [frequency]",
					Short:          "Create a new notification-settings",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "userAddress"}, {ProtoField: "emailEnabled"}, {ProtoField: "pushEnabled"}, {ProtoField: "smsEnabled"}, {ProtoField: "emailAddress"}, {ProtoField: "phoneNumber"}, {ProtoField: "notificationTypes"}, {ProtoField: "frequency"}},
				},
				{
					RpcMethod:      "UpdateNotificationSettings",
					Use:            "update-notification-settings [index] [userAddress] [emailEnabled] [pushEnabled] [smsEnabled] [emailAddress] [phoneNumber] [notificationTypes] [frequency]",
					Short:          "Update notification-settings",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "userAddress"}, {ProtoField: "emailEnabled"}, {ProtoField: "pushEnabled"}, {ProtoField: "smsEnabled"}, {ProtoField: "emailAddress"}, {ProtoField: "phoneNumber"}, {ProtoField: "notificationTypes"}, {ProtoField: "frequency"}},
				},
				{
					RpcMethod:      "DeleteNotificationSettings",
					Use:            "delete-notification-settings [index]",
					Short:          "Delete notification-settings",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
