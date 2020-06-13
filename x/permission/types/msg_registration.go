package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgRegister
// This creates a request to register the data subject and data pointer
// -------------------------------------------------------------------

var _ sdk.Msg = &MsgRegister{}

// MsgRegister - struct for unjailing jailed validator
type MsgRegister struct {
	Subject     sdk.AccAddress `json:"subject" yaml:"subject"`         // address of the owner of the data ~(data subject DS)
	Controller  sdk.AccAddress `json:"controller" yaml:"controller"`   // address of the controller of the data (DC)
	DataPointer string         `json:"dataPointer" yaml:"dataPointer"` // pointer to location of data being accessed
	DataHash    string         `json:"dataHash" yaml:"dataHash"`       // hash of data being accessed
}

// NewMsgRegister creates a new MsgRegister instance
func NewMsgRegister(subject sdk.AccAddress, controller sdk.AccAddress, dataPointer string, dataHash string) MsgRegister {
	return MsgRegister{
		Subject:     subject,
		Controller:  controller,
		DataPointer: dataPointer,
		DataHash:    dataHash,
	}
}

// RegistrationConst is Registration Constant
const RegistrationConst = "Registration"

// Route Returns a human-readable string for the message, intended to be the name of the module
func (msg MsgRegister) Route() string { return RouterKey }

// Type - returns the message type as defined by AccessRequestConst
func (msg MsgRegister) Type() string { return AccessRequestConst }

// GetSigners - returns the address of the signers that must sign, in this case the data subject
func (msg MsgRegister) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Subject}
}

// GetSignBytes defines hopw the message gets encoded for signing, in this case marshall to sorted JSON
func (msg MsgRegister) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is used to provide some basic stateless checks on the validity of the message
func (msg MsgRegister) ValidateBasic() error {
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
