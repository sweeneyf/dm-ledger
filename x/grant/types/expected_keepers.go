package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	perm "github.com/sweeneyf/dm-ledger/x/permission/types"
)

// ParamSubspace defines the expected Subspace interfacace
type ParamSubspace interface {
	WithKeyTable(table params.KeyTable) params.Subspace
	Get(ctx sdk.Context, key []byte, ptr interface{})
	GetParamSet(ctx sdk.Context, ps params.ParamSet)
	SetParamSet(ctx sdk.Context, ps params.ParamSet)
}

// BankKeeper - When a module wishes to interact with another module, it is good practice to define what it will use
//as an interface so the module cannot use things that are not permitted.
type BankKeeper interface {
	SubtractCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, error)
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

// PermissionKeeper - need to be able to gte the permisison details from the permission keeper
type PermissionKeeper interface {
	GetPermission(ctx sdk.Context, key string) (*perm.Permission, error)
}
