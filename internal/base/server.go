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

func NewServer(name string, log Logger) *Server {
	server := Server{
		BaseWorker: NewWorker(name, log),
	}

	return &server
}

func (srv *Server) Name() string {
	return srv.name
}

func (srv *Server) SetRouter(r *Router) {
	srv.router = r
}

func (srv *Server) SetService(s Service) {
	srv.service = s
}

func (srv *Server) Start(port int) (err error) {
	if srv.router == nil {
		return errors.New("server router not set")
	}

	p := fmt.Sprintf(":%d", port)

	srv.Log().Infof("Server %s initializing at port %s", srv.Name(), p)

	err = http.ListenAndServe(p, srv.router)

	return err
}

func (srv *Server) Service() Service {
	return srv.service
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
