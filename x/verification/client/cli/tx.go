package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SkillChainLab/skillchain/x/verification/types"
	"github.com/cosmos/cosmos-sdk/client"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateVerifiedInstitution())
	cmd.AddCommand(CmdCreateVerificationRequest())
	cmd.AddCommand(CmdApproveVerificationRequest())
	cmd.AddCommand(CmdRejectVerificationRequest())

	return cmd
}
