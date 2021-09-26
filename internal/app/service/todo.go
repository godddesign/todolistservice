// Package service provides application resources for managing todo lists.
package service

import (
	"errors"
	"net/http"

	"github.com/adrianpk/cirrustodo/internal/app/adapter/repo"
	"github.com/adrianpk/cirrustodo/internal/app/domain/service"
	"github.com/adrianpk/cirrustodo/internal/base"
)

type (
	Todo struct {
		base.Worker
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

func (todo *Todo) CreateList(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
