package profile

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/SkillChainLab/skillchain/testutil/sample"
	profilesimulation "github.com/SkillChainLab/skillchain/x/profile/simulation"
	"github.com/SkillChainLab/skillchain/x/profile/types"
)

// avoid unused import issue
var (
	_ = profilesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateProfile = "op_weight_msg_create_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProfile int = 100

	opWeightMsgDeleteProfile = "op_weight_msg_delete_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteProfile int = 100

	opWeightMsgUpdateProfile = "op_weight_msg_update_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProfile int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	profileGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&profileGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateProfile, &weightMsgCreateProfile, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProfile = defaultWeightMsgCreateProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProfile,
		profilesimulation.SimulateMsgCreateProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteProfile, &weightMsgDeleteProfile, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProfile = defaultWeightMsgDeleteProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProfile,
		profilesimulation.SimulateMsgDeleteProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateProfile, &weightMsgUpdateProfile, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProfile = defaultWeightMsgUpdateProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProfile,
		profilesimulation.SimulateMsgUpdateProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateProfile,
			defaultWeightMsgCreateProfile,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgCreateProfile(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteProfile,
			defaultWeightMsgDeleteProfile,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgDeleteProfile(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateProfile,
			defaultWeightMsgUpdateProfile,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgUpdateProfile(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
