package jsonapi

import (
	"net/http"

	"github.com/adrianpk/godddtodo/internal/base"
)

func (server *Server) InitJSONAPIRouter(h http.Handler) {
	r := base.NewRouter("json-api-router", server.Log())
	r.Mount("/", h)

	server.SetRouter(r)
}
