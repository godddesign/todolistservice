package jsonapi

import (
	"errors"
	"net/http"

	"github.com/adrianpk/godddtodo/internal/app/cqrs/command"
	"github.com/adrianpk/godddtodo/internal/base"
)

type (
	RequestManager struct {
		cqrs *base.CQRSManager
	}
)

func NewRequestManager(cqrs *base.CQRSManager) (rm *RequestManager, err error) {
	if cqrs == nil {
		return rm, errors.New("nil CQRS manager")
	}

	return &RequestManager{
		cqrs: cqrs,
	}, nil

}

func (rm *RequestManager) CreateList(w http.ResponseWriter, r *http.Request) {
	// WIP: Hardcoded command name, implement a pre dinamic dispatcher
	c, ok := rm.cqrs.FindCommand("create-list")
	if !ok {
		// TODO: Write error response
		panic("not implemented")
	}

	switch cmd := c.(type) {
	case command.CreateListCommand:
		data := ToCreateListCommandData(r)
		cmd.Handle(r.Context(), data)

	default:
		// TODO: Write error response
		panic("not implemented")
	}
}

// WIP: Implement, rename and move to ports(?) package
func ToCreateListCommandData(r *http.Request) command.CreateListCommandData {
	// TODO: Extract data from request and create a command data
	return command.CreateListCommandData{}
}
