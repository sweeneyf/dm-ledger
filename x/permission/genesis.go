package permission

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper /* TODO: Define what keepers the module needs */, data GenesisState) {
	/*TODO add in the genesis funtion here
	for _, record := range data.AccessControlList {
		keeper..SetWhois(ctx, record.Value, record)
	} */
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
	// TODO: Define logic for exporting state
	return NewGenesisState(data.AccessControlList)
}
