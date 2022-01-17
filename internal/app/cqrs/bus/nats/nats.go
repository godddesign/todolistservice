package nats

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"

	"github.com/godddesign/todo/list/internal/app/config"
	"github.com/godddesign/todo/list/internal/base"
)

type (
	Client struct {
		*base.BaseWorker
		config *config.Config
		conn   *nats.Conn
	}
)

const (
	defaultHost = "localhost"
	defaultPort = 4222
)

const (
	commandsSubj = "commands"
)

func NewClient(name string, cfg *config.Config, log base.Logger) *Client {
	return &Client{
		BaseWorker: base.NewWorker(name, log),
		config:     cfg,
	}
}

func (c *Client) Start() error {
	c.Log().Infof("NATS client connecting to %s", c.address())

	var err error
	c.conn, err = nats.Connect(c.address())
	if err != nil {
		return fmt.Errorf("nats connection cannot be established: %w", err)
	}

	// Subscriptions
	// WIP: Move this up, just only to verify subscriptions are working
	c.SubscribeToCommands()

	return nil
}

func (c *Client) address() (address string) {
	host := defaultHost
	if c.config.NATS.Host == "" {
		host = c.config.NATS.Host
	}

	port := defaultPort
	if c.config.NATS.Port == 0 {
		port = defaultPort
	}

	return fmt.Sprintf("nats://%s:%d", host, port)
}

func (c *Client) PublishCommand(name string, commandEvent []byte) {
	c.Log().Infof("NATS publishing through: %s", c.conn.ConnectedAddr())

	err := c.conn.Publish(commandsSubj, commandEvent)
	if err != nil {
		c.Log().Errorf("NATS command publishing error:", err.Error())
	}
}

func (c *Client) SubscribeToCommands() {
	c.Log().Infof("NATS subscribed through: %s", c.conn.ConnectedAddr())

	var err error
	_, err = c.conn.Subscribe(commandsSubj, func(m *nats.Msg) {
		buf := bytes.NewBuffer(m.Data)
		dec := gob.NewDecoder(buf)

		ce := CommandEvent{}
		err := dec.Decode(&ce)
		if err != nil {
			c.Log().Errorf("Cannot decode command event: %s", err.Error())
		}

		c.Log().Infof("Received a command event with ID: %s", ce.TracingID)
	})

	if err != nil {
		c.Log().Errorf("NATS command subscription error: %s", err.Error())
	}

	err = c.conn.Flush()
	if err != nil {
		c.Log().Errorf("NATS flush error: %s", err.Error())
	}

	err = c.conn.LastError()
	if err != nil {
		c.Log().Error(err.Error())
	}

	c.Log().Infof("Listening on '%s' subject", commandsSubj)

	runtime.Goexit()

}
