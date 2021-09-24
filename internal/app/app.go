package app

import (
	"sync"

	"github.com/adrianpk/cirrus"
	"github.com/adrianpk/cirrustodo/internal/app/adapter/jsonapi"
	"github.com/adrianpk/cirrustodo/internal/app/cqrs/command"
	"github.com/adrianpk/cirrustodo/internal/app/ports/openapi"
	"github.com/adrianpk/cirrustodo/internal/app/service"
)

type (
	// App description
	App struct {
		*cirrus.App

		Config

		JSONAPIServer *jsonapi.Server
		//WebServer     *web.Server
		//GRPCServer    *grpc.Server
	}
)

type (
	Command struct {
		name string
	}
)

// NewApp initializes new App worker instance
func NewApp(name string, ts *service.Todo, cfg *Config) (*App, error) {
	app := App{
		App: cirrus.NewApp(name),
	}

	// Server
	jas, err := jsonapi.NewServer("json-api-server", ts, jsonapi.Config{
		TracingLevel: cfg.Tracing.Level,
	})
	if err != nil {
		return nil, err
	}

	// Server
	app.JSONAPIServer = jas

	// Router
	h := openapi.Handler(ts)
	jas.InitJSONAPIRouter(h)

	return &app, nil
}

// Init app
func (app *App) Init() (err error) {
	return app.initCommands()
}

// Start app
func (app *App) Start() error {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		app.JSONAPIServer.Start(app.Config.Server.JSONAPIPort)
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func (app *App) Stop() {
	// TODO: Gracefully stop the app
}

func (app *App) initCommands() (err error) {
	server := app.JSONAPIServer
	app.AddCommand(&cirrus.SampleCommand)
	app.AddCommand(command.NewCreateListCommand(server.Todo))
	//app.AddCommand(command.NewAddItemCommand(app.TodoService))
	//app.AddCommand(command.NewGetItemCommand(app.TodoService))
	//app.AddCommand(command.NewUpdateItemCommand(app.TodoService))
	//app.AddCommand(command.NewDeleteItemCommand(app.TodoService))
	//app.AddCommand(command.NewDeleteListCommand(app.TodoService))
	return err
}
