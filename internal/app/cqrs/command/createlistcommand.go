package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/adrianpk/godddtodo/internal/app/service"
	"github.com/adrianpk/godddtodo/internal/base"
)

type (
	CreateListCommandData struct {
		Name        string
		Description string
	}

	CreateListCommand struct {
		*base.BaseWorker
		*base.BaseCommand
		todoService *service.Todo
	}
)

func NewCreateListCommand(todoService *service.Todo, tracinglevel string) *CreateListCommand {
	if todoService == nil {
		panic("nil Todo service")
	}

	return &CreateListCommand{
		BaseWorker:  base.NewWorker("create-list-command", tracinglevel),
		BaseCommand: base.NewBaseCommand("create-list"),
		todoService: todoService,
	}
}

func (c *CreateListCommand) Name() string {
	return c.BaseCommand.Name()
}

func (c *CreateListCommand) HandleFunc() (f func(ctx context.Context, data interface{}) error) {
	return c.handle
}

func (c *CreateListCommand) handle(ctx context.Context, data interface{}) (err error) {
	defer func() {
		// TODO: Trace command error: name, data, error
		if err != nil {
			c.SendDebugf("error handling %s: %v", c.Name(), err)
		}
	}()

	switch d := data.(type) {
	case CreateListCommandData:
		// TODO: Use c.todoService to do something meaningful with data
		fmt.Printf("Procesing %+v: ", d)

	default:
		// TODO: Write error response
		return errors.New("wrong command data")
	}

	return nil
}
