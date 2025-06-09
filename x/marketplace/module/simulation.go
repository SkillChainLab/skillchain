package marketplace

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"skillchain/testutil/sample"
	marketplacesimulation "skillchain/x/marketplace/simulation"
	"skillchain/x/marketplace/types"
)

// avoid unused import issue
var (
	_ = marketplacesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateJobPosting = "op_weight_msg_job_posting"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateJobPosting int = 100

	opWeightMsgUpdateJobPosting = "op_weight_msg_job_posting"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateJobPosting int = 100

	opWeightMsgDeleteJobPosting = "op_weight_msg_job_posting"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteJobPosting int = 100

	opWeightMsgCreateProposal = "op_weight_msg_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProposal int = 100

	opWeightMsgUpdateProposal = "op_weight_msg_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProposal int = 100

	opWeightMsgDeleteProposal = "op_weight_msg_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteProposal int = 100

	opWeightMsgCreateProject = "op_weight_msg_project"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProject int = 100

	opWeightMsgUpdateProject = "op_weight_msg_project"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProject int = 100

	opWeightMsgDeleteProject = "op_weight_msg_project"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteProject int = 100

	opWeightMsgCreateMilestone = "op_weight_msg_milestone"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMilestone int = 100

	opWeightMsgUpdateMilestone = "op_weight_msg_milestone"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMilestone int = 100

	opWeightMsgDeleteMilestone = "op_weight_msg_milestone"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMilestone int = 100

	opWeightMsgAcceptProposal = "op_weight_msg_accept_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAcceptProposal int = 100

	opWeightMsgCompleteMilestone = "op_weight_msg_complete_milestone"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCompleteMilestone int = 100

	opWeightMsgReleasePayment = "op_weight_msg_release_payment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReleasePayment int = 100

	opWeightMsgDisputeProject = "op_weight_msg_dispute_project"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDisputeProject int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	marketplaceGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		JobPostingList: []types.JobPosting{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		ProposalList: []types.Proposal{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		ProjectList: []types.Project{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		MilestoneList: []types.Milestone{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&marketplaceGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateJobPosting int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateJobPosting, &weightMsgCreateJobPosting, nil,
		func(_ *rand.Rand) {
			weightMsgCreateJobPosting = defaultWeightMsgCreateJobPosting
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateJobPosting,
		marketplacesimulation.SimulateMsgCreateJobPosting(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateJobPosting int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateJobPosting, &weightMsgUpdateJobPosting, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateJobPosting = defaultWeightMsgUpdateJobPosting
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateJobPosting,
		marketplacesimulation.SimulateMsgUpdateJobPosting(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteJobPosting int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteJobPosting, &weightMsgDeleteJobPosting, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteJobPosting = defaultWeightMsgDeleteJobPosting
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteJobPosting,
		marketplacesimulation.SimulateMsgDeleteJobPosting(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateProposal, &weightMsgCreateProposal, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProposal = defaultWeightMsgCreateProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProposal,
		marketplacesimulation.SimulateMsgCreateProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateProposal, &weightMsgUpdateProposal, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProposal = defaultWeightMsgUpdateProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProposal,
		marketplacesimulation.SimulateMsgUpdateProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteProposal, &weightMsgDeleteProposal, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProposal = defaultWeightMsgDeleteProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProposal,
		marketplacesimulation.SimulateMsgDeleteProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateProject int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateProject, &weightMsgCreateProject, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProject = defaultWeightMsgCreateProject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProject,
		marketplacesimulation.SimulateMsgCreateProject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProject int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateProject, &weightMsgUpdateProject, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProject = defaultWeightMsgUpdateProject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProject,
		marketplacesimulation.SimulateMsgUpdateProject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteProject int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteProject, &weightMsgDeleteProject, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProject = defaultWeightMsgDeleteProject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProject,
		marketplacesimulation.SimulateMsgDeleteProject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateMilestone int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateMilestone, &weightMsgCreateMilestone, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMilestone = defaultWeightMsgCreateMilestone
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMilestone,
		marketplacesimulation.SimulateMsgCreateMilestone(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMilestone int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateMilestone, &weightMsgUpdateMilestone, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMilestone = defaultWeightMsgUpdateMilestone
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMilestone,
		marketplacesimulation.SimulateMsgUpdateMilestone(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMilestone int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteMilestone, &weightMsgDeleteMilestone, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMilestone = defaultWeightMsgDeleteMilestone
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMilestone,
		marketplacesimulation.SimulateMsgDeleteMilestone(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAcceptProposal int
	simState.AppParams.GetOrGenerate(opWeightMsgAcceptProposal, &weightMsgAcceptProposal, nil,
		func(_ *rand.Rand) {
			weightMsgAcceptProposal = defaultWeightMsgAcceptProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAcceptProposal,
		marketplacesimulation.SimulateMsgAcceptProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCompleteMilestone int
	simState.AppParams.GetOrGenerate(opWeightMsgCompleteMilestone, &weightMsgCompleteMilestone, nil,
		func(_ *rand.Rand) {
			weightMsgCompleteMilestone = defaultWeightMsgCompleteMilestone
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCompleteMilestone,
		marketplacesimulation.SimulateMsgCompleteMilestone(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReleasePayment int
	simState.AppParams.GetOrGenerate(opWeightMsgReleasePayment, &weightMsgReleasePayment, nil,
		func(_ *rand.Rand) {
			weightMsgReleasePayment = defaultWeightMsgReleasePayment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReleasePayment,
		marketplacesimulation.SimulateMsgReleasePayment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDisputeProject int
	simState.AppParams.GetOrGenerate(opWeightMsgDisputeProject, &weightMsgDisputeProject, nil,
		func(_ *rand.Rand) {
			weightMsgDisputeProject = defaultWeightMsgDisputeProject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDisputeProject,
		marketplacesimulation.SimulateMsgDisputeProject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateJobPosting,
			defaultWeightMsgCreateJobPosting,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgCreateJobPosting(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateJobPosting,
			defaultWeightMsgUpdateJobPosting,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgUpdateJobPosting(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteJobPosting,
			defaultWeightMsgDeleteJobPosting,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgDeleteJobPosting(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateProposal,
			defaultWeightMsgCreateProposal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgCreateProposal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateProposal,
			defaultWeightMsgUpdateProposal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgUpdateProposal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteProposal,
			defaultWeightMsgDeleteProposal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgDeleteProposal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateProject,
			defaultWeightMsgCreateProject,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgCreateProject(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateProject,
			defaultWeightMsgUpdateProject,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgUpdateProject(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteProject,
			defaultWeightMsgDeleteProject,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgDeleteProject(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateMilestone,
			defaultWeightMsgCreateMilestone,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgCreateMilestone(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateMilestone,
			defaultWeightMsgUpdateMilestone,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgUpdateMilestone(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteMilestone,
			defaultWeightMsgDeleteMilestone,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgDeleteMilestone(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAcceptProposal,
			defaultWeightMsgAcceptProposal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgAcceptProposal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCompleteMilestone,
			defaultWeightMsgCompleteMilestone,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgCompleteMilestone(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgReleasePayment,
			defaultWeightMsgReleasePayment,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgReleasePayment(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDisputeProject,
			defaultWeightMsgDisputeProject,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				marketplacesimulation.SimulateMsgDisputeProject(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
