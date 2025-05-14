package profile

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/SkillChainLab/skillchain/api/skillchain/profile"
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
					RpcMethod:      "ShowProfile",
					Use:            "show-profile [username]",
					Short:          "Query show-profile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "username"}},
				},

				{
					RpcMethod:      "ListProfile",
					Use:            "list-profile",
					Short:          "Query list-profile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
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
					RpcMethod:      "CreateProfile",
					Use:            "create-profile [username] [bio]",
					Short:          "Send a create-profile tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "username"}, {ProtoField: "bio"}},
				},
				{
					RpcMethod:      "DeleteProfile",
					Use:            "delete-profile [username]",
					Short:          "Send a delete-profile tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "username"}},
				},
				{
					RpcMethod:      "UpdateProfile",
					Use:            "update-profile [username] [bio]",
					Short:          "Send a update-profile tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "username"}, {ProtoField: "bio"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
