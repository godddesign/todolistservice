package jsonapi

import (
	"encoding/json"
	"fmt"
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

	cmd, ok := rm.cqrs.FindCommand(name)
	if !ok {
		err := fmt.Errorf("command '%s' not found", name)

		rm.Log().Error(err.Error())

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch cmd := cmd.(type) {
	case *command.CreateListCommand:
		data, err := ToCreateListCommandData(r)
		if err != nil {
			err := fmt.Errorf("wrong '%s' data: %+v", cmd.Name(), data)

			rm.Log().Error(err.Error())

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = cmd.HandleFunc()(r.Context(), data)
		if err != nil {
			err := fmt.Errorf("error: %s", err.Error())

			rm.Log().Errorf("Request Manager %s", err.Error())

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		rm.Log().Errorf("wrong command: %+v", cmd)
	}
}

// ToCreateListCommandData command
func ToCreateListCommandData(r *http.Request) (cmdData command.CreateListCommandData, err error) {
	err = json.NewDecoder(r.Body).Decode(&cmdData)
	if err != nil {
		return cmdData, err
	}

	return cmdData, nil
}
