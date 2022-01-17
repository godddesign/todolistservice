package mongo

import (
	"context"
	"errors"

	"github.com/godddesign/todo/list/internal/app/config"
	"github.com/godddesign/todo/list/internal/app/core"
	"github.com/godddesign/todo/list/internal/base"

	"github.com/google/uuid"

	db "github.com/godddesign/todo/list/internal/base/db/mongo"
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

func NewListRead(name string, conn *db.Client, cfg *config.Config, log base.Logger) *ListRead {
	return &ListRead{
		Repo: NewRepo(name, conn, listColl, cfg, log),
	}
}

func NewListWrite(name string, conn *db.Client, cfg *config.Config, log base.Logger) *ListWrite {
	return &ListWrite{
		Repo: NewRepo(name, conn, listColl, cfg, log),
	}
}

func (lr *ListRead) GetAll(ctx context.Context) (lists []core.List, err error) {
	return lists, err
}

func (lr *ListRead) Get(ctx context.Context, uid uuid.UUID) (list core.List, err error) {
	return list, err
}

func (lr *ListRead) GetBySlug(ctx context.Context, slug string) (list core.List, err error) {
	return list, err
}

func (lr *ListRead) GetByName(ctx context.Context, name string) (list core.List, err error) {
	return list, err
}

func (lr *ListRead) GetBySlugAndToken(ctx context.Context, slug, token string) (list *core.List, err error) {
	return list, err
}

func (lw *ListWrite) Create(ctx context.Context, list core.List) error {
	return errors.New("not implemented")
}

func (lw *ListWrite) Update(ctx context.Context, list core.List) (err error) {
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
