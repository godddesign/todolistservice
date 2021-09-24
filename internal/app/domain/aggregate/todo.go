package aggregate

import (
	"github.com/adrianpk/cirrustodo/internal/app/domain/entity"
	"github.com/google/uuid"
)

type (
	todo struct {
		userID uuid.UUID
		list   *entity.List
		items  []*entity.Item
	}
)
