package jsonapi

import (
	"github.com/adrianpk/godddtodo/internal/base"
)

type (
	Server struct {
		*base.Server
		Config
		*base.CQRSManager
	}

	Config struct {
		TracingLevel string
	}
)

func NewServer(name string, cfg Config) (server *Server, err error) {
	jas, err := base.NewServer(name, cfg.TracingLevel)
	if err != nil {
		return nil, err
	}

	return &Server{
		Server: jas,
		Config: cfg,
	}, nil
}
