package marketplace

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "skillchain/api/skillchain/marketplace"
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
					RpcMethod: "JobPostingAll",
					Use:       "list-job-posting",
					Short:     "List all jobPosting",
				},
				{
					RpcMethod:      "JobPosting",
					Use:            "show-job-posting [id]",
					Short:          "Shows a jobPosting",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ProposalAll",
					Use:       "list-proposal",
					Short:     "List all proposal",
				},
				{
					RpcMethod:      "Proposal",
					Use:            "show-proposal [id]",
					Short:          "Shows a proposal",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ProjectAll",
					Use:       "list-project",
					Short:     "List all project",
				},
				{
					RpcMethod:      "Project",
					Use:            "show-project [id]",
					Short:          "Shows a project",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "MilestoneAll",
					Use:       "list-milestone",
					Short:     "List all milestone",
				},
				{
					RpcMethod:      "Milestone",
					Use:            "show-milestone [id]",
					Short:          "Shows a milestone",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "ListJobPostings",
					Use:            "list-job-postings",
					Short:          "Query list-job-postings",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "ListProposals",
					Use:            "list-proposals",
					Short:          "Query list-proposals",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "ListProjects",
					Use:            "list-projects",
					Short:          "Query list-projects",
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
					RpcMethod:      "CreateJobPosting",
					Use:            "create-job-posting [index] [clientAddress] [title] [description] [skillsRequired] [budgetAmount] [paymentCurrency] [deadline] [isActive] [createdAt]",
					Short:          "Create a new jobPosting",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "clientAddress"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "skillsRequired"}, {ProtoField: "budgetAmount"}, {ProtoField: "paymentCurrency"}, {ProtoField: "deadline"}, {ProtoField: "isActive"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "UpdateJobPosting",
					Use:            "update-job-posting [index] [clientAddress] [title] [description] [skillsRequired] [budgetAmount] [paymentCurrency] [deadline] [isActive] [createdAt]",
					Short:          "Update jobPosting",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "clientAddress"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "skillsRequired"}, {ProtoField: "budgetAmount"}, {ProtoField: "paymentCurrency"}, {ProtoField: "deadline"}, {ProtoField: "isActive"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "DeleteJobPosting",
					Use:            "delete-job-posting [index]",
					Short:          "Delete jobPosting",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateProposal",
					Use:            "create-proposal [index] [jobPostingId] [freelancerAddress] [proposedAmount] [proposedCurrency] [proposedTimeline] [coverLetter] [freelancerReputation] [isAccepted] [createdAt]",
					Short:          "Create a new proposal",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "jobPostingId"}, {ProtoField: "freelancerAddress"}, {ProtoField: "proposedAmount"}, {ProtoField: "proposedCurrency"}, {ProtoField: "proposedTimeline"}, {ProtoField: "coverLetter"}, {ProtoField: "freelancerReputation"}, {ProtoField: "isAccepted"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "UpdateProposal",
					Use:            "update-proposal [index] [jobPostingId] [freelancerAddress] [proposedAmount] [proposedCurrency] [proposedTimeline] [coverLetter] [freelancerReputation] [isAccepted] [createdAt]",
					Short:          "Update proposal",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "jobPostingId"}, {ProtoField: "freelancerAddress"}, {ProtoField: "proposedAmount"}, {ProtoField: "proposedCurrency"}, {ProtoField: "proposedTimeline"}, {ProtoField: "coverLetter"}, {ProtoField: "freelancerReputation"}, {ProtoField: "isAccepted"}, {ProtoField: "createdAt"}},
				},
				{
					RpcMethod:      "DeleteProposal",
					Use:            "delete-proposal [index]",
					Short:          "Delete proposal",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateProject",
					Use:            "create-project [index] [jobPostingId] [proposalId] [clientAddress] [freelancerAddress] [totalAmount] [paidAmount] [escrowAmount] [status] [startDate] [expectedEndDate] [actualEndDate]",
					Short:          "Create a new project",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "jobPostingId"}, {ProtoField: "proposalId"}, {ProtoField: "clientAddress"}, {ProtoField: "freelancerAddress"}, {ProtoField: "totalAmount"}, {ProtoField: "paidAmount"}, {ProtoField: "escrowAmount"}, {ProtoField: "status"}, {ProtoField: "startDate"}, {ProtoField: "expectedEndDate"}, {ProtoField: "actualEndDate"}},
				},
				{
					RpcMethod:      "UpdateProject",
					Use:            "update-project [index] [jobPostingId] [proposalId] [clientAddress] [freelancerAddress] [totalAmount] [paidAmount] [escrowAmount] [status] [startDate] [expectedEndDate] [actualEndDate]",
					Short:          "Update project",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "jobPostingId"}, {ProtoField: "proposalId"}, {ProtoField: "clientAddress"}, {ProtoField: "freelancerAddress"}, {ProtoField: "totalAmount"}, {ProtoField: "paidAmount"}, {ProtoField: "escrowAmount"}, {ProtoField: "status"}, {ProtoField: "startDate"}, {ProtoField: "expectedEndDate"}, {ProtoField: "actualEndDate"}},
				},
				{
					RpcMethod:      "DeleteProject",
					Use:            "delete-project [index]",
					Short:          "Delete project",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateMilestone",
					Use:            "create-milestone [index] [projectId] [title] [description] [amount] [dueDate] [status] [isCompleted] [isPaid] [submittedAt] [approvedAt]",
					Short:          "Create a new milestone",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "projectId"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "amount"}, {ProtoField: "dueDate"}, {ProtoField: "status"}, {ProtoField: "isCompleted"}, {ProtoField: "isPaid"}, {ProtoField: "submittedAt"}, {ProtoField: "approvedAt"}},
				},
				{
					RpcMethod:      "UpdateMilestone",
					Use:            "update-milestone [index] [projectId] [title] [description] [amount] [dueDate] [status] [isCompleted] [isPaid] [submittedAt] [approvedAt]",
					Short:          "Update milestone",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "projectId"}, {ProtoField: "title"}, {ProtoField: "description"}, {ProtoField: "amount"}, {ProtoField: "dueDate"}, {ProtoField: "status"}, {ProtoField: "isCompleted"}, {ProtoField: "isPaid"}, {ProtoField: "submittedAt"}, {ProtoField: "approvedAt"}},
				},
				{
					RpcMethod:      "DeleteMilestone",
					Use:            "delete-milestone [index]",
					Short:          "Delete milestone",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "AcceptProposal",
					Use:            "accept-proposal [proposal-id]",
					Short:          "Send a acceptProposal tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proposalId"}},
				},
				{
					RpcMethod:      "CompleteMilestone",
					Use:            "complete-milestone [milestone-id] [delivery-notes]",
					Short:          "Send a completeMilestone tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "milestoneId"}, {ProtoField: "deliveryNotes"}},
				},
				{
					RpcMethod:      "ReleasePayment",
					Use:            "release-payment [milestone-id] [rating] [feedback]",
					Short:          "Send a releasePayment tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "milestoneId"}, {ProtoField: "rating"}, {ProtoField: "feedback"}},
				},
				{
					RpcMethod:      "DisputeProject",
					Use:            "dispute-project [project-id] [reason] [evidence]",
					Short:          "Send a disputeProject tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectId"}, {ProtoField: "reason"}, {ProtoField: "evidence"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
