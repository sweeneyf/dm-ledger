package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sweeneyf/dm-ledger/x/permission/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the permission Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetPermission returns the permission record
func (k Keeper) GetPermission(ctx sdk.Context, key string) (*types.Permission, error) {
	store := ctx.KVStore(k.storeKey)
	var item types.Permission
	byteKey := []byte(key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// SetPermission inserts or overwrites the permission
func (k Keeper) SetPermission(ctx sdk.Context, key string, value *types.Permission) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(*value)
	store.Set([]byte(key), bz)
}

// DeletePermission deletes the permission
func (k Keeper) DeletePermission(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key))
}

// GetPermissionsIterator - Get an iterator over all permissions in which the keys
func (k Keeper) GetPermissionsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
