package profile

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "skillchain/api/skillchain/profile"
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
					RpcMethod: "UserProfileAll",
					Use:       "list-user-profile",
					Short:     "List all user-profile",
				},
				{
					RpcMethod:      "UserProfile",
					Use:            "show-user-profile [id]",
					Short:          "Shows a user-profile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "UserSkillAll",
					Use:       "list-user-skill",
					Short:     "List all user-skill",
				},
				{
					RpcMethod:      "UserSkill",
					Use:            "show-user-skill [id]",
					Short:          "Shows a user-skill",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "SkillEndorsementAll",
					Use:       "list-skill-endorsement",
					Short:     "List all skill-endorsement",
				},
				{
					RpcMethod:      "SkillEndorsement",
					Use:            "show-skill-endorsement [id]",
					Short:          "Shows a skill-endorsement",
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
					RpcMethod:      "CreateProfile",
					Use:            "create-profile [display-name] [bio] [location] [website] [github] [linkedin] [twitter] [avatar]",
					Short:          "Send a create-profile tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "displayName"}, {ProtoField: "bio"}, {ProtoField: "location"}, {ProtoField: "website"}, {ProtoField: "github"}, {ProtoField: "linkedin"}, {ProtoField: "twitter"}, {ProtoField: "avatar"}},
				},
				{
					RpcMethod:      "CreateUserProfile",
					Use:            "create-user-profile [index] [owner] [displayName] [bio] [location] [website] [github] [linkedin] [twitter] [avatar] [reputationScore] [createdAt] [updatedAt]",
					Short:          "Create a new user-profile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "owner"}, {ProtoField: "displayName"}, {ProtoField: "bio"}, {ProtoField: "location"}, {ProtoField: "website"}, {ProtoField: "github"}, {ProtoField: "linkedin"}, {ProtoField: "twitter"}, {ProtoField: "avatar"}, {ProtoField: "reputationScore"}, {ProtoField: "createdAt"}, {ProtoField: "updatedAt"}},
				},
				{
					RpcMethod:      "UpdateUserProfile",
					Use:            "update-user-profile [index] [owner] [displayName] [bio] [location] [website] [github] [linkedin] [twitter] [avatar] [reputationScore] [createdAt] [updatedAt]",
					Short:          "Update user-profile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "owner"}, {ProtoField: "displayName"}, {ProtoField: "bio"}, {ProtoField: "location"}, {ProtoField: "website"}, {ProtoField: "github"}, {ProtoField: "linkedin"}, {ProtoField: "twitter"}, {ProtoField: "avatar"}, {ProtoField: "reputationScore"}, {ProtoField: "createdAt"}, {ProtoField: "updatedAt"}},
				},
				{
					RpcMethod:      "DeleteUserProfile",
					Use:            "delete-user-profile [index]",
					Short:          "Delete user-profile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateUserSkill",
					Use:            "create-user-skill [index] [owner] [skillName] [proficiencyLevel] [yearsExperience] [verified] [verifiedBy] [verificationDate] [endorsementCount]",
					Short:          "Create a new user-skill",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "owner"}, {ProtoField: "skillName"}, {ProtoField: "proficiencyLevel"}, {ProtoField: "yearsExperience"}, {ProtoField: "verified"}, {ProtoField: "verifiedBy"}, {ProtoField: "verificationDate"}, {ProtoField: "endorsementCount"}},
				},
				{
					RpcMethod:      "UpdateUserSkill",
					Use:            "update-user-skill [index] [owner] [skillName] [proficiencyLevel] [yearsExperience] [verified] [verifiedBy] [verificationDate] [endorsementCount]",
					Short:          "Update user-skill",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "owner"}, {ProtoField: "skillName"}, {ProtoField: "proficiencyLevel"}, {ProtoField: "yearsExperience"}, {ProtoField: "verified"}, {ProtoField: "verifiedBy"}, {ProtoField: "verificationDate"}, {ProtoField: "endorsementCount"}},
				},
				{
					RpcMethod:      "DeleteUserSkill",
					Use:            "delete-user-skill [index]",
					Short:          "Delete user-skill",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "EndorseSkill",
					Use:            "endorse-skill [target-user] [skill-name] [endorsement-type] [comment]",
					Short:          "Send a endorse-skill tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "targetUser"}, {ProtoField: "skillName"}, {ProtoField: "endorsementType"}, {ProtoField: "comment"}},
				},
				{
					RpcMethod:      "CreateSkillEndorsement",
					Use:            "create-skill-endorsement [index] [endorser] [targetUser] [skillName] [endorsementType] [comment] [createdAt] [skillTokensStaked]",
					Short:          "Create a new skill-endorsement",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "endorser"}, {ProtoField: "targetUser"}, {ProtoField: "skillName"}, {ProtoField: "endorsementType"}, {ProtoField: "comment"}, {ProtoField: "createdAt"}, {ProtoField: "skillTokensStaked"}},
				},
				{
					RpcMethod:      "UpdateSkillEndorsement",
					Use:            "update-skill-endorsement [index] [endorser] [targetUser] [skillName] [endorsementType] [comment] [createdAt] [skillTokensStaked]",
					Short:          "Update skill-endorsement",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "endorser"}, {ProtoField: "targetUser"}, {ProtoField: "skillName"}, {ProtoField: "endorsementType"}, {ProtoField: "comment"}, {ProtoField: "createdAt"}, {ProtoField: "skillTokensStaked"}},
				},
				{
					RpcMethod:      "DeleteSkillEndorsement",
					Use:            "delete-skill-endorsement [index]",
					Short:          "Delete skill-endorsement",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "DisputeEndorsement",
					Use:       "dispute-endorsement",
					Short:     "Dispute an endorsement for failed delivery",
				},
				{
					RpcMethod:      "WithdrawStakedTokens",
					Use:            "withdraw-staked-tokens [skill-name]",
					Short:          "Withdraw your staked tokens for a specific skill",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "skillName"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
