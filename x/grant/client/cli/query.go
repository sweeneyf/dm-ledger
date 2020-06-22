package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sweeneyf/dm-ledger/x/grant/types"
)

// GetQueryCmd -
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	grantQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the grant module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	grantQueryCmd.AddCommand(flags.GetCommands(
		GetCmdGetGrant(storeKey, cdc),
		GetCmdListGrants(storeKey, cdc),
	)...)

	return grantQueryCmd
}

// GetCmdGetGrant queries information about a grant
func GetCmdGetGrant(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "detail [id]",
		Short: "gives the detail of a grant",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGrantDetail, id), nil)
			if err != nil {
				fmt.Printf("could not query grant - %s %v\n", id, err)
				return nil
			}

			var out types.QueryResGrant
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdListGrants queries a list of all grants in the system
func GetCmdListGrants(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			fullQueryRoute := fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryGrantList)
			res, _, err := cliCtx.QueryWithData(fullQueryRoute, nil)
			if err != nil {
				fmt.Printf("could not get query grant list full route is %s - error is %v\n ", fullQueryRoute, err)
				return nil
			}

			var out types.QueryResGrants
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
