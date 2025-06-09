package simulation

import (
	"math/rand"

	"skillchain/x/profile/keeper"
	"skillchain/x/profile/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgEndorseSkill(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgEndorseSkill{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the EndorseSkill simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "EndorseSkill simulation not implemented"), nil, nil
	}
}
