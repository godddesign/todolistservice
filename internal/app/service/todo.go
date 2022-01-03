// Package service provides application resources for managing todo lists.
package service

import (
	"errors"

	"github.com/adrianpk/godddtodo/internal/app/adapter/driver/repo"
	"github.com/adrianpk/godddtodo/internal/app/domain/service"
	"github.com/adrianpk/godddtodo/internal/base"
)

type (
	Todo struct {
		*base.BaseWorker
		config      Config
		cqrs        *base.CQRSManager
		repoRead    repo.ListRead
		repoWrite   repo.ListWrite
		listService *service.List
	}
)

type (
	Config struct {
		TracingLevel string
	}
)

func NewTodo(name string, rr repo.ListRead, rw repo.ListWrite, cfg Config, log base.Logger) (Todo, error) {
	var svc Todo

	if rr == nil {
		return svc, errors.New("no read repo")
	}

	if rw == nil {
		return svc, errors.New("no write repo")
	}

	svc = Todo{
		BaseWorker: base.NewWorker(name, log),
		config:     cfg,
		repoRead:   rr,
		repoWrite:  rw,
	}

	return svc, nil
}

func (t *Todo) CreateList(name, description string) error {
	t.Log().Infof("CreateList name: '%s', description: '%s'", name, description)
	return nil
}
