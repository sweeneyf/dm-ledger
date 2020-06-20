package types

import "strings"

const (
	// QueryGrantDetail -
	QueryGrantDetail = "detail"
	// QueryGrantList - a list of a grant permissions
	QueryGrantList = "list"
)

// QueryResGrant Queries Result Payload for a resolve query
type QueryResGrant struct {
	Value Grant `json:"grant"`
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
