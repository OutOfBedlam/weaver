package server

type Server struct {
	dataRoot string
}

type Option func(*Server)

func New(opts ...Option) *Server {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (svr *Server) Start() error {
	return nil
}

func (svr *Server) Stop() {

}
