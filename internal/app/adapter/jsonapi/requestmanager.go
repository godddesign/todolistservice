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

// Dispatch is a WIP: This can be improved,
// Maybe delegate the action to the command itself right here
// but removing first the concrete actions from the OpenAPI specifications.
// so the handlers don't have to adjust to the interface (response/request interface)
// and code and checks duplication is avoided.
func (rm *RequestManager) Dispatch(w http.ResponseWriter, r *http.Request, commandName string) {
	//path := strings.Split(r.URL.Path, "/")
	//name := path[len(path)-1]

	switch commandName {
	case "create-list":
		rm.CreateList(w, r)

	//case "update-list":
	//	rm.UpdateList(w, r)

	//case "delete-list":
	//	rm.DeleteList(w, r)

	//case "get-all-lists":
	//	rm.GetAllLists(w, r)

	default:
		err := fmt.Errorf("command '%s' not found", commandName)
		rm.Error(err, w)
	}
}

func (rm *RequestManager) CreateList(w http.ResponseWriter, r *http.Request) {
	name := "create-list"

	cmd, ok := rm.cqrs.FindCommand(name)
	if !ok {
		err := fmt.Errorf("command '%s' not found", name)
		rm.Error(err, w)
		return
	}

	switch cmd := cmd.(type) {
	case *command.CreateListCommand:
		data, err := ToCreateListCommandData(r)
		if err != nil {
			err := fmt.Errorf("wrong '%s' data: %+v", cmd.Name(), data)
			rm.Error(err, w)
			return
		}

		err = cmd.HandleFunc()(r.Context(), data)
		if err != nil {
			err := fmt.Errorf("error: %s", err.Error())
			rm.Error(err, w)
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

func (rm *RequestManager) Error(err error, w http.ResponseWriter) {
	rm.Log().Error(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
