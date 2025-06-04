package skilltoken

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/SkillChainLab/skillchain/api/skillchain/skilltoken"
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
					RpcMethod:      "Token",
					Use:            "token [symbol]",
					Short:          "Query token",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "symbol"}},
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
					RpcMethod:      "CreateToken",
					Use:            "create-token [symbol] [name] [decimals] [mintable] [burnable] [transferable]",
					Short:          "Send a create-token tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "symbol"}, {ProtoField: "name"}, {ProtoField: "decimals"}, {ProtoField: "mintable"}, {ProtoField: "burnable"}, {ProtoField: "transferable"}},
				},
				{
					RpcMethod:      "TransferToken",
					Use:            "transfer-token [to] [amount]",
					Short:          "Send a transfer-token tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "to"}, {ProtoField: "amount"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
