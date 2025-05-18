package job

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/SkillChainLab/skillchain/testutil/sample"
	jobsimulation "github.com/SkillChainLab/skillchain/x/job/simulation"
	"github.com/SkillChainLab/skillchain/x/job/types"
)

// avoid unused import issue
var (
	_ = jobsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateJob = "op_weight_msg_create_job"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateJob int = 100

	opWeightMsgApplyJob = "op_weight_msg_apply_job"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApplyJob int = 100

	opWeightMsgReviewApplication = "op_weight_msg_review_application"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReviewApplication int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	jobGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&jobGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateJob int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateJob, &weightMsgCreateJob, nil,
		func(_ *rand.Rand) {
			weightMsgCreateJob = defaultWeightMsgCreateJob
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateJob,
		jobsimulation.SimulateMsgCreateJob(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApplyJob int
	simState.AppParams.GetOrGenerate(opWeightMsgApplyJob, &weightMsgApplyJob, nil,
		func(_ *rand.Rand) {
			weightMsgApplyJob = defaultWeightMsgApplyJob
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApplyJob,
		jobsimulation.SimulateMsgApplyJob(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReviewApplication int
	simState.AppParams.GetOrGenerate(opWeightMsgReviewApplication, &weightMsgReviewApplication, nil,
		func(_ *rand.Rand) {
			weightMsgReviewApplication = defaultWeightMsgReviewApplication
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReviewApplication,
		jobsimulation.SimulateMsgReviewApplication(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateJob,
			defaultWeightMsgCreateJob,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				jobsimulation.SimulateMsgCreateJob(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgApplyJob,
			defaultWeightMsgApplyJob,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				jobsimulation.SimulateMsgApplyJob(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgReviewApplication,
			defaultWeightMsgReviewApplication,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				jobsimulation.SimulateMsgReviewApplication(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
