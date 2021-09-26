package repo

import (
	"context"

	"github.com/adrianpk/cirrustodo/internal/app/domain/aggregate"
	"github.com/google/uuid"
)

type (
	// ListRepo interface
	ListWrite interface {
		Create(ctx context.Context, list *aggregate.List) error
		Update(ctx context.Context, list *aggregate.List) (err error)
		Delete(ctx context.Context, uid uuid.UUID) error
		DeleteBySlug(ctx context.Context, slug string) error
		DeleteByName(ctx context.Context, listname string) error
	}

	ListRead interface {
		GetAll(ctx context.Context) (lists []*aggregate.List, err error)
		Get(ctx context.Context, uid uuid.UUID) (list *aggregate.List, err error)
		GetBySlug(ctx context.Context, slug string) (list *aggregate.List, err error)
		GetByName(ctx context.Context, listname string) (*aggregate.List, error)
	}
)
