package models

type (
	BasicEntity interface {
		GetName() string
		GetDescription() string
	}

	Resource interface {
		BasicEntity
		GetType() ResourceType
		GetGroup() ResourceGroup
	}
)
