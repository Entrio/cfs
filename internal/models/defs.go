package models

import "github.com/google/uuid"

type (
	BasicEntity interface {
		GetID() uuid.UUID
		GetName() string
		GetDescription() string
	}

	Resource interface {
		BasicEntity
		GetType() ResourceType
		GetGroup() ResourceGroup
	}
)
