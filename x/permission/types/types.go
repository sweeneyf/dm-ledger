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

// NewPermission creates a new Permission instance
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

// UpdatePolicy - returns a policy intialised with just accessed granted to the subject
func (p *Policy) UpdatePolicy(processor sdk.AccAddress, create, read, update, delete bool) {
	p.Create = updateAccList(p.Create, processor, create)
	p.Read = updateAccList(p.Read, processor, read)
	p.Update = updateAccList(p.Update, processor, update)
	p.Delete = updateAccList(p.Delete, processor, delete)

}

// UpdateAccList takes a accList of SDk addresses and checks if an address is there
// depending on whther the adress is required it is added or deleted
func updateAccList(accList []sdk.AccAddress, acc sdk.AccAddress, required bool) []sdk.AccAddress {
	// first look for the item
	found := false
	for i, item := range accList {
		if item.String() == acc.String() {
			found = true
			if required == false { // thern we need to delete it
				accList[i] = accList[len(accList)-1] // Copy last element to index i.
				accList[len(accList)-1] = nil        // Erase last element (write zero value).
				accList = accList[:len(accList)-1]
			}
			break
		}
	}
	if !found && required == true { // then we need to add it
		accList = append(accList, acc)
	}
	return accList
}
