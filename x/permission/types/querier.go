package types

import "strings"

const (
	// QueryPermissionDetail -
	QueryPermissionDetail = "detail"
	// QueryPermissionList - a list of a access permissions
	QueryPermissionList = "list"
)

// QueryResPermission Queries Result Payload for a resolve query
type QueryResPermission struct {
	Value Permission `json:"permission"`
}

// implement fmt.Stringer
func (r QueryResPermission) String() string {
	// TODO return r.Value
	return ""
}

// QueryResPermissions Queries Result Payload for a permissions query
type QueryResPermissions []string

// implement the fmt.Stringer interface
func (n QueryResPermissions) String() string {
	return strings.Join(n[:], "\n")
}
