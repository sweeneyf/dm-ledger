package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/sweeneyf/dm-ledger/x/grant/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	grantTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	grantTxCmd.AddCommand(flags.PostCommands(
		GetCmdRequestAcess(cdc),
		GetCmdValidateToken(cdc),
	)...)

	return grantTxCmd
}

// GetCmdRequestAcess is the CLI command for requesting grant
func GetCmdRequestAcess(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "request [data-subject] [data-controller] [data-processor]",
		Short: "request grant to a subjects data with a particlular controller",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			subjectAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			controllerAddress, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}
			processorAddress, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoins(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestAccess(subjectAddress, controllerAddress, processorAddress, coins)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdValidateToken is the CLI command for requesting grant
func GetCmdValidateToken(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "validate [service provider] [token] ",
		Short: "validae a grant by token",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			serviceProvider, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			token := args[1]

			msg := types.NewMsgValidateToken(serviceProvider, token)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
