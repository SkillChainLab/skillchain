package analytics

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"skillchain/testutil/sample"
	analyticssimulation "skillchain/x/analytics/simulation"
	"skillchain/x/analytics/types"
)

// avoid unused import issue
var (
	_ = analyticssimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreatePlatformMetric = "op_weight_msg_platform_metric"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePlatformMetric int = 100

	opWeightMsgUpdatePlatformMetric = "op_weight_msg_platform_metric"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePlatformMetric int = 100

	opWeightMsgDeletePlatformMetric = "op_weight_msg_platform_metric"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePlatformMetric int = 100

	opWeightMsgCreateUserActivity = "op_weight_msg_user_activity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateUserActivity int = 100

	opWeightMsgUpdateUserActivity = "op_weight_msg_user_activity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateUserActivity int = 100

	opWeightMsgDeleteUserActivity = "op_weight_msg_user_activity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteUserActivity int = 100

	opWeightMsgCreateRevenueRecord = "op_weight_msg_revenue_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRevenueRecord int = 100

	opWeightMsgUpdateRevenueRecord = "op_weight_msg_revenue_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRevenueRecord int = 100

	opWeightMsgDeleteRevenueRecord = "op_weight_msg_revenue_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteRevenueRecord int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	analyticsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PlatformMetricList: []types.PlatformMetric{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		UserActivityList: []types.UserActivity{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		RevenueRecordList: []types.RevenueRecord{
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
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&analyticsGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePlatformMetric int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePlatformMetric, &weightMsgCreatePlatformMetric, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePlatformMetric = defaultWeightMsgCreatePlatformMetric
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePlatformMetric,
		analyticssimulation.SimulateMsgCreatePlatformMetric(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePlatformMetric int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePlatformMetric, &weightMsgUpdatePlatformMetric, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePlatformMetric = defaultWeightMsgUpdatePlatformMetric
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePlatformMetric,
		analyticssimulation.SimulateMsgUpdatePlatformMetric(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePlatformMetric int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePlatformMetric, &weightMsgDeletePlatformMetric, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePlatformMetric = defaultWeightMsgDeletePlatformMetric
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePlatformMetric,
		analyticssimulation.SimulateMsgDeletePlatformMetric(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateUserActivity int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateUserActivity, &weightMsgCreateUserActivity, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUserActivity = defaultWeightMsgCreateUserActivity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUserActivity,
		analyticssimulation.SimulateMsgCreateUserActivity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateUserActivity int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateUserActivity, &weightMsgUpdateUserActivity, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateUserActivity = defaultWeightMsgUpdateUserActivity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateUserActivity,
		analyticssimulation.SimulateMsgUpdateUserActivity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteUserActivity int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteUserActivity, &weightMsgDeleteUserActivity, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUserActivity = defaultWeightMsgDeleteUserActivity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteUserActivity,
		analyticssimulation.SimulateMsgDeleteUserActivity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateRevenueRecord int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateRevenueRecord, &weightMsgCreateRevenueRecord, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRevenueRecord = defaultWeightMsgCreateRevenueRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRevenueRecord,
		analyticssimulation.SimulateMsgCreateRevenueRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRevenueRecord int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateRevenueRecord, &weightMsgUpdateRevenueRecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRevenueRecord = defaultWeightMsgUpdateRevenueRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRevenueRecord,
		analyticssimulation.SimulateMsgUpdateRevenueRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteRevenueRecord int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteRevenueRecord, &weightMsgDeleteRevenueRecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRevenueRecord = defaultWeightMsgDeleteRevenueRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteRevenueRecord,
		analyticssimulation.SimulateMsgDeleteRevenueRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePlatformMetric,
			defaultWeightMsgCreatePlatformMetric,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				analyticssimulation.SimulateMsgCreatePlatformMetric(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdatePlatformMetric,
			defaultWeightMsgUpdatePlatformMetric,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				analyticssimulation.SimulateMsgUpdatePlatformMetric(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeletePlatformMetric,
			defaultWeightMsgDeletePlatformMetric,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				analyticssimulation.SimulateMsgDeletePlatformMetric(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateUserActivity,
			defaultWeightMsgCreateUserActivity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				analyticssimulation.SimulateMsgCreateUserActivity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateUserActivity,
			defaultWeightMsgUpdateUserActivity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				analyticssimulation.SimulateMsgUpdateUserActivity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteUserActivity,
			defaultWeightMsgDeleteUserActivity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				analyticssimulation.SimulateMsgDeleteUserActivity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateRevenueRecord,
			defaultWeightMsgCreateRevenueRecord,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				analyticssimulation.SimulateMsgCreateRevenueRecord(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateRevenueRecord,
			defaultWeightMsgUpdateRevenueRecord,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				analyticssimulation.SimulateMsgUpdateRevenueRecord(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteRevenueRecord,
			defaultWeightMsgDeleteRevenueRecord,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				analyticssimulation.SimulateMsgDeleteRevenueRecord(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
