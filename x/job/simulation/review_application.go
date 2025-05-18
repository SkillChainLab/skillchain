package simulation

import (
	"math/rand"

	"github.com/SkillChainLab/skillchain/x/job/keeper"
	"github.com/SkillChainLab/skillchain/x/job/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgReviewApplication(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgReviewApplication{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ReviewApplication simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "ReviewApplication simulation not implemented"), nil, nil
	}
}
