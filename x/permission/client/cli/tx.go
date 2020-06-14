package cli

import (
	"bufio"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/sweeneyf/dm-ledger/x/permission/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	permissionTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	permissionTxCmd.AddCommand(flags.PostCommands(
		GetCmdRequestAcess(cdc),
		GetCmdCreatePermission(cdc),
		GetCmdDeletePermission(cdc),
	)...)

	return permissionTxCmd
}

// GetCmdCreatePermission is the CLI command for registering a data subject with a data controller
func GetCmdCreatePermission(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create [data-subject] [data-controller] [data-pointer] [data-hash]",
		Short: "register a data subjects data with a particlular controller",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			/*subjectAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}*/

			controllerAddress, err := getAccAddress(cliCtx.Input, args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePermission(cliCtx.GetFromAddress(), controllerAddress, args[2], args[3])
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdRequestAcess is the CLI command for requesting access
func GetCmdRequestAcess(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "request-access [data-subject] [data-controller]",
		Short: "request access to a subjects data with a particlular controller",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgAccessRequest(cliCtx.GetFromAddress(), cliCtx.GetFromAddress(), cliCtx.GetFromAddress(), args[2], args[3], coins)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdDeletePermission is the CLI command for creating a permission on a subjects data
func GetCmdDeletePermission(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [data-subject] [data-controller]",
		Short: "delete permission to a subjects data location with a particlular controller",
		Args:  cobra.ExactArgs(2),
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

			msg := types.NewMsgDeletePermission(subjectAddress, controllerAddress)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// getAccAddress returns an account address and Keybase name given an account identifier of either an address or key name passed
func getAccAddress(input io.Reader, accID string) (sdk.AccAddress, error) {
	if accID == "" {
		return nil, nil
	}

	keybase, err := keys.NewKeyring(sdk.KeyringServiceName(),
		viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), input)
	if err != nil {
		return nil, err
	}

	var info keys.Info
	if addr, err := sdk.AccAddressFromBech32(accID); err == nil {
		info, err = keybase.GetByAddress(addr)
		if err != nil {
			return nil, err
		}
	} else {
		info, err = keybase.Get(accID)
		if err != nil {
			return nil, err
		}
	}

	return info.GetAddress(), nil
}
