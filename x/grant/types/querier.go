package types

import "strings"

const (
	// QueryGrantDetail -
	QueryGrantDetail = "detail"
	// QueryGrantDetail - a list of a access grants
	QueryGrantList = "list"
)

// QueryResGrant Queries Result Payload for a resolve query
type QueryResGrant struct {
	Value AccessControlGrant `json:"value"`
}

// implement fmt.Stringer
func (r QueryResGrant) String() string {
	// TODO return r.Value
	return ""
}

// QueryResGrants Queries Result Payload for a grants query
type QueryResGrants []string

// implement the fmt.Stringer interface
func (n QueryResGrants) String() string {
	return strings.Join(n[:], "\n")
}
