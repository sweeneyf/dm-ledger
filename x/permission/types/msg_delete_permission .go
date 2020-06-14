package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgDeletePermission
// This creates a request to delete a permission
// -------------------------------------------------

var _ sdk.Msg = &MsgDeletePermission{}

// MsgDeletePermission - struct for unjailing jailed validator
type MsgDeletePermission struct {
	Subject    sdk.AccAddress `json:"subject" yaml:"subject"`       // address of the owner of the data ~(data subject DS)
	Controller sdk.AccAddress `json:"controller" yaml:"controller"` // address of the controller of the data (DC)
}

// NewMsgDeletePermission creates a new MsgDeletePermission instance
func NewMsgDeletePermission(subject sdk.AccAddress, controller sdk.AccAddress) MsgDeletePermission {
	return MsgDeletePermission{
		Subject:    subject,
		Controller: controller,
	}
}

// DeletePermissionConst is DeletePermission Constant
const DeletePermissionConst = "DeletePermission"

// Route Returns a human-readable string for the message, intended to be the name of the module
func (msg MsgDeletePermission) Route() string { return RouterKey }

// Type - returns the message type as defined by AccessRequestConst
func (msg MsgDeletePermission) Type() string { return AccessRequestConst }

// GetSigners - returns the address of the signers that must sign, in this case the data subject
func (msg MsgDeletePermission) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Subject}
}

// GetSignBytes defines hopw the message gets encoded for signing, in this case marshall to sorted JSON
func (msg MsgDeletePermission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is used to provide some basic stateless checks on the validity of the message
func (msg MsgDeletePermission) ValidateBasic() error {
	if msg.Subject.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Subject cannot be empty")
	}
	if msg.Controller.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Controller cannot be empty")
	}
	return nil
}
