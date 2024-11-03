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

func Start(tornado *Tornado) {
	println("Starting new tornado server")
}
func newServer(opts Options) *Tornado {
	return &Tornado{
		opts,
	}
}