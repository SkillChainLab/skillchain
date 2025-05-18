package cli

import (
	"encoding/json"
	"fmt"

	"github.com/SkillChainLab/skillchain/x/profile/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdUpdateProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-profile [username] [bio]",
		Short: "Update a profile",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get arguments
			username := args[0]
			bio := args[1]

			// Get optional flags
			experiencesStr, _ := cmd.Flags().GetString("experiences")

			// Parse experiences
			var experiences []types.Experience
			if experiencesStr != "" {
				err = json.Unmarshal([]byte(experiencesStr), &experiences)
				if err != nil {
					return fmt.Errorf("invalid experiences format: %w", err)
				}
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateProfile(
				clientCtx.GetFromAddress().String(),
				username,
				bio,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("experiences", "", "JSON array of experiences")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
} 