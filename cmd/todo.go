package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/adrianpk/godddtodo/internal/app/cqrs/bus/nats"
	"github.com/adrianpk/godddtodo/internal/app/repo/mongo"
	"github.com/adrianpk/godddtodo/internal/base"

	"github.com/adrianpk/godddtodo/internal/app"
	"github.com/adrianpk/godddtodo/internal/app/service"
	db "github.com/adrianpk/godddtodo/internal/base/db/mongo"
)

type contextKey string

const (
	name    = "todo"
	version = "v0.0.1"
)

var (
	a   *app.App
	log base.Logger
)

func main() {
	log = base.NewLogger("debug", true)

	// App
	a := app.NewApp(name, version, log)
	cfg := a.LoadConfig()

	// Context
	ctx, cancel := context.WithCancel(context.Background())
	initExitMonitor(ctx, cancel)

	// Database
	mgo := db.NewMongoClient("mongo-client", db.Config{
		Host:       cfg.Mongo.Host,
		Port:       cfg.Mongo.Port,
		User:       cfg.Mongo.User,
		Pass:       cfg.Mongo.Pass,
		Database:   cfg.Mongo.Database,
		MaxRetries: cfg.Mongo.MaxRetriesUInt64(),
	}, log)

	// Repo
	lrr := mongo.NewListRead("list-read-repo", mgo, mongo.Config{}, log)

	lwr := mongo.NewListWrite("list-write-repo", mgo, mongo.Config{}, log)

	// Service
	ts, err := service.NewTodo("todo-app-service", lrr, lwr, service.Config{}, log)

	if err != nil {
		exit(err)
	}

	a.TodoService = &ts

	// Bus (TODO: Get config values from flags / environment)
	a.NATS = nats.NewNATSClient("nats-client", nats.Config{
		Host: "localhost",
		Port: 4222,
	}, log)

	// Init & Start
	err = a.InitAndStart()
	if err != nil {
		exit(err)
	}

	log.Errorf("%s stopped: %s (%s)", a.Name(), a.Version(), err)
}

func exit(err error) {
	log.Fatal(err)
}

func initExitMonitor(ctx context.Context, cancel context.CancelFunc) {
	go checkSigterm(cancel)
	go checkCancel(ctx)
}

func checkSigterm(cancel context.CancelFunc) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	cancel()
}

func checkCancel(ctx context.Context) {
	<-ctx.Done()
	a.Stop()
	os.Exit(1)
}
