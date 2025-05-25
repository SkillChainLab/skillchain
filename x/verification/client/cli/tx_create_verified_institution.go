package cli

import (
	"github.com/SkillChainLab/skillchain/x/verification/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateVerifiedInstitution() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-verified-institution [address] [name] [website] [verification-categories] [verification-level]",
		Short: "Create a new verified institution",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get arguments
			address := args[0]
			name := args[1]
			website := args[2]
			verificationCategories := args[3]
			// verificationLevel := args[4] // Unused, so remove to fix build error

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVerifiedInstitution(
				clientCtx.GetFromAddress().String(),
				address,
				name,
				website,
				[]string{verificationCategories},
				uint32(0), // TODO: Parse verification level
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
