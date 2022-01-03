package base

type (
	Service interface {
		Worker
	}

	BaseService struct {
		*BaseWorker
	}
)

func NewBaseService(name string, log Logger) *BaseService {
	return &BaseService{
		BaseWorker: NewWorker(name, log),
	}
}
