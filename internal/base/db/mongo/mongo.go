// Package mongo is a thin wrapper over official Mongo client including support
// for exponential backoff connection retries and cirrus.Service interface
// implementation.
package mongo

import (
	"fmt"
	"github.com/adrianpk/godddtodo/internal/base"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Client struct {
		base.Service
		*mongo.Client
		config Config
	}

	Config struct {
		Host       string
		Port       int
		User       string
		Pass       string
		Database   string
		MaxRetries uint64
	}
)

// NewMongoClient
// NOTE: Other config parameters should be passed
func NewMongoClient(name string, cfg Config, log base.Logger) *Client {
	return &Client{
		Service: base.NewBaseService(name, log),
		config:  cfg,
	}
}

func (c *Client) Init() (ok chan bool) {
	ok = make(chan bool)

	go func() {
		defer close(ok)

		err := c.connect()
		if err != nil {
			ok <- false
			return
		}

		c.Log().Infof("%s service initialized", c.Name())

		ok <- true
	}()

	return ok
}

func (c *Client) Start() error {
	// Do nothing for now
	return nil
}

func (c *Client) connect() error {
	r := <-c.retryConnection()

	if r.Error != nil {
		return r.Error
	}

	c.Client = r.Client

	return nil
}

func (c *Client) URL() string {
	cfg := c.config
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/auth?authSource=admin", cfg.User,
		cfg.Pass, cfg.Host, cfg.Port)
}

func (c *Client) Db() string {
	return c.config.Database
}
