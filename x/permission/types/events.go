package types

// permission module event types
const (
	EventTypeCreatepermission = "Createpermission"
	EventTypeDeletepermission = "Deletepermission"
	EventTypeAccessRequest    = "AccessRequest"

	AttributeSubject     = "subject"
	AttributeController  = "controller"
	AttributeProcessor   = "processor"
	AttributeDataPointer = "dataPointer"
	AttributeAccessType  = "acessType"
	AttributeReward      = "reward"

	AttributeValueCategory = ModuleName
)
