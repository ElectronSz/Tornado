package tornado

type Options struct {
	Port int32
}

type Tornado struct {
	Options
}

type Server interface {
	Start()
}

func (*Tornado) Start() {
	println("Starting new tornado server")
}
func (*Tornado) newServer(opts Options) error {
	println("Stating server", opts.port)
	return nil
}
