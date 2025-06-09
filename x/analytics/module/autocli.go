package analytics

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "skillchain/api/skillchain/analytics"
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
					RpcMethod: "PlatformMetricAll",
					Use:       "list-platform-metric",
					Short:     "List all platform-metric",
				},
				{
					RpcMethod:      "PlatformMetric",
					Use:            "show-platform-metric [id]",
					Short:          "Shows a platform-metric",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "UserActivityAll",
					Use:       "list-user-activity",
					Short:     "List all user-activity",
				},
				{
					RpcMethod:      "UserActivity",
					Use:            "show-user-activity [id]",
					Short:          "Shows a user-activity",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "RevenueRecordAll",
					Use:       "list-revenue-record",
					Short:     "List all revenue-record",
				},
				{
					RpcMethod:      "RevenueRecord",
					Use:            "show-revenue-record [id]",
					Short:          "Shows a revenue-record",
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
					RpcMethod:      "CreatePlatformMetric",
					Use:            "create-platform-metric [index] [metricName] [metricValue] [metricType] [period] [timestamp] [metadata]",
					Short:          "Create a new platform-metric",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "metricName"}, {ProtoField: "metricValue"}, {ProtoField: "metricType"}, {ProtoField: "period"}, {ProtoField: "timestamp"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "UpdatePlatformMetric",
					Use:            "update-platform-metric [index] [metricName] [metricValue] [metricType] [period] [timestamp] [metadata]",
					Short:          "Update platform-metric",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "metricName"}, {ProtoField: "metricValue"}, {ProtoField: "metricType"}, {ProtoField: "period"}, {ProtoField: "timestamp"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "DeletePlatformMetric",
					Use:            "delete-platform-metric [index]",
					Short:          "Delete platform-metric",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateUserActivity",
					Use:            "create-user-activity [index] [userAddress] [activityType] [action] [resourceId] [timestamp] [ipAddress] [userAgent] [metadata]",
					Short:          "Create a new user-activity",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "userAddress"}, {ProtoField: "activityType"}, {ProtoField: "action"}, {ProtoField: "resourceId"}, {ProtoField: "timestamp"}, {ProtoField: "ipAddress"}, {ProtoField: "userAgent"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "UpdateUserActivity",
					Use:            "update-user-activity [index] [userAddress] [activityType] [action] [resourceId] [timestamp] [ipAddress] [userAgent] [metadata]",
					Short:          "Update user-activity",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "userAddress"}, {ProtoField: "activityType"}, {ProtoField: "action"}, {ProtoField: "resourceId"}, {ProtoField: "timestamp"}, {ProtoField: "ipAddress"}, {ProtoField: "userAgent"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "DeleteUserActivity",
					Use:            "delete-user-activity [index]",
					Short:          "Delete user-activity",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateRevenueRecord",
					Use:            "create-revenue-record [index] [transactionType] [amount] [currency] [fromAddress] [toAddress] [timestamp] [feeAmount] [projectId] [platformFee]",
					Short:          "Create a new revenue-record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "transactionType"}, {ProtoField: "amount"}, {ProtoField: "currency"}, {ProtoField: "fromAddress"}, {ProtoField: "toAddress"}, {ProtoField: "timestamp"}, {ProtoField: "feeAmount"}, {ProtoField: "projectId"}, {ProtoField: "platformFee"}},
				},
				{
					RpcMethod:      "UpdateRevenueRecord",
					Use:            "update-revenue-record [index] [transactionType] [amount] [currency] [fromAddress] [toAddress] [timestamp] [feeAmount] [projectId] [platformFee]",
					Short:          "Update revenue-record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "transactionType"}, {ProtoField: "amount"}, {ProtoField: "currency"}, {ProtoField: "fromAddress"}, {ProtoField: "toAddress"}, {ProtoField: "timestamp"}, {ProtoField: "feeAmount"}, {ProtoField: "projectId"}, {ProtoField: "platformFee"}},
				},
				{
					RpcMethod:      "DeleteRevenueRecord",
					Use:            "delete-revenue-record [index]",
					Short:          "Delete revenue-record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
