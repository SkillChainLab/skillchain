package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/SkillChainLab/skillchain/x/verification/types"
)

func CmdRejectVerificationRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reject-verification-request [request-id] [reason]",
		Short: "Reject a verification request",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get request ID and reason
			requestId := args[0]
			reason := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRejectVerificationRequest(
				clientCtx.GetFromAddress().String(),
				requestId,
				reason,
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