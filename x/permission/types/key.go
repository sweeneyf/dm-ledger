package types

const (
	// ModuleName is the name of the module
	ModuleName = "permission"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName

	// PermissionPrefix allows us to to recognize that the hash is a permission hash by checking first 2 bytes of key
	PermissionPrefix = "pm-"

	// AccessGrantPrefix allows us to to recognize that the hash is a acessToken hash by checking first 2 bytes of key
	AccessGrantPrefix = "ag-"
)
