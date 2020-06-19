package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgAccessRequest
// This creates a reques to ask for acces to the data
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgAccessRequest{}

// MsgAccessRequest - struct for unjailing jailed validator
type MsgAccessRequest struct {
	Subject    sdk.AccAddress `json:"subject" yaml:"subject"`       // address of the owner of the data ~(data subject DS)
	Controller sdk.AccAddress `json:"controller" yaml:"controller"` // address of the controller of the data (DC)
	Processor  sdk.AccAddress `json:"processor" yaml:"processor"`   // address of the data processor requesting access (DP)
	Reward     sdk.Coins      `json:"reward" yaml:"reward"`         // optional reward for allowing access
}

// NewMsgAccessRequest creates a new MsgAccessRequest instance
func NewMsgAccessRequest(subject sdk.AccAddress, controller sdk.AccAddress, processor sdk.AccAddress, reward sdk.Coins) MsgAccessRequest {
	return MsgAccessRequest{
		Subject:    subject,
		Controller: controller,
		Processor:  processor,
		Reward:     reward,
	}
}

// AccessRequestConst is AccessRequest Constant
const AccessRequestConst = "AccessRequest"

// Route Returns a human-readable string for the message, intended to be the name of the module
func (msg MsgAccessRequest) Route() string { return RouterKey }

// Type - returns the message type as defined by AccessRequestConst
func (msg MsgAccessRequest) Type() string { return AccessRequestConst }

// GetSigners - returns the address of the signers that must sign, in this case te access requestors
func (msg MsgAccessRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Processor)}
}

// GetSignBytes defines hopw the message gets encoded for signing, in this case marshall to sorted JSON
func (msg MsgAccessRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is used to provide some basic stateless checks on the validity of the message
func (msg MsgAccessRequest) ValidateBasic() error {
	if msg.Processor.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "processor cannot be empty")
	}
	if msg.Subject.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "subject cannot be empty")
	}
	if msg.Controller.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "controller cannot be empty")
	}
	return nil
}
