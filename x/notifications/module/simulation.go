package notifications

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"skillchain/testutil/sample"
	notificationssimulation "skillchain/x/notifications/simulation"
	"skillchain/x/notifications/types"
)

// avoid unused import issue
var (
	_ = notificationssimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateNotification = "op_weight_msg_notification"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateNotification int = 100

	opWeightMsgUpdateNotification = "op_weight_msg_notification"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateNotification int = 100

	opWeightMsgDeleteNotification = "op_weight_msg_notification"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteNotification int = 100

	opWeightMsgCreateNotificationSettings = "op_weight_msg_notification_settings"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateNotificationSettings int = 100

	opWeightMsgUpdateNotificationSettings = "op_weight_msg_notification_settings"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateNotificationSettings int = 100

	opWeightMsgDeleteNotificationSettings = "op_weight_msg_notification_settings"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteNotificationSettings int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	notificationsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		NotificationList: []types.Notification{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		NotificationSettingsList: []types.NotificationSettings{
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
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&notificationsGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateNotification int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateNotification, &weightMsgCreateNotification, nil,
		func(_ *rand.Rand) {
			weightMsgCreateNotification = defaultWeightMsgCreateNotification
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateNotification,
		notificationssimulation.SimulateMsgCreateNotification(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateNotification int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateNotification, &weightMsgUpdateNotification, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateNotification = defaultWeightMsgUpdateNotification
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateNotification,
		notificationssimulation.SimulateMsgUpdateNotification(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteNotification int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteNotification, &weightMsgDeleteNotification, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteNotification = defaultWeightMsgDeleteNotification
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteNotification,
		notificationssimulation.SimulateMsgDeleteNotification(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateNotificationSettings int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateNotificationSettings, &weightMsgCreateNotificationSettings, nil,
		func(_ *rand.Rand) {
			weightMsgCreateNotificationSettings = defaultWeightMsgCreateNotificationSettings
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateNotificationSettings,
		notificationssimulation.SimulateMsgCreateNotificationSettings(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateNotificationSettings int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateNotificationSettings, &weightMsgUpdateNotificationSettings, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateNotificationSettings = defaultWeightMsgUpdateNotificationSettings
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateNotificationSettings,
		notificationssimulation.SimulateMsgUpdateNotificationSettings(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteNotificationSettings int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteNotificationSettings, &weightMsgDeleteNotificationSettings, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteNotificationSettings = defaultWeightMsgDeleteNotificationSettings
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteNotificationSettings,
		notificationssimulation.SimulateMsgDeleteNotificationSettings(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateNotification,
			defaultWeightMsgCreateNotification,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				notificationssimulation.SimulateMsgCreateNotification(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateNotification,
			defaultWeightMsgUpdateNotification,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				notificationssimulation.SimulateMsgUpdateNotification(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteNotification,
			defaultWeightMsgDeleteNotification,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				notificationssimulation.SimulateMsgDeleteNotification(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateNotificationSettings,
			defaultWeightMsgCreateNotificationSettings,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				notificationssimulation.SimulateMsgCreateNotificationSettings(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateNotificationSettings,
			defaultWeightMsgUpdateNotificationSettings,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				notificationssimulation.SimulateMsgUpdateNotificationSettings(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteNotificationSettings,
			defaultWeightMsgDeleteNotificationSettings,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				notificationssimulation.SimulateMsgDeleteNotificationSettings(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
