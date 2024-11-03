package tornado

type Options struct {
	port int32
}

type Tornado struct {
	opts Options
}

type Server interface {
	Start(*Tornado)
}

func Start() {
	println("Starting new tornado server")
}
func (*Tornado) newServer(opts Options) error {
	println("Stating server")
	return nil
}
