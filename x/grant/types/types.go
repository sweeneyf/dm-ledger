package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccessControlGrant - hold the access control information for data
type AccessControlGrant struct {
	Subject    sdk.AccAddress     `json:"subject"`
	Controller sdk.AccAddress     `json:"controller"`
	Processor  sdk.AccAddress     `json:"processor"`
	Datasets   map[string]Dataset `json:"dataset"`
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

// AddDataset - adds a dataset to the accesscontrolgrant
func (a *AccessControlGrant) AddDataset(location string, encrKey string, policy Policy) {
	a.Datasets[location] = Dataset{
		Location: location,
		EncrKey:  encrKey,
		Policy:   policy,
	}
}

// RemoveDataset - adds a dataset to the accesscontrolgrant
func (a *AccessControlGrant) RemoveDataset(location string, encrKey string) error {
	_, ok := a.Datasets[location]
	if ok {
		delete(a.Datasets, location)
	} else {
		return ErrGrantLocationDoesNotExist
	}
	return nil
}
