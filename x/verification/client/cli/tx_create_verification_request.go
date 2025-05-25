package cli

import (
	"github.com/SkillChainLab/skillchain/x/verification/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateVerificationRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-verification-request [institution-address] [skills] [evidence]",
		Short: "Create a new verification request",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get arguments
			institutionAddress := args[0]
			skills := args[1]
			evidence := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVerificationRequest(
				clientCtx.GetFromAddress().String(),
				institutionAddress,
				[]string{skills},
				evidence,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
