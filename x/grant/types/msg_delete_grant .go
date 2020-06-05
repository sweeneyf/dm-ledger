package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgDeleteGrant
// This creates a request to delete a grant
// -------------------------------------------------

var _ sdk.Msg = &MsgDeleteGrant{}

// MsgDeleteGrant - struct for unjailing jailed validator
type MsgDeleteGrant struct {
	Subject    sdk.AccAddress `json:"subject" yaml:"subject"`       // address of the owner of the data ~(data subject DS)
	Controller sdk.AccAddress `json:"controller" yaml:"controller"` // address of the controller of the data (DC)
	Processor  sdk.AccAddress `json:"requestor" yaml:"requestor"`   // address of the service provider requesting access (SP)
	Location   string         `json:"location" yaml:"location"`     // pointer to location of data being accessed
}

// NewMsgDeleteGrant creates a new MsgDeleteGrant instance
func NewMsgDeleteGrant(subject sdk.AccAddress, controller sdk.AccAddress, processor sdk.AccAddress, location string) MsgDeleteGrant {
	return MsgDeleteGrant{
		Subject:    subject,
		Controller: controller,
		Processor:  processor,
		Location:   location,
	}
}

// DeleteGrantConst is DeleteGrant Constant
const DeleteGrantConst = "DeleteGrant"

// Route Returns a human-readable string for the message, intended to be the name of the module
func (msg MsgDeleteGrant) Route() string { return RouterKey }

// Type - returns the message type as defined by AccessRequestConst
func (msg MsgDeleteGrant) Type() string { return AccessRequestConst }

// GetSigners - returns the address of the signers that must sign, in this case the data subject
func (msg MsgDeleteGrant) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Subject}
}

// GetSignBytes defines hopw the message gets encoded for signing, in this case marshall to sorted JSON
func (msg MsgDeleteGrant) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is used to provide some basic stateless checks on the validity of the message
func (msg MsgDeleteGrant) ValidateBasic() error {
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
