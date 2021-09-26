package jsonapi

import (
	"errors"

	"github.com/adrianpk/cirrustodo/internal/app/service"
	"github.com/adrianpk/cirrustodo/internal/base"
)

type (
	Server struct {
		*base.Server
		Config
		*service.Todo
	}

	Config struct {
		TracingLevel string
	}
)

func NewServer(name string, ts *service.Todo, cfg Config) (server *Server, err error) {
	if ts == nil {
		return server, errors.New("todo app service is nil")
	}

	jas, err := base.NewServer(name, cfg.TracingLevel)
	if err != nil {
		return nil, err
	}

	return &Server{
		Server: jas,
		Config: cfg,
		Todo:   ts,
	}, nil
}
