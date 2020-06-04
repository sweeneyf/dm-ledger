package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// AccessControlGrant - hold the access control information for data
type AccessControlGrant struct {
	Subject    sdk.AccAddress `json:"subject"`
	Controller sdk.AccAddress `json:"controller"`
	Processor  sdk.AccAddress `json:"processor"`
	Datasets   []Dataset      `json:"dataset"`
}

// Dataset hold information of the dataset being permissioned
type Dataset struct {
	Location string `json:"location"`
	EncrKey  string `json:"encrKey"`
	Policy   Policy `json:"policy"`
}

// Policy for managing data access
type Policy struct {
	AccessType string `json:"accessType"`
}
