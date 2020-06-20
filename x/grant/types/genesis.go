package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// GenesisState - all permission state that must be provided at genesis
type GenesisState struct {
	AccessControlList []Grant `json:"access_control_list"`
}

// NewGenesisState - Create a new empty access control list
func NewGenesisState(accessControlList []Grant) GenesisState {
	return GenesisState{AccessControlList: nil}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		AccessControlList: []Grant{},
	}
}

// ValidateGenesis validates the permission genesis parameters
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.AccessControlList {
		if record.Token != "" {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "token cannot be empty")
		}
	}
	return nil
}
