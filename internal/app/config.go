package app

import "flag"

type (
	Config struct {
		Server
		Mongo
		Logging
	}

	Server struct {
		JSONAPIPort int
	}

	Mongo struct {
		Host       string
		Port       int
		User       string
		Pass       string
		Database   string
		MaxRetries int
	}

	Logging struct {
		Level string
	}
)

func (app *App) LoadConfig() *Config {
	if app.Config == nil {
		app.Config = &Config{}
	}

	cfg := app.Config

	// Server
	flag.IntVar(&cfg.Server.JSONAPIPort, "json-api-port", 8081, "JSON API server port")

	// Mongo
	flag.StringVar(&cfg.Mongo.Host, "mongo-host", "localhost", "Mongo host")
	flag.IntVar(&cfg.Mongo.Port, "mongo-port", 8081, "Mongo port")
	flag.StringVar(&cfg.Mongo.User, "mongo-user", "", "Mongo user")
	flag.StringVar(&cfg.Mongo.Pass, "mongo-pass", "", "Mongo pass")
	flag.StringVar(&cfg.Mongo.Database, "mongo-database", "", "Mongo database")
	flag.IntVar(&cfg.Mongo.MaxRetries, "mongo-max-reties", 10, "Mongo port")

	return cfg
}

func (m *Mongo) MaxRetriesUInt64() uint64 {
	return uint64(m.MaxRetries)
}
