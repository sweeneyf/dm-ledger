package permission

import (
	"github.com/sweeneyf/dm-ledger/x/permission/keeper"
	"github.com/sweeneyf/dm-ledger/x/permission/types"
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
	// TODO: Fill out function aliases

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

	// MsgRegsiter -
	MsgRegister = types.MsgRegister
	// MsgAccessRequest -
	MsgAccessRequest = types.MsgAccessRequest
	// MsgCreatepermission -
	MsgCreatepermission = types.MsgCreatepermission
	// MsgDeletepermission -
	MsgDeletepermission = types.MsgDeletepermission

	// AccessPermission -
	AccessPermission = types.AccessPermission
	// Policy -
	Policy = types.Policy
)
