package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgValidateToken
// This creates a reques to ask for acces to the data
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgValidateToken{}

// MsgValidateToken - struct for unjailing jailed validator
type MsgValidateToken struct {
	ServiceProvider sdk.AccAddress `json:"serviceProvider" yaml:"serviceProvider"` // the service provider(Resource Server)
	Token           string         `json:"token" yaml:"token"`                     // token identifying the grant
}

// NewMsgValidateToken creates a new MsgValidateToken instance
func NewMsgValidateToken(serviceProvider sdk.AccAddress, token string) MsgValidateToken {
	return MsgValidateToken{
		ServiceProvider: serviceProvider,
		Token:           token,
	}
}

// ValidateTokenConst is ValidateToken Constant
const ValidateTokenConst = "ValidateToken"

// Route Returns a human-readable string for the message, intended to be the name of the module
func (msg MsgValidateToken) Route() string { return RouterKey }

// Type - returns the message type as defined by ValidateTokenConst
func (msg MsgValidateToken) Type() string { return ValidateTokenConst }

// GetSigners - returns the address of the signers that must sign, in this case te access requestors
func (msg MsgValidateToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ServiceProvider)}
}

// GetSignBytes defines hopw the message gets encoded for signing, in this case marshall to sorted JSON
func (msg MsgValidateToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic is used to provide some basic stateless checks on the validity of the message
func (msg MsgValidateToken) ValidateBasic() error {
	if msg.ServiceProvider.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "service provider cannot be empty")
	}
	if msg.Token == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "token cannot be empty")
	}
	return nil
}
