package repo

import (
	"context"

	"github.com/adrianpk/cirrustodo/internal/app/domain/entity"
	"github.com/google/uuid"
)

type (
	// ListRepo interface
	ListWrite interface {
		Create(ctx context.Context, list *entity.List) error
		Update(ctx context.Context, list *entity.List) (err error)
		Delete(ctx context.Context, uid uuid.UUID) error
		DeleteBySlug(ctx context.Context, slug string) error
		DeleteByName(ctx context.Context, listname string) error
	}

	ListRead interface {
		GetAll(ctx context.Context) (lists []*entity.List, err error)
		Get(ctx context.Context, uid uuid.UUID) (list *entity.List, err error)
		GetBySlug(ctx context.Context, slug string) (list *entity.List, err error)
		GetByName(ctx context.Context, listname string) (*entity.List, error)
	}
)
