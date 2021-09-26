package base

import "errors"

type (
	CommandReg map[string]Command

	Command interface {
		Name() string
	}

	BaseCommand struct {
		name string
	}
)

var (
	SampleCommand = BaseCommand{name: "sample-command"}
)

func newCommandReg() CommandReg {
	return CommandReg{}
}

func (cr CommandReg) Add(c Command) error {
	if c.Name() == "" {
		errors.New("command name is empty")
	}

	cr[c.Name()] = c

	return nil
}

func (bc *BaseCommand) Name() string {
	return bc.name
}
