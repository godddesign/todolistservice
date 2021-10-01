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
		base.Worker
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

func NewTodo(name string, rr repo.ListRead, rw repo.ListWrite, cfg Config) (svc Todo, err error) {
	if rr == nil {
		return svc, errors.New("nil read repo")
	}

	if rw == nil {
		return svc, errors.New("nil write repo")
	}

	svc = Todo{
		Worker:    base.NewWorker(name, cfg.TracingLevel),
		repoRead:  rr,
		repoWrite: rw,
	}

	return svc, nil
}
