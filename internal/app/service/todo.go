// Package service provides application primitives for managing todo lists.
package service

import (
	"errors"
	"net/http"

	"github.com/adrianpk/cirrus"
	"github.com/adrianpk/cirrustodo/internal/app/adapter/repo"
	"github.com/adrianpk/cirrustodo/internal/app/domain/service"
)

type (
	Todo struct {
		cirrus.Service
		repoRead      repo.ListRead
		repoWrite     repo.ListWrite
		domainService *service.Todo
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
		Service:   cirrus.NewService(name, cfg.TracingLevel),
		repoRead:  rr,
		repoWrite: rw,
	}

	return svc, nil
}

func (todo *Todo) CreateList(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
