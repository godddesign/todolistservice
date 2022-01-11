package nats

import (
	"fmt"

	nats "github.com/nats-io/nats.go"

	"github.com/adrianpk/godddtodo/internal/base"
)

type (
	NATSClient struct {
		*base.BaseWorker
		config Config
		conn   *nats.Conn
	}

	Config struct {
		Host string
		Port int
	}
)

const (
	defaultHost = "localhost"
	defaultPort = 4222
)

func NewNATSClient(name string, cfg Config, log base.Logger) *NATSClient {
	return &NATSClient{
		BaseWorker: base.NewWorker(name, log),
		config:     cfg,
	}

}

func (nc *NATSClient) Init() error {
	nc.Log().Infof("Starting NATS client (%s)", nc.address())

	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return fmt.Errorf("nats connection cannot be established: %w", err)
	}

	nc.conn = conn

	return nil
}

func (nc *NATSClient) address() (address string) {
	host := defaultHost
	if nc.config.Host == "" {
		host = nc.config.Host
	}

	port := defaultPort
	if nc.config.Port == 0 {
		port = defaultPort
	}

	return fmt.Sprintf("nats://%s:%d", host, port)
}
