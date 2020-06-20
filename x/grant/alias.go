package grant

import (
	"github.com/sweeneyf/dm-ledger/x/grant/keeper"
	"github.com/sweeneyf/dm-ledger/x/grant/types"
)

// The main reason for having this file is to prevent import cycles

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	//QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// ModuleCdc -
	ModuleCdc = types.ModuleCdc
	// TODO: Fill out variable aliases
)

type (
	//Keeper -
	Keeper = keeper.Keeper
	//GenesisState -
	GenesisState = types.GenesisState
	// Params -
	Params = types.Params

	// MsgRequestAccess -
	MsgRequestAccess = types.MsgRequestAccess

	// Grant is the result of the Access Request
	Grant = types.Grant
)
