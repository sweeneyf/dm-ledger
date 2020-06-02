package types

// QueryResGrant Queries Result Payload for a resolve query
type QueryResGrant struct {
	Value AccessControlGrant `json:"value"`
}

// implement fmt.Stringer
func (r QueryResGrant) String() string {
	// TODO return r.Value
	return ""
}
