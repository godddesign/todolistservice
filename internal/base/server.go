package base

import (
	"errors"
	"fmt"
	"net/http"
)

type (
	Server struct {
		*BaseWorker

		service Service
		router  *Router
	}
)

type (
	ContextKey string
)

type (
	Identifiable interface {
		GetSlug() string
	}
)

const (
	SessionCtxKey ContextKey = "session"
)

const (
	GetMethod    = "GET"
	PostMethod   = "POST"
	PutMethod    = "PUT"
	PatchMethod  = "PATCH"
	DeleteMethod = "DELETE"
)

const (
	SessionKey = "session"
)

func NewServer(name, tracingLevel string) (*Server, error) {
	server := Server{
		BaseWorker: NewWorker(name, tracingLevel),
	}

	return &server, nil
}

func (server *Server) Name() string {
	return server.Name()
}

func (server *Server) SetRouter(r *Router) {
	server.router = r
}

func (server *Server) SetService(s Service) {
	server.service = s
}

func (server *Server) Start(port int) (err error) {
	if server.router == nil {
		return errors.New("server router not set")
	}

	p := fmt.Sprintf(":%d", port)

	server.SendInfof("Server %s initializing at port %s", server.Name, port)

	err = http.ListenAndServe(p, server.router)

	return err
}

func (server *Server) Service() Service {
	return server.service
}

// Resource path functions

// IndexPath returns index path under resource root path.
func IndexPath() string {
	return ""
}

// EditPath returns edit path under resource root path.
func EditPath() string {
	return "/{id}/edit"
}

// NewPath returns new path under resource root path.
func NewPath() string {
	return "/new"
}

// ShowPath returns show path under resource root path.
func ShowPath() string {
	return "/{id}"
}

// CreatePath returns create path under resource root path.
func CreatePath() string {
	return ""
}

// UpdatePath returns update path under resource root path.
func UpdatePath() string {
	return "/{id}"
}

// DeletePath returns delete path under resource root path.
func DeletePath() string {
	return "/{id}"
}

// SignupPath returns signup path.
func SignupPath() string {
	return "/signup"
}

// LoginPath returns login path.
func LoginPath() string {
	return "/login"
}

// ResPath
func ResPath(rootPath string) string {
	return "/" + rootPath + IndexPath()
}

// ResPathEdit
func ResPathEdit(rootPath string, r Identifiable) string {
	return fmt.Sprintf("/%s/%s/edit", rootPath, r.GetSlug())
}

// ResPathNew
func ResPathNew(rootPath string) string {
	return fmt.Sprintf("/%s/new", rootPath)
}

// ResPathInitDelete
func ResPathInitDelete(rootPath string, r Identifiable) string {
	return fmt.Sprintf("/%s/%s/init-delete", rootPath, r.GetSlug())
}

// ResPathSlug
func ResPathSlug(rootPath string, r Identifiable) string {
	return fmt.Sprintf("/%s/%s", rootPath, r.GetSlug())
}

// Admin
func ResAdmin(path, adminPathPfx string) string {
	return fmt.Sprintf("/%s/%s", adminPathPfx, path)
}
