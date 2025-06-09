package filestorage

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"skillchain/testutil/sample"
	filestoragesimulation "skillchain/x/filestorage/simulation"
	"skillchain/x/filestorage/types"
)

// avoid unused import issue
var (
	_ = filestoragesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateFileRecord = "op_weight_msg_file_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateFileRecord int = 100

	opWeightMsgUpdateFileRecord = "op_weight_msg_file_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateFileRecord int = 100

	opWeightMsgDeleteFileRecord = "op_weight_msg_file_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteFileRecord int = 100

	opWeightMsgCreateFilePermission = "op_weight_msg_file_permission"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateFilePermission int = 100

	opWeightMsgUpdateFilePermission = "op_weight_msg_file_permission"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateFilePermission int = 100

	opWeightMsgDeleteFilePermission = "op_weight_msg_file_permission"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteFilePermission int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	filestorageGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		FileRecordList: []types.FileRecord{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		FilePermissionList: []types.FilePermission{
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
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&filestorageGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateFileRecord int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateFileRecord, &weightMsgCreateFileRecord, nil,
		func(_ *rand.Rand) {
			weightMsgCreateFileRecord = defaultWeightMsgCreateFileRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateFileRecord,
		filestoragesimulation.SimulateMsgCreateFileRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateFileRecord int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateFileRecord, &weightMsgUpdateFileRecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateFileRecord = defaultWeightMsgUpdateFileRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateFileRecord,
		filestoragesimulation.SimulateMsgUpdateFileRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteFileRecord int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteFileRecord, &weightMsgDeleteFileRecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteFileRecord = defaultWeightMsgDeleteFileRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteFileRecord,
		filestoragesimulation.SimulateMsgDeleteFileRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateFilePermission int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateFilePermission, &weightMsgCreateFilePermission, nil,
		func(_ *rand.Rand) {
			weightMsgCreateFilePermission = defaultWeightMsgCreateFilePermission
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateFilePermission,
		filestoragesimulation.SimulateMsgCreateFilePermission(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateFilePermission int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateFilePermission, &weightMsgUpdateFilePermission, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateFilePermission = defaultWeightMsgUpdateFilePermission
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateFilePermission,
		filestoragesimulation.SimulateMsgUpdateFilePermission(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteFilePermission int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteFilePermission, &weightMsgDeleteFilePermission, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteFilePermission = defaultWeightMsgDeleteFilePermission
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteFilePermission,
		filestoragesimulation.SimulateMsgDeleteFilePermission(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateFileRecord,
			defaultWeightMsgCreateFileRecord,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filestoragesimulation.SimulateMsgCreateFileRecord(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateFileRecord,
			defaultWeightMsgUpdateFileRecord,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filestoragesimulation.SimulateMsgUpdateFileRecord(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteFileRecord,
			defaultWeightMsgDeleteFileRecord,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filestoragesimulation.SimulateMsgDeleteFileRecord(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateFilePermission,
			defaultWeightMsgCreateFilePermission,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filestoragesimulation.SimulateMsgCreateFilePermission(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateFilePermission,
			defaultWeightMsgUpdateFilePermission,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filestoragesimulation.SimulateMsgUpdateFilePermission(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteFilePermission,
			defaultWeightMsgDeleteFilePermission,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				filestoragesimulation.SimulateMsgDeleteFilePermission(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
