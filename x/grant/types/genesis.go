package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// GenesisState - all grant state that must be provided at genesis
type GenesisState struct {
	AccessControlList []AccessControlGrant `json:"access_control_list"`
}

// NewGenesisState - Create a new empty access control list
func NewGenesisState(accessControlList []AccessControlGrant) GenesisState {
	return GenesisState{AccessControlList: nil}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		AccessControlList: []AccessControlGrant{},
	}
}

// ValidateGenesis validates the grant genesis parameters
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.AccessControlList {
		if record.Processor.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "requestor cannot be empty")
		}
		if record.Subject.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "subject cannot be empty")
		}
		if record.Controller.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "controller cannot be empty")
		}
		if record.GDPRData.Location == "" {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "data location cannot be empty")
		}
	}
	return nil
}
