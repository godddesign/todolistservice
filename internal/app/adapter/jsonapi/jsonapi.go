package jsonapi

import (
	"net/http"

	"github.com/adrianpk/godddtodo/internal/base"
)

type (
	Server struct {
		*base.Server
		*base.CQRSManager
		Config *Config
		Router http.Handler
		log    base.Logger
	}

	Config struct{}
)

func NewServer(name string, cfg *Config, log base.Logger) (server *Server, err error) {
	jas, err := base.NewServer(name, log)
	if err != nil {
		return nil, err
	}

	return &Server{
		Server: jas,
		Config: cfg,
	}, nil
}
