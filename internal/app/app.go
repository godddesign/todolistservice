package app

import (
	"fmt"
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
		Config *Config

		// Service
		TodoService *service.Todo

		// CQRS
		CQRS *base.CQRSManager

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

func NewApp(name, version string, log base.Logger) *App {
	return &App{
		App:  base.NewApp(name, version, log),
		CQRS: base.NewCQRSManager(),
	}
}

func (app *App) SetLogLevel(level string) {
	app.Log().SetLevel(app.Config.Level)
}

// Init app
func (app *App) Init() error {
	// Server
	jas, err := jsonapi.NewServer("json-api-server", &jsonapi.Config{}, app.Log())
	if err != nil {
		return err
	}

	// Server
	app.JSONAPIServer = jas

	// Commands
	app.initCommands()

	// Router
	rm := jsonapi.NewRequestManager(app.CQRS, app.Log())
	h := openapi.Handler(rm)

	jas.InitJSONAPIRouter(h)

	return nil
}

// Start app
func (app *App) Start() error {
	var err error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		err = app.JSONAPIServer.Start(app.Config.Server.JSONAPIPort)
		wg.Done()
	}()

	wg.Wait()
	return err
}

func (app *App) InitAndStart() error {
	err := app.Init()
	if err != nil {
		return fmt.Errorf("%s init error: %w", app.Name(), err)
	}

	err = app.Start()
	if err != nil {
		return fmt.Errorf("%s start error: %w", app.Name(), err)
	}

	return nil
}

func (app *App) Stop() {
	// TODO: Gracefully stop the app
}

func (app *App) initCommands() {
	log := app.Log()
	app.AddCommand(&base.SampleCommand) // TODO: Remove
	app.AddCommand(command.NewCreateListCommand(app.TodoService, log))
	//app.AddCommand(command.NewAddItemCommand(app.TodoService))
	//app.AddCommand(command.NewGetItemCommand(app.TodoService))
	//app.AddCommand(command.NewUpdateItemCommand(app.TodoService))
	//app.AddCommand(command.NewDeleteItemCommand(app.TodoService))
	//app.AddCommand(command.NewDeleteListCommand(app.TodoService))
}

func (app *App) AddCommand(command base.Command) {
	app.CQRS.AddCommand(command)
}

func (app *App) AddQuery(query base.Query) {
	app.CQRS.AddQuery(query)
}
