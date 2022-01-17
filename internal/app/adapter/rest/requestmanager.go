package rest

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"

	"github.com/godddesign/todo/list/internal/app/cqrs/bus/nats"
	"github.com/godddesign/todo/list/internal/app/cqrs/command"
	"github.com/godddesign/todo/list/internal/base"
)

type (
	RequestManager struct {
		*base.BaseWorker
		cqrs *base.CQRSManager
		nats *nats.Client
	}
)

func NewRequestManager(cqrs *base.CQRSManager, natsClient *nats.Client, log base.Logger) (rm *RequestManager) {
	return &RequestManager{
		BaseWorker: base.NewWorker("request-manager", log),
		cqrs:       cqrs,
		nats:       natsClient,
	}

}

// Dispatch is a WIP: This can be improved,
// Maybe delegate the action to the command itself right here
// but removing first the concrete actions from the OpenAPI specifications.
// so the handlers don't have to adjust to the interface (response/request interface)
// and code and checks duplication is avoided.
func (rm *RequestManager) Dispatch(w http.ResponseWriter, r *http.Request, commandName string) {
	reqID := genReqID(r)

	// TODO: Sending raw body for now
	// See for a better approach later.
	payload, err := body(r)
	if err != nil {
		err := fmt.Errorf("send command error: %w", err)
		rm.Error(err, w)
	}

	switch commandName {
	case "create-list":
		rm.SendCommand("create-list-command", payload, reqID)

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

func (rm *RequestManager) SendCommand(cmd string, payload []byte, tracingID string) {
	ce, err := encodeCommand(cmd, payload, tracingID)
	if err != nil {
		rm.Log().Errorf("Send command error: %s", err.Error())
		return
	}

	rm.nats.PublishCommand(cmd, ce)
}

//func (rm *RequestManager) Receive(commandName string, payload []byte) {
//	switch commandName {
//	case "create-list":
//		rm.CreateList(payload)
//
//	//case "update-list":
//	//	rm.UpdateList(w, r)
//
//	//case "delete-list":
//	//	rm.DeleteList(w, r)
//
//	//case "get-all-lists":
//	//	rm.GetAllLists(w, r)
//
//	default:
//		err := fmt.Errorf("command '%s' not found", commandName)
//		rm.Log().Error(err.Error())
//		// NOTE: Send an error event?
//	}
//}

//func (rm *RequestManager) CreateList(payload []byte) {
//	cmd, ok := rm.cqrs.FindCommand(name)
//	if !ok {
//		err := fmt.Errorf("command '%s' not found", name)
//		rm.Log().Errorf("Create list command error: %w", err)
//	}
//
//	switch cmd := cmd.(type) {
//	case *command.CreateListCommand:
//		data, err := ToCreateListCommandData(r)
//		if err != nil {
//			err := fmt.Errorf("wrong '%s' data: %+v", cmd.Name(), data)
//			rm.Error(err, w)
//			return
//		}
//
//		ctx := context.Background()
//		err = cmd.HandleFunc()(ctx, data)
//		if err != nil {
//			err := fmt.Errorf("error: %s", err.Error())
//			rm.Error(err, w)
//			return
//		}
//
//	default:
//		rm.Log().Errorf("wrong command: %+v", cmd)
//	}
//}

func (rm *RequestManager) CreateListReqRes(w http.ResponseWriter, r *http.Request) {
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

// Helpers
func body(r *http.Request) (body []byte, err error) {
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}

func genReqID(r *http.Request) (id string) {
	id = r.Header.Get("X-Request-ID")
	if id == "" {
		return uuid.New().String()
	}

	return id
}

func encodeCommand(cmd string, payload []byte, tracingID string) (cmdEvent []byte, err error) {
	ce := nats.CommandEvent{
		Command:   cmd,
		Payload:   payload,
		TracingID: tracingID,
	}

	buf := bytes.Buffer{}
	err = gob.NewEncoder(&buf).Encode(ce)
	if err != nil {
		return cmdEvent, err
	}

	cmdEvent, err = ioutil.ReadAll(&buf)
	if err != nil {
		return cmdEvent, err
	}

	return cmdEvent, err
}
