package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "profile",
		Short:                      "Profile transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateProfile())
	cmd.AddCommand(CmdUpdateProfile())
	cmd.AddCommand(CmdDeleteProfile())

	return cmd
}

func GetQueryCmd(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "profile",
		Short:                      "Querying commands for the profile module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdShowProfile())
	cmd.AddCommand(CmdListProfile())

	return cmd
}
