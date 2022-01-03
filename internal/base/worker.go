package base

type (
	Worker interface {
		Name() string
		Init() error
		Start() error
		Log() Logger
	}
)

type (
	BaseWorker struct {
		name     string
		didInit  bool
		didStart bool
		log      Logger
	}
)

func NewWorker(name string, log Logger) *BaseWorker {
	name = GenName(name, "worker")

	return &BaseWorker{
		name: name,
		log:  log,
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

func (bw BaseWorker) Log() Logger {
	return bw.log
}
