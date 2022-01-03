package jsonapi

import (
	"encoding/json"
	"net/http"

	"github.com/adrianpk/godddtodo/internal/app/cqrs/command"
	"github.com/adrianpk/godddtodo/internal/base"
)

type (
	RequestManager struct {
		*base.BaseWorker
		cqrs *base.CQRSManager
	}
)

func NewRequestManager(cqrs *base.CQRSManager, log base.Logger) (rm *RequestManager) {
	return &RequestManager{
		BaseWorker: base.NewWorker("request-manager", log),
		cqrs:       cqrs,
	}

}

func (rm *RequestManager) CreateList(w http.ResponseWriter, r *http.Request) {
	name := "create-list"
	// WIP: Hardcoded command name, implement a pre dynamic dispatcher
	cmd, ok := rm.cqrs.FindCommand(name)
	if !ok {
		rm.Log().Errorf("command not found: %+w", cmd)
		panic("write error response")
	}

	switch cmd := cmd.(type) {
	case *command.CreateListCommand:
		data, err := ToCreateListCommandData(r)
		if err != nil {
			rm.Log().Errorf("create list error: %w", err)
		}

		err = cmd.HandleFunc()(r.Context(), data)
		if err != nil {
			rm.Log().Errorf("create list error: %w", err)
		}

	default:
		rm.Log().Errorf("wrong command: %+v", cmd)
	}
}

// ToCreateListCommandData command
// WIP: Implement, rename and move to ports(?) package
func ToCreateListCommandData(r *http.Request) (command.CreateListCommandData, error) {
	cmdData := command.CreateListCommandData{}

	err := json.NewDecoder(r.Body).Decode(&cmdData)
	if err != nil {
		return cmdData, err
	}

	return cmdData, nil
}
