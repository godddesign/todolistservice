package app

import (
	"flag"
	"fmt"
	"sync"

	"github.com/godddesign/todo/list/internal/app/adapter/rest"
	"github.com/godddesign/todo/list/internal/app/config"
	"github.com/godddesign/todo/list/internal/app/cqrs/bus/nats"
	"github.com/godddesign/todo/list/internal/app/cqrs/command"
	"github.com/godddesign/todo/list/internal/app/ports/openapi"
	"github.com/godddesign/todo/list/internal/app/service"
	"github.com/godddesign/todo/list/internal/base"
)

type (
	// App description
	App struct {
		*base.App
		Cfg *config.Config

		// Service
		TodoService *service.Todo

		// CQRS
		CQRS *base.CQRSManager

		// Bus
		// NATS
		Bus *nats.BusManager

		RESTServer *rest.Server
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
	app.Log().SetLevel(app.Cfg.Level)
}

// Init app
func (app *App) Init() (err error) {
	// Commands
	app.initCommands()

	// Router
	if app.RESTServer != nil {
		rm := rest.NewRequestManager(app.CQRS, app.Bus, app.Log())
		h := openapi.Handler(rm)
		app.RESTServer.InitRESTRouter(h)
	}

	return nil
}

// Start app
func (app *App) Start() error {
	var errREST error
	var errBus error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		errREST = app.RESTServer.Start(app.Cfg.Server.JSONAPIPort)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		errBus = app.Bus.Start()
		wg.Done()
	}()

	wg.Wait()

	if errREST != nil {
		return fmt.Errorf("cannot start server: %w", errREST)
	}

	if errBus != nil {
		return fmt.Errorf("cannot start server: %w", errBus)
	}

	return fmt.Errorf("cannot start server:\n\t%s\n\t%s\n", errREST.Error(), errBus.Error())
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

func (app *App) LoadConfig() config.Config {
	if app.Cfg == nil {
		app.Cfg = &config.Config{}
	}

	cfg := config.Config{}

	// Server
	flag.IntVar(&cfg.Server.JSONAPIPort, "json-api-port", 8081, "JSON API server port")

	// Mongo
	flag.StringVar(&cfg.Mongo.Host, "mongo-host", "localhost", "Mongo host")
	flag.IntVar(&cfg.Mongo.Port, "mongo-port", 8081, "Mongo port")
	flag.StringVar(&cfg.Mongo.User, "mongo-user", "", "Mongo user")
	flag.StringVar(&cfg.Mongo.Pass, "mongo-pass", "", "Mongo pass")
	flag.StringVar(&cfg.Mongo.Database, "mongo-database", "", "Mongo database")
	flag.IntVar(&cfg.Mongo.MaxRetries, "mongo-max-reties", 10, "Mongo port")

	// NATS
	flag.StringVar(&cfg.NATS.Host, "nats-host", "0.0.0.0", "NATS host")
	flag.IntVar(&cfg.NATS.Port, "nats-port", 4222, "NATS port")

	app.Cfg = &cfg

	return cfg
}
