package job

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/SkillChainLab/skillchain/api/skillchain/job"
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
					RpcMethod:      "ListJob",
					Use:            "list-job",
					Short:          "Query list-job",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "ShowJob",
					Use:            "show-job [id]",
					Short:          "Query show-job",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListApplication",
					Use:            "list-application",
					Short:          "Query list-application",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "ListMyApplications",
					Use:            "list-my-applications [applicant]",
					Short:          "Query list-my-applications",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "applicant"}},
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
					RpcMethod:      "CreateJob",
					Use:            "create-job [title] [description] [budget]",
					Short:          "Send a create-job tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "budget"}},
				},
				{
					RpcMethod:      "ApplyJob",
					Use:            "apply-job [job-id] [cover-letter]",
					Short:          "Send a apply-job tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "jobId"}, {ProtoField: "coverLetter"}},
				},
				{
					RpcMethod:      "ReviewApplication",
					Use:            "review-application [job-id] [applicant] [status]",
					Short:          "Send a review-application tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "jobId"}, {ProtoField: "applicant"}, {ProtoField: "status"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
