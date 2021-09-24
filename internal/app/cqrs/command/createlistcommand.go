package command

import (
	"context"

	"github.com/adrianpk/cirrustodo/internal/app/service"
)

type (
	CreateListCommandData struct {
		Name        string
		Description string
	}

	CreateListCommand struct {
		name        string
		todoService *service.Todo
	}
)

func NewCreateListCommand(todoService *service.Todo) CreateListCommand {
	if todoService == nil {
		panic("nil Todo service")
	}

	return CreateListCommand{
		name:        "create-list",
		todoService: todoService,
	}
}

func (c CreateListCommand) Name() string {
	return c.name
}

func (c CreateListCommand) Handle(ctx context.Context, data CreateListCommandData) (err error) {
	defer func() {
		// TODO: Trace command error: name, data, error
	}()

	// TODO: Use handler associated todoService to create list

	return nil
}
