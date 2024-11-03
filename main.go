package tornado

type Options struct {
	Port int32
}

type Tornado struct {
	Options
}

type Server interface {
	Start()
	NewServer(opts Options) error
}

func (*Tornado) Start() {
	println("Starting new tornado server")
}
func (*Tornado) NewServer(opts Options) error {
	println("Stating server", opts.Port)
	return nil
}
