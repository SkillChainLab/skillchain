package cli

import (
	"github.com/SkillChainLab/skillchain/x/profile/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-profile [username] [bio] [email] [wallet-address]",
		Short: "Create a new profile",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get arguments
			username := args[0]
			bio := args[1]
			email := args[2]
			walletAddress := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateProfile(
				walletAddress,
				username,
				bio,
				nil, // skills
				nil, // experiences
				"",  // website
				"",  // github
				"",  // linkedin
				"",  // twitter
				"",  // avatar
				"",  // location
				email,
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

func CmdDeleteProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-profile [username]",
		Short: "Delete a profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			username := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteProfile(
				clientCtx.GetFromAddress().String(),
				username,
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
