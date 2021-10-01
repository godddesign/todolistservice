package jsonapi

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (server *Server) InitJSONAPIRouter(h http.Handler) {
	r := chi.NewRouter()
	r.Mount("/", h)
}
