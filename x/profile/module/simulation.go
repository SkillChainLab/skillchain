package profile

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"skillchain/testutil/sample"
	profilesimulation "skillchain/x/profile/simulation"
	"skillchain/x/profile/types"
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

	opWeightMsgCreateUserProfile = "op_weight_msg_user_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateUserProfile int = 100

	opWeightMsgUpdateUserProfile = "op_weight_msg_user_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateUserProfile int = 100

	opWeightMsgDeleteUserProfile = "op_weight_msg_user_profile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteUserProfile int = 100

	opWeightMsgCreateUserSkill = "op_weight_msg_user_skill"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateUserSkill int = 100

	opWeightMsgUpdateUserSkill = "op_weight_msg_user_skill"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateUserSkill int = 100

	opWeightMsgDeleteUserSkill = "op_weight_msg_user_skill"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteUserSkill int = 100

	opWeightMsgEndorseSkill = "op_weight_msg_endorse_skill"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEndorseSkill int = 100

	opWeightMsgCreateSkillEndorsement = "op_weight_msg_skill_endorsement"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSkillEndorsement int = 100

	opWeightMsgUpdateSkillEndorsement = "op_weight_msg_skill_endorsement"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSkillEndorsement int = 100

	opWeightMsgDeleteSkillEndorsement = "op_weight_msg_skill_endorsement"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSkillEndorsement int = 100

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
		UserProfileList: []types.UserProfile{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		UserSkillList: []types.UserSkill{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		SkillEndorsementList: []types.SkillEndorsement{
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

	var weightMsgCreateUserProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateUserProfile, &weightMsgCreateUserProfile, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUserProfile = defaultWeightMsgCreateUserProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUserProfile,
		profilesimulation.SimulateMsgCreateUserProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateUserProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateUserProfile, &weightMsgUpdateUserProfile, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateUserProfile = defaultWeightMsgUpdateUserProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateUserProfile,
		profilesimulation.SimulateMsgUpdateUserProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteUserProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteUserProfile, &weightMsgDeleteUserProfile, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUserProfile = defaultWeightMsgDeleteUserProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteUserProfile,
		profilesimulation.SimulateMsgDeleteUserProfile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateUserSkill int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateUserSkill, &weightMsgCreateUserSkill, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUserSkill = defaultWeightMsgCreateUserSkill
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUserSkill,
		profilesimulation.SimulateMsgCreateUserSkill(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateUserSkill int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateUserSkill, &weightMsgUpdateUserSkill, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateUserSkill = defaultWeightMsgUpdateUserSkill
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateUserSkill,
		profilesimulation.SimulateMsgUpdateUserSkill(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteUserSkill int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteUserSkill, &weightMsgDeleteUserSkill, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUserSkill = defaultWeightMsgDeleteUserSkill
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteUserSkill,
		profilesimulation.SimulateMsgDeleteUserSkill(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEndorseSkill int
	simState.AppParams.GetOrGenerate(opWeightMsgEndorseSkill, &weightMsgEndorseSkill, nil,
		func(_ *rand.Rand) {
			weightMsgEndorseSkill = defaultWeightMsgEndorseSkill
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEndorseSkill,
		profilesimulation.SimulateMsgEndorseSkill(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateSkillEndorsement int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateSkillEndorsement, &weightMsgCreateSkillEndorsement, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSkillEndorsement = defaultWeightMsgCreateSkillEndorsement
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSkillEndorsement,
		profilesimulation.SimulateMsgCreateSkillEndorsement(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSkillEndorsement int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateSkillEndorsement, &weightMsgUpdateSkillEndorsement, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSkillEndorsement = defaultWeightMsgUpdateSkillEndorsement
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSkillEndorsement,
		profilesimulation.SimulateMsgUpdateSkillEndorsement(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteSkillEndorsement int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteSkillEndorsement, &weightMsgDeleteSkillEndorsement, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSkillEndorsement = defaultWeightMsgDeleteSkillEndorsement
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSkillEndorsement,
		profilesimulation.SimulateMsgDeleteSkillEndorsement(am.accountKeeper, am.bankKeeper, am.keeper),
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
			opWeightMsgCreateUserProfile,
			defaultWeightMsgCreateUserProfile,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgCreateUserProfile(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateUserProfile,
			defaultWeightMsgUpdateUserProfile,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgUpdateUserProfile(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteUserProfile,
			defaultWeightMsgDeleteUserProfile,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgDeleteUserProfile(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateUserSkill,
			defaultWeightMsgCreateUserSkill,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgCreateUserSkill(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateUserSkill,
			defaultWeightMsgUpdateUserSkill,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgUpdateUserSkill(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteUserSkill,
			defaultWeightMsgDeleteUserSkill,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgDeleteUserSkill(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgEndorseSkill,
			defaultWeightMsgEndorseSkill,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgEndorseSkill(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateSkillEndorsement,
			defaultWeightMsgCreateSkillEndorsement,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgCreateSkillEndorsement(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateSkillEndorsement,
			defaultWeightMsgUpdateSkillEndorsement,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgUpdateSkillEndorsement(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteSkillEndorsement,
			defaultWeightMsgDeleteSkillEndorsement,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgDeleteSkillEndorsement(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
