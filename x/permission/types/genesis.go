package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// GenesisState - all permission state that must be provided at genesis
type GenesisState struct {
	AccessControlList []Permission `json:"access_control_list"`
}

// NewGenesisState - Create a new empty access control list
func NewGenesisState(accessControlList []Permission) GenesisState {
	return GenesisState{AccessControlList: nil}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		AccessControlList: []Permission{},
	}
}

// ValidateGenesis validates the permission genesis parameters
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.AccessControlList {
		if record.Subject.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "subject cannot be empty")
		}
		if record.Controller.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "controller cannot be empty")
		}
		if record.DataPointer == "" {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "data pointer cannot be empty")
		}
	}
	return nil
}
