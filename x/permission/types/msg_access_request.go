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
	Requestor  sdk.AccAddress `json:"requestor" yaml:"requestor"`   // address of the service provider requesting access (SP)
	Owner      sdk.AccAddress `json:"owner" yaml:"owner"`           // address of the owner of the data ~(data subject DS)
	Controller sdk.AccAddress `json:"controller" yaml:"controller"` // address of the controller of the data (DC)
	AccessType string         `json:"accessType" yaml:"accessType"` // what type of access read/write/delete
	Location   string         `json:"location" yaml:"location"`     // pointer to location of data being accessed
	Reward     sdk.Coins      `json:"reward" yaml:"reward"`         // reward for allowing access
}

// NewMsgAccessRequest creates a new MsgAccessRequest instance
func NewMsgAccessRequest(requestor sdk.AccAddress, owner sdk.AccAddress, controller sdk.AccAddress, accessType, location string, reward sdk.Coins) MsgAccessRequest {
	return MsgAccessRequest{
		Requestor:  requestor,
		Owner:      owner,
		Controller: controller,
		AccessType: accessType,
		Location:   location,
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
	return []sdk.AccAddress{sdk.AccAddress(msg.Requestor)}
}

// GetSignBytes defines hopw the message gets encoded for signing, in this case marshall to sorted JSON
func (msg MsgAccessRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is used to provide some basic stateless checks on the validity of the message
func (msg MsgAccessRequest) ValidateBasic() error {
	if msg.Requestor.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "requestor cannot be empty")
	}
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner cannot be empty")
	}
	if msg.Controller.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "controller cannot be empty")
	}
	if msg.Location == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "data location cannot empty")
	}
	return nil
}
