package types

// permission module event types
const (
	EventTypeCreatePermission = "CreatePermission"
	EventTypeUpdatePermission = "UpdatePermission"
	EventTypeDeletePermission = "DeletePermission"
	EventTypeAccessRequest    = "AccessRequest"

	AttributeSubject     = "subject"
	AttributeController  = "controller"
	AttributeProcessor   = "processor"
	AttributeDataPointer = "dataPointer"
	AttributeDataHash    = "dataHash"
	AttributeReward      = "reward"

	AttributeValueCategory = ModuleName
)
