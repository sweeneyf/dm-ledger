package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgDeletepermission
// This creates a request to delete a permission
// -------------------------------------------------

var _ sdk.Msg = &MsgDeletepermission{}

// MsgDeletepermission - struct for unjailing jailed validator
type MsgDeletepermission struct {
	Subject    sdk.AccAddress `json:"subject" yaml:"subject"`       // address of the owner of the data ~(data subject DS)
	Controller sdk.AccAddress `json:"controller" yaml:"controller"` // address of the controller of the data (DC)
	Processor  sdk.AccAddress `json:"requestor" yaml:"requestor"`   // address of the service provider requesting access (SP)
	Location   string         `json:"location" yaml:"location"`     // pointer to location of data being accessed
}

// NewMsgDeletepermission creates a new MsgDeletepermission instance
func NewMsgDeletepermission(subject sdk.AccAddress, controller sdk.AccAddress, processor sdk.AccAddress, location string) MsgDeletepermission {
	return MsgDeletepermission{
		Subject:    subject,
		Controller: controller,
		Processor:  processor,
		Location:   location,
	}
}

// DeletepermissionConst is Deletepermission Constant
const DeletepermissionConst = "Deletepermission"

// Route Returns a human-readable string for the message, intended to be the name of the module
func (msg MsgDeletepermission) Route() string { return RouterKey }

// Type - returns the message type as defined by AccessRequestConst
func (msg MsgDeletepermission) Type() string { return AccessRequestConst }

// GetSigners - returns the address of the signers that must sign, in this case the data subject
func (msg MsgDeletepermission) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Subject}
}

// GetSignBytes defines hopw the message gets encoded for signing, in this case marshall to sorted JSON
func (msg MsgDeletepermission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is used to provide some basic stateless checks on the validity of the message
func (msg MsgDeletepermission) ValidateBasic() error {
	if msg.Processor.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Processor cannot be empty")
	}
	if msg.Subject.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Subject cannot be empty")
	}
	if msg.Controller.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Controller cannot be empty")
	}
	if msg.Location == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "data location cannot empty")
	}
	return nil
}
