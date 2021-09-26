package app

import (
	"sync"

	"github.com/adrianpk/godddtodo/internal/app/adapter/jsonapi"
	"github.com/adrianpk/godddtodo/internal/app/cqrs/command"
	"github.com/adrianpk/godddtodo/internal/app/ports/openapi"
	"github.com/adrianpk/godddtodo/internal/app/service"
	"github.com/adrianpk/godddtodo/internal/base"
)

type (
	// App description
	App struct {
		*base.App
		Config

		// Service
		TodoService *service.Todo

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
func NewApp(name string, svc *service.Todo, cfg *Config) (*App, error) {
	app := App{
		App:         base.NewApp(name),
		TodoService: svc,
	}

	// Server
	jas, err := jsonapi.NewServer("json-api-server", svc, jsonapi.Config{
		TracingLevel: cfg.Tracing.Level,
	})
	if err != nil {
		return nil, err
	}

	// Server
	app.JSONAPIServer = jas

	// Router
	h := openapi.Handler(svc)
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
	app.AddCommand(&base.SampleCommand)
	app.AddCommand(command.NewCreateListCommand(server.Todo))
	//app.AddCommand(command.NewAddItemCommand(server.TodoService))
	//app.AddCommand(command.NewGetItemCommand(server.TodoService))
	//app.AddCommand(command.NewUpdateItemCommand(server.TodoService))
	//app.AddCommand(command.NewDeleteItemCommand(server.TodoService))
	//app.AddCommand(command.NewDeleteListCommand(server.TodoService))
	return err
}
