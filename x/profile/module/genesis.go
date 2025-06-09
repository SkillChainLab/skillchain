package profile

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"skillchain/x/profile/keeper"
	"skillchain/x/profile/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the userProfile
	for _, elem := range genState.UserProfileList {
		k.SetUserProfile(ctx, elem)
	}
	// Set all the userSkill
	for _, elem := range genState.UserSkillList {
		k.SetUserSkill(ctx, elem)
	}
	// Set all the skillEndorsement
	for _, elem := range genState.SkillEndorsementList {
		k.SetSkillEndorsement(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.UserProfileList = k.GetAllUserProfile(ctx)
	genesis.UserSkillList = k.GetAllUserSkill(ctx)
	genesis.SkillEndorsementList = k.GetAllSkillEndorsement(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
