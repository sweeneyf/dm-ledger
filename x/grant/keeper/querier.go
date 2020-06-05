package keeper

import (
	"fmt"

	"github.com/sweeneyf/dm-ledger/x/grant/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryGrantDetail:
			return queryGrant(ctx, path[1:], req, keeper)
		case types.QueryGrantList:
			return queryGrants(ctx, req, keeper)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown grant query endpoint")
		}
	}
}

// nolint: unparam
func queryGrant(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	value, err := keeper.GetAccessControlGrant(ctx, path[0])

	if err != nil {
		return []byte{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("could not query grant detail -%v", err))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResGrant{Value: *value})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// this function returns a list of all the keys for grants
func queryGrants(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var grantsList types.QueryResGrants

	iterator := keeper.GetGrantsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		grantsList = append(grantsList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, grantsList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
