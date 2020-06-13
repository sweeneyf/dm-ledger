package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgCreatepermission
// This creates a reques to ask for acces to the data
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgCreatepermission{}

// MsgCreatepermission - struct for unjailing jailed validator
type MsgCreatepermission struct {
	Subject    sdk.AccAddress `json:"subject" yaml:"subject"`       // address of the owner of the data ~(data subject DS)
	Controller sdk.AccAddress `json:"controller" yaml:"controller"` // address of the controller of the data (DC)
	Processor  sdk.AccAddress `json:"requestor" yaml:"requestor"`   // address of the service provider requesting access (SP)
	AccessType string         `json:"accessType" yaml:"accessType"` // what type of access read/write/delete
	Location   string         `json:"location" yaml:"location"`     // pointer to location of data being accessed
	Reward     sdk.Coins      `json:"reward" yaml:"reward"`         // reward for allowing access
}

// NewMsgCreatepermission creates a new MsgCreatepermission instance
func NewMsgCreatepermission(subject sdk.AccAddress, controller sdk.AccAddress, processor sdk.AccAddress, accessType, location string, reward sdk.Coins) MsgCreatepermission {
	return MsgCreatepermission{
		Subject:    subject,
		Controller: controller,
		Processor:  processor,
		AccessType: accessType,
		Location:   location,
		Reward:     reward,
	}
}

// CreatepermissionConst is AccessRequest Constant
const CreatepermissionConst = "Createpermission"

// Route Returns a human-readable string for the message, intended to be the name of the module
func (msg MsgCreatepermission) Route() string { return RouterKey }

// Type - returns the message type as defined by AccessRequestConst
func (msg MsgCreatepermission) Type() string { return AccessRequestConst }

// GetSigners - returns the address of the signers that must sign, in this case the data subject
func (msg MsgCreatepermission) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Subject}
}

// GetSignBytes defines hopw the message gets encoded for signing, in this case marshall to sorted JSON
func (msg MsgCreatepermission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is used to provide some basic stateless checks on the validity of the message
func (msg MsgCreatepermission) ValidateBasic() error {
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
