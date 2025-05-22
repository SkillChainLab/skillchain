package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/SkillChainLab/skillchain/x/verification/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdQueryVerificationRequest(),
		GetCmdQueryAllVerificationRequests(),
		GetCmdQueryVerifiedInstitution(),
		GetCmdQueryAllVerifiedInstitutions(),
	)

	return cmd
}

// GetCmdQueryVerificationRequest implements the query verification request command
func GetCmdQueryVerificationRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verification-request [request-id]",
		Short: "Query a verification request by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.VerificationRequest(cmd.Context(), &types.QueryVerificationRequestRequest{
				RequestId: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAllVerificationRequests implements the query all verification requests command
func GetCmdQueryAllVerificationRequests() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verification-requests",
		Short: "Query all verification requests",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.VerificationRequestAll(cmd.Context(), &types.QueryAllVerificationRequestRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryVerifiedInstitution implements the query verified institution command
func GetCmdQueryVerifiedInstitution() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verified-institution [address]",
		Short: "Query a verified institution by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.VerifiedInstitution(cmd.Context(), &types.QueryVerifiedInstitutionRequest{
				Address: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAllVerifiedInstitutions implements the query all verified institutions command
func GetCmdQueryAllVerifiedInstitutions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verified-institutions",
		Short: "Query all verified institutions",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.VerifiedInstitutionAll(cmd.Context(), &types.QueryAllVerifiedInstitutionRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
} 