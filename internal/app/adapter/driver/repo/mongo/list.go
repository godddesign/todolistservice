package mongo

import (
	"context"

	"github.com/adrianpk/godddtodo/internal/base"

	"github.com/google/uuid"

	"github.com/adrianpk/godddtodo/internal/app/domain/aggregate"
	db "github.com/adrianpk/godddtodo/internal/base/db/mongo"
)

type (
	ListRead struct {
		*Repo
	}

	ListWrite struct {
		*Repo
	}
)

const listColl = "list"

func NewListRead(name string, conn *db.Client, cfg Config, log base.Logger) *ListRead {
	return &ListRead{
		Repo: NewRepo(name, conn, listColl, cfg, log),
	}
}

func NewListWrite(name string, conn *db.Client, cfg Config, log base.Logger) *ListWrite {
	return &ListWrite{
		Repo: NewRepo(name, conn, listColl, cfg, log),
	}
}

func (lr *ListRead) GetAll(ctx context.Context) (lists []*aggregate.List, err error) {
	return lists, err
}

func (lr *ListRead) Get(ctx context.Context, uid uuid.UUID) (list *aggregate.List, err error) {
	return list, err
}

func (lr *ListRead) GetBySlug(ctx context.Context, slug string) (list *aggregate.List, err error) {
	return list, err
}

func (lr *ListRead) GetByName(ctx context.Context, name string) (list *aggregate.List, err error) {
	return list, err
}

func (lr *ListRead) GetBySlugAndToken(ctx context.Context, slug, token string) (list *aggregate.List, err error) {
	return list, err
}

func (lw *ListWrite) Create(ctx context.Context, list *aggregate.List) error {
	return nil
}

func (lw *ListWrite) Update(ctx context.Context, list *aggregate.List) (err error) {
	return err
}

func (lw *ListWrite) Delete(ctx context.Context, uid uuid.UUID) (err error) {
	return err
}

func (lw *ListWrite) DeleteBySlug(ctx context.Context, slug string) (err error) {
	return err
}

func (lw *ListWrite) DeleteByName(ctx context.Context, name string) (err error) {
	return err
}
