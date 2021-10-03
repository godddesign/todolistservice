package app

import (
	"sync"

	"github.com/adrianpk/godddtodo/internal/app/adapter/driving/jsonapi"
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
func NewApp(name string, svc *service.Todo, cfg *Config) (app *App, err error) {
	app = &App{
		App:         base.NewApp(name),
		TodoService: svc,
	}

	err = app.initCommands()
	if err != nil {
		return nil, err
	}

	// Server
	jas, err := jsonapi.NewServer("json-api-server", jsonapi.Config{
		TracingLevel: cfg.Tracing.Level,
	})
	if err != nil {
		return nil, err
	}

	// Server
	app.JSONAPIServer = jas

	// Router
	rm, err := jsonapi.NewRequestManager(app.CQRS)
	if err != nil {
		return nil, err
	}

	h := openapi.Handler(rm)
	jas.InitJSONAPIRouter(h)

	return app, nil
}

// Init app
func (app *App) Init() (err error) {
	// TODO
	return nil
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
	tl := app.JSONAPIServer.TracingLevel
	app.AddCommand(&base.SampleCommand) // TODO: Remove
	app.AddCommand(command.NewCreateListCommand(app.TodoService, tl))
	//app.AddCommand(command.NewAddItemCommand(app.TodoService))
	//app.AddCommand(command.NewGetItemCommand(app.TodoService))
	//app.AddCommand(command.NewUpdateItemCommand(app.TodoService))
	//app.AddCommand(command.NewDeleteItemCommand(app.TodoService))
	//app.AddCommand(command.NewDeleteListCommand(app.TodoService))

	// TODO: Implement error return for command creation fail

	return nil
}
