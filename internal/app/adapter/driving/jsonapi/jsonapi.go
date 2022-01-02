package jsonapi

import (
	"github.com/adrianpk/godddtodo/internal/base"
	"net/http"
)

type (
	Server struct {
		*base.Server
		*base.CQRSManager
		Config *Config
		Router http.Handler
	}
)

func NewServer(name string, cfg *Config) (server *Server, err error) {
	jas, err := base.NewServer(name, cfg.TracingLevel)
	if err != nil {
		return nil, err
	}

	return &Server{
		Server: jas,
		Config: cfg,
	}, nil
}
