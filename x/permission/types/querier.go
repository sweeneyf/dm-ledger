package types

import "strings"

const (
	// QuerypermissionDetail -
	QuerypermissionDetail = "detail"
	// QuerypermissionDetail - a list of a access permissions
	QuerypermissionList = "list"
)

// QueryRespermission Queries Result Payload for a resolve query
type QueryRespermission struct {
	Value AccessPermission `json:"value"`
}

// implement fmt.Stringer
func (r QueryRespermission) String() string {
	// TODO return r.Value
	return ""
}

// QueryRespermissions Queries Result Payload for a permissions query
type QueryRespermissions []string

// implement the fmt.Stringer interface
func (n QueryRespermissions) String() string {
	return strings.Join(n[:], "\n")
}
