package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/sweeneyf/dm-ledger/x/grant/types"
	"github.com/sweeneyf/dm-ledger/x/permission"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc              *codec.Codec // The wire codec for binary encoding/decoding.
	bankKeeper       bank.Keeper
	PermissionKeeper permission.Keeper
}

// NewKeeper creates new instances of the access Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, permissionKeeper permission.Keeper, bankKeeper bank.Keeper) Keeper {
	return Keeper{
		cdc:              cdc,
		storeKey:         storeKey,
		PermissionKeeper: permissionKeeper,
		bankKeeper:       bankKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetGrant returns the grant record
func (k Keeper) GetGrant(ctx sdk.Context, key string) (*types.Grant, error) {
	store := ctx.KVStore(k.storeKey)
	var item types.Grant
	byteKey := []byte(key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// SetGrant inserts or overwrites the grant
func (k Keeper) SetGrant(ctx sdk.Context, key string, value *types.Grant) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(*value)
	store.Set([]byte(key), bz)
}

// DeleteGrant revokes the grant
func (k Keeper) DeleteGrant(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key))
}

// GetGrantIterator - Get an iterator over all grants
func (k Keeper) GetGrantIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
