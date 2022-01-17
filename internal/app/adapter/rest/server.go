package rest

import (
	"net/http"

	"github.com/godddesign/todo/list/internal/app/config"
	"github.com/godddesign/todo/list/internal/base"
)

type (
	Server struct {
		*base.Server
		*base.CQRSManager
		Config *config.Config
		Router http.Handler
		log    base.Logger
	}
)

func NewServer(name string, cfg *config.Config, log base.Logger) (server *Server) {
	return &Server{
		Server: base.NewServer(name, log),
		Config: cfg,
	}
}
