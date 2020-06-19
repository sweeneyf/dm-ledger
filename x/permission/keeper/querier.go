package keeper

import (
	"fmt"

	"github.com/sweeneyf/dm-ledger/x/permission/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryPermissionDetail:
			return getPermission(ctx, path[1:], req, keeper)
		case types.QueryPermissionList:
			return listPermissions(ctx, req, keeper)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown permission query endpoint -"+path[0])
		}
	}
}

// removePrefixFromKey removes the prefix from the key
func removePrefixFromKey(key []byte, prefix []byte) (hash []byte) {
	hash = key[len(prefix):]
	return hash
}

// nolint: unparam
func getPermission(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	value, err := keeper.GetPermission(ctx, path[0])

	if err != nil {
		return []byte{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("could not query permission detail -%v", err))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResPermission{Value: *value})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// this function returns a list of all the keys for permissions
func listPermissions(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var permissionsList types.QueryResPermissions

	iterator := keeper.GetPermissionsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		permissionsList = append(permissionsList, string(removePrefixFromKey(iterator.Key(), []byte(types.PermissionPrefix))))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, permissionsList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
