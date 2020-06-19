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

	// NewPermission - create a new permission
	NewPermission = types.NewPermission
	FindAccInACL  = types.FindAccInACL
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

	// MsgAccessRequest -
	MsgAccessRequest = types.MsgAccessRequest
	// MsgCreatePermission -
	MsgCreatePermission = types.MsgCreatePermission
	// MsgUpdatePermission -
	MsgUpdatePermission = types.MsgUpdatePermission
	// MsgDeletePermission -
	MsgDeletePermission = types.MsgDeletePermission

	// Permission -
	Permission = types.Permission
	// Policy -
	Policy = types.Policy

	// AccessGrant is the result of the Access Request
	AccessGrant = types.AccessGrant
)
