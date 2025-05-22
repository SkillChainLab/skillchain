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
			skillsStr, _ := cmd.Flags().GetString("skills")
			experiencesStr, _ := cmd.Flags().GetString("experiences")
			website, _ := cmd.Flags().GetString("website")
			github, _ := cmd.Flags().GetString("github")
			linkedin, _ := cmd.Flags().GetString("linkedin")
			twitter, _ := cmd.Flags().GetString("twitter")
			avatar, _ := cmd.Flags().GetString("avatar")
			location, _ := cmd.Flags().GetString("location")
			email, _ := cmd.Flags().GetString("email")

			// Parse skills
			var skills []string
			if skillsStr != "" {
				err = json.Unmarshal([]byte(skillsStr), &skills)
				if err != nil {
					return fmt.Errorf("invalid skills format: %w", err)
				}
			}

			// Parse experiences
			var experiences []types.Experience
			var experiencePtrs []*types.Experience
			if experiencesStr != "" {
				err = json.Unmarshal([]byte(experiencesStr), &experiences)
				if err != nil {
					return fmt.Errorf("invalid experiences format: %w", err)
				}
				// []Experience -> []*Experience dönüştür
				for i := range experiences {
					experiencePtrs = append(experiencePtrs, &experiences[i])
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
				skills,
				experiencePtrs,
				website,
				github,
				linkedin,
				twitter,
				avatar,
				location,
				email,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("skills", "", "JSON array of skills")
	cmd.Flags().String("experiences", "", "JSON array of experiences")
	cmd.Flags().String("website", "", "Website URL")
	cmd.Flags().String("github", "", "GitHub profile URL")
	cmd.Flags().String("linkedin", "", "LinkedIn profile URL")
	cmd.Flags().String("twitter", "", "Twitter profile URL")
	cmd.Flags().String("avatar", "", "Avatar image URL")
	cmd.Flags().String("location", "", "Location")
	cmd.Flags().String("email", "", "Email address")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
