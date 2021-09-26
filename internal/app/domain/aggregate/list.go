package aggregate

import (
	"github.com/adrianpk/cirrustodo/internal/app/domain/entity"
	"github.com/google/uuid"
)

type (
	List struct {
		id          uuid.UUID
		userID      uuid.UUID
		slug        string
		tenantID    uuid.UUID
		orgID       uuid.UUID
		ownerID     uuid.UUID
		name        string
		description string
		items       []*entity.Item
	}
)
