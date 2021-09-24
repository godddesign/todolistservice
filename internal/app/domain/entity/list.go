package entity

import (
	"github.com/google/uuid"
)

type (
	List struct {
		id          uuid.UUID
		slug        string
		tenantID    uuid.UUID
		orgID       uuid.UUID
		ownerID     uuid.UUID
		name        string
		description string
		items       []Item
	}
)
