package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgUpdatePermission
// This updates the permission registered between data subject and data controller to the data processor
// -------------------------------------------------------------------

var _ sdk.Msg = &MsgUpdatePermission{}

// MsgUpdatePermission - struct for unjailing jailed validator
type MsgUpdatePermission struct {
	Subject    sdk.AccAddress `json:"subject" yaml:"subject"`       // address of the owner of the data ~(data subject DS)
	Controller sdk.AccAddress `json:"controller" yaml:"controller"` // address of the controller of the data (DC)
	Processor  sdk.AccAddress `json:"processor" yaml:"processor"`   // address of the processor of the data (DP)
	Create     bool           `json:"create" yaml:"create"`         // has permission create
	Read       bool           `json:"read" yaml:"read"`             // has permission read
	Update     bool           `json:"update" yaml:"update"`         // has permission update
	Delete     bool           `json:"delete" yaml:"delete"`         // has permission delete
}

// NewMsgUpdatePermission creates a new MsgUpdatePermission instance
func NewMsgUpdatePermission(subject sdk.AccAddress, controller sdk.AccAddress, processor sdk.AccAddress, create, read, update, delete bool) MsgUpdatePermission {
	return MsgUpdatePermission{
		Subject:    subject,
		Controller: controller,
		Processor:  processor,
		Create:     create,
		Read:       read,
		Update:     update,
		Delete:     delete,
	}
}

// UpdatePermissionConst is Update Constant
const UpdatePermissionConst = "Update"

// Route Returns a human-readable string for the message, intended to be the name of the module
func (msg MsgUpdatePermission) Route() string { return RouterKey }

// Type - returns the message type as defined by UpdateConst
func (msg MsgUpdatePermission) Type() string { return UpdatePermissionConst }

// GetSigners - returns the address of the signers that must sign, in this case the data subject
func (msg MsgUpdatePermission) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Subject}
}

// GetSignBytes defines hopw the message gets encoded for signing, in this case marshall to sorted JSON
func (msg MsgUpdatePermission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is used to provide some basic stateless checks on the validity of the message
func (msg MsgUpdatePermission) ValidateBasic() error {
	if msg.Subject.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Subject cannot be empty")
	}
	if msg.Controller.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Controller cannot be empty")
	}
	if msg.Processor.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Controller cannot be empty")
	}
	if !msg.Create && !msg.Read && !msg.Update && !msg.Delete {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "at least one permission type must be set")
	}
	return nil
}
