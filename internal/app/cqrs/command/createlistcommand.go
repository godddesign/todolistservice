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

func NewCreateListCommand(todoService *service.Todo, log base.Logger) *CreateListCommand {
	if todoService == nil {
		panic("nil Todo service")
	}

	return &CreateListCommand{
		BaseWorker:  base.NewWorker("create-list-command", log),
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

func (c *CreateListCommand) handle(ctx context.Context, data interface{}) error {
	var err error

	defer func() {
		if err != nil {
			c.Log().Errorf("command %s error: %w", c.Name(), err)
		}
	}()

	switch d := data.(type) {
	case CreateListCommandData:
		c.Log().Debugf("Processing %s with %+v", c.Name(), d)

		err = c.todoService.CreateList(d.Name, d.Description)
		if err != nil {
			return fmt.Errorf("%s error: %w", c.Name(), err)
		}

	default:
		return errors.New("wrong command data")
	}

	return nil
}
