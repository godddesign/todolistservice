package base

import "errors"

type (
	CQRSManager struct {
		Service  *Service
		Commands CommandSet
		Queries  QuerySet
	}

	CommandSet map[string]Command
	QuerySet   map[string]Query

	Command interface {
		Name() string
	}

	Query interface {
		Name() string
	}

	BaseCommand struct {
		name string
	}

	BaseQuery struct {
		name string
	}
)

var (
	SampleCommand = BaseCommand{name: "sample-command"}
	SampleQuery   = BaseQuery{name: "sample-query"}
)

func NewCQRSManager(svc *Service) *CQRSManager {
	return &CQRSManager{
		Service:  svc,
		Commands: CommandSet{},
		Queries:  QuerySet{},
	}
}

func (cqrs CQRSManager) AddCommand(c Command) error {
	if c.Name() == "" {
		errors.New("command name is empty")
	}

	cqrs.Commands[c.Name()] = c

	return nil
}

func (cqrs CQRSManager) AddQuery(q Query) error {
	if q.Name() == "" {
		errors.New("query name is empty")
	}

	cqrs.Queries[q.Name()] = q

	return nil
}

func (bc *BaseCommand) Name() string {
	return bc.name
}
