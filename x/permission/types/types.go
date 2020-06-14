package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Permission - hold the access control information for data
type Permission struct {
	Subject     sdk.AccAddress `json:"subject"`
	Controller  sdk.AccAddress `json:"controller"`
	DataPointer string         `json:"dataPointer"`
	DataHash    string         `json:"dataHash"`
	Policy      Policy         `json:"policy"`
}

// Policy is simply defined as an ACL sepcifying 4 access rights, asscoiated with each access right
// is a list of permissioned parties under their public keys
type Policy struct {
	Create []sdk.AccAddress `json:"create"`
	Read   []sdk.AccAddress `json:"read"`
	Update []sdk.AccAddress `json:"update"`
	Delete []sdk.AccAddress `json:"delete"`
}

// NewPermission creates a new MsgCreatePermission instance
func NewPermission(subject sdk.AccAddress, controller sdk.AccAddress, dataPointer string, dataHash string) Permission {
	return Permission{
		Subject:     subject,
		Controller:  controller,
		DataPointer: dataPointer,
		DataHash:    dataHash,
		Policy:      NewPolicy(subject),
	}
}

// NewPolicy - returns a policy intialised with just accessed granted to the subject
func NewPolicy(subject sdk.AccAddress) Policy {
	return Policy{
		Create: []sdk.AccAddress{subject},
		Read:   []sdk.AccAddress{subject},
		Update: []sdk.AccAddress{subject},
		Delete: []sdk.AccAddress{subject},
	}
}
