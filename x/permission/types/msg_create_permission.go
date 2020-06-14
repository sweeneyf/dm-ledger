package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgCreatePermission
// This creates a request to register the data subject and data pointer
// -------------------------------------------------------------------

var _ sdk.Msg = &MsgCreatePermission{}

// MsgCreatePermission - struct for unjailing jailed validator
type MsgCreatePermission struct {
	Subject     sdk.AccAddress `json:"subject" yaml:"subject"`         // address of the owner of the data ~(data subject DS)
	Controller  sdk.AccAddress `json:"controller" yaml:"controller"`   // address of the controller of the data (DC)
	DataPointer string         `json:"dataPointer" yaml:"dataPointer"` // pointer to location of data being accessed
	DataHash    string         `json:"dataHash" yaml:"dataHash"`       // hash of data being accessed
}

// NewMsgCreatePermission creates a new MsgCreatePermission instance
func NewMsgCreatePermission(subject sdk.AccAddress, controller sdk.AccAddress, dataPointer string, dataHash string) MsgCreatePermission {
	return MsgCreatePermission{
		Subject:     subject,
		Controller:  controller,
		DataPointer: dataPointer,
		DataHash:    dataHash,
	}
}

// RegistrationConst is Registration Constant
const RegistrationConst = "Registration"

// Route Returns a human-readable string for the message, intended to be the name of the module
func (msg MsgCreatePermission) Route() string { return RouterKey }

// Type - returns the message type as defined by AccessRequestConst
func (msg MsgCreatePermission) Type() string { return AccessRequestConst }

// GetSigners - returns the address of the signers that must sign, in this case the data subject
func (msg MsgCreatePermission) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Subject}
}

// GetSignBytes defines hopw the message gets encoded for signing, in this case marshall to sorted JSON
func (msg MsgCreatePermission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is used to provide some basic stateless checks on the validity of the message
func (msg MsgCreatePermission) ValidateBasic() error {
	if msg.Subject.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Subject cannot be empty")
	}
	if msg.Controller.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Controller cannot be empty")
	}
	if msg.DataPointer == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "data location cannot empty")
	}
	if msg.DataHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "data hash cannot empty")
	}
	return nil
}
