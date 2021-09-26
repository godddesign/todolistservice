package base

type (
	Worker interface {
		Name() string
		Init() error
		Start() error
	}
)

type (
	BaseWorker struct {
		name     string
		didInit  bool
		didStart bool
		*Tracer
	}
)

func NewWorker(name string, tracingLevel string) *BaseWorker {
	name = genName(name, "worker")

	return &BaseWorker{
		name:   name,
		Tracer: NewTracer(tracingLevel),
	}
}

func (bw BaseWorker) Name() string {
	return bw.name
}

func (bw BaseWorker) SetName(name string) {
	bw.name = name
}

func (bw BaseWorker) Init() error {
	return bw.Init()
}

func (bw BaseWorker) Start() error {
	return bw.Start()
}
