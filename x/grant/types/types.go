package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccessControlGrant - hold the access control information for data
type AccessControlGrant struct {
	Subject    sdk.AccAddress `json:"subject"`
	Controller sdk.AccAddress `json:"controller"`
	Processor  sdk.AccAddress `json:"processor"`
	GDPRData   GDPRData       `json:"gdpr_dataset"`
}

// GDPRData hold information of the data being permissioned
type GDPRData struct {
	Location string `json:"location"`
	EncrKey  string `json:"encrKey"`
	Policy   Policy `json:"policy"`
}

// Policy for managing data access
type Policy struct {
	AccessType string `json:"accessType"`
}

/* AddGDPRData - adds data to the accesscontrolgrant
func (a *AccessControlGrant) AddGDPRData(location string, encrKey string, policy Policy) {
	xgData := GDPRData{
		Location: location,
		EncrKey:  encrKey,
		Policy:   policy,
	}
	a.GDPRDataset[location] = &xgData
}

// RemoveGDPRData - removes data from the accesscontrolgrant
func (a *AccessControlGrant) RemoveGDPRData(location string, encrKey string) error {
	_, ok := a.GDPRDataset[location]
	if ok {
		delete(a.GDPRDataset, location)
	} else {
		return ErrGrantLocationDoesNotExist
	}
	return nil
} */
