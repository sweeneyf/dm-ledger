package types

// grant module event types
const (
	EventTypeRequestAccess = "RequestAccess"
	EventTypeValidateToken = "ValidateToken"

	AttributeSubject     = "subject"
	AttributeController  = "controller"
	AttributeProcessor   = "processor"
	AttributeDataPointer = "dataPointer"
	AttributeDataHash    = "dataHash"
	AttributeReward      = "reward"
	AttributeToken       = "token"

	AttributeValueCategory = ModuleName
)
