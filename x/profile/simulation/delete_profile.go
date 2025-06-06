package simulation

import (
	"math/rand"

	"github.com/SkillChainLab/skillchain/x/profile/keeper"
	"github.com/SkillChainLab/skillchain/x/profile/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgDeleteProfile(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDeleteProfile{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the DeleteProfile simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "DeleteProfile simulation not implemented"), nil, nil
	}
}
