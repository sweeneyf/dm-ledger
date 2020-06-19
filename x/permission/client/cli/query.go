package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sweeneyf/dm-ledger/x/permission/types"
)

// GetQueryCmd -
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	permissionQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the permission module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	permissionQueryCmd.AddCommand(flags.GetCommands(
		GetCmdGetPermission(storeKey, cdc),
		GetCmdListPermissions(storeKey, cdc),
	)...)

	return permissionQueryCmd
}

// GetCmdGetPermission queries information about a permission
func GetCmdGetPermission(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "detail [id]",
		Short: "gives the detail of a permission",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryPermissionDetail, id), nil)
			if err != nil {
				fmt.Printf("could not query permission - %s %v\n", id, err)
				return nil
			}

			var out types.QueryResPermission
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdListPermissions queries a list of all permissions in the system
func GetCmdListPermissions(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			fullQueryRoute := fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryPermissionList)
			res, _, err := cliCtx.QueryWithData(fullQueryRoute, nil)
			if err != nil {
				fmt.Printf("could not get query permission list full route is %s - error is %v\n ", fullQueryRoute, err)
				return nil
			}

			var out types.QueryResPermissions
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
