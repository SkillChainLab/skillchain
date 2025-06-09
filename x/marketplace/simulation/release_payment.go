package simulation

import (
	"math/rand"

	"skillchain/x/marketplace/keeper"
	"skillchain/x/marketplace/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgReleasePayment(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgReleasePayment{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ReleasePayment simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "ReleasePayment simulation not implemented"), nil, nil
	}
}
