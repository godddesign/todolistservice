package rest

import (
	"net/http"

	"github.com/godddesign/todo/list/internal/base"
)

func (server *Server) InitRESTRouter(h http.Handler) {
	r := base.NewRouter("rest-router", server.Log())
	r.Mount("/api/v1", h)

	server.SetRouter(r)
}
