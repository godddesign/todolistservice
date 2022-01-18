package nats

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"time"

	"github.com/godddesign/todo/list/internal/app/config"
	"github.com/godddesign/todo/list/internal/base"
)

const (
	commandSubject = "command"
	querySubject   = "query"
	eventSubject   = "event"
)

type (
	// BusManager is an implementation of bus.Manager built on top of NATS.
	BusManager struct {
		*base.BaseWorker
		config *config.Config
		nats   *Client
	}
)

func NewBusManager(name string, cfg *config.Config, nats *Client, log base.Logger) *BusManager {
	return &BusManager{
		BaseWorker: base.NewWorker(name, log),
		config:     cfg,
		nats:       nats,
	}
}

func (bm *BusManager) Start() error {
	return bm.nats.Start()
}

func (bm *BusManager) SendCommand(cmd string, payload []byte, tracingID string) error {
	ce, err := encodeCommand(cmd, payload, tracingID)
	if err != nil {
		bm.Log().Errorf("Send command error: %s", err.Error())
		return err
	}

	return bm.nats.PublishEvent(commandSubject, ce)
}

func (bm *BusManager) SendEvent(event string, payload []byte, tracingID string) error {
	panic("not implemented")
}

func (bm *BusManager) Query(query string, payload []byte, timeout time.Duration, tracingID string) (response []byte, err error) {
	panic("not implemented")
}

// Helpers

func encodeCommand(cmd string, payload []byte, tracingID string) (cmdEvent []byte, err error) {
	ce := CommandEvent{
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
