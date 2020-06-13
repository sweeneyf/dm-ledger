package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccessPermission - hold the access control information for data
type AccessPermission struct {
	Subject     sdk.AccAddress `json:"subject"`
	Controller  sdk.AccAddress `json:"controller"`
	DataPointer string         `json:"dataPointer"`
	Policy      Policy         `json:"policy"`
}

// Policy is simply defined as an ACL sepcifying 4 access rights, asscoiated with each access right
// is a list of permissioned parties under their public keys
type Policy struct {
	Create []string `json:"create"`
	Read   []string `json:"read"`
	Update []string `json:"update"`
	Delete []string `json:"delete"`
}
