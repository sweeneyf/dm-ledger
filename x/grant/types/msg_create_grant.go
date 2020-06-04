package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgCreateGrant
// This creates a reques to ask for acces to the data
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgCreateGrant{}

// MsgCreateGrant - struct for unjailing jailed validator
type MsgCreateGrant struct {
	Subject    sdk.AccAddress `json:"subject" yaml:"subject"`       // address of the owner of the data ~(data subject DS)
	Controller sdk.AccAddress `json:"controller" yaml:"controller"` // address of the controller of the data (DC)
	Processor  sdk.AccAddress `json:"requestor" yaml:"requestor"`   // address of the service provider requesting access (SP)
	AccessType string         `json:"accessType" yaml:"accessType"` // what type of access read/write/delete
	Location   string         `json:"location" yaml:"location"`     // pointer to location of data being accessed
	Reward     sdk.Coins      `json:"reward" yaml:"reward"`         // reward for allowing access
}

// NewMsgCreateGrant creates a new MsgCreateGrant instance
func NewMsgCreateGrant(subject sdk.AccAddress, controller sdk.AccAddress, processor sdk.AccAddress, accessType, location string, reward sdk.Coins) MsgCreateGrant {
	return MsgCreateGrant{
		Subject:    subject,
		Controller: controller,
		Processor:  processor,
		AccessType: accessType,
		Location:   location,
		Reward:     reward,
	}
}

// CreateGrantConst is AccessRequest Constant
const CreateGrantConst = "CreateGrant"

// Route Returns a human-readable string for the message, intended to be the name of the module
func (msg MsgCreateGrant) Route() string { return RouterKey }

// Type - returns the message type as defined by AccessRequestConst
func (msg MsgCreateGrant) Type() string { return AccessRequestConst }

// GetSigners - returns the address of the signers that must sign, in this case the data subject
func (msg MsgCreateGrant) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Subject}
}

// GetSignBytes defines hopw the message gets encoded for signing, in this case marshall to sorted JSON
func (msg MsgCreateGrant) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is used to provide some basic stateless checks on the validity of the message
func (msg MsgCreateGrant) ValidateBasic() error {
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
