package jsonapi

import (
	"errors"

	"github.com/adrianpk/cirrus"
	"github.com/adrianpk/cirrustodo/internal/app/service"
)

type (
	Server struct {
		*cirrus.Server
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

	jas, err := cirrus.NewServer(name, cfg.TracingLevel)
	if err != nil {
		return nil, err
	}

	return &Server{
		Server: jas,
		Config: cfg,
		Todo:   ts,
	}, nil
}
