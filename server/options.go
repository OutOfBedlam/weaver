package server

func WithData(path string) Option {
	return func(s *Server) {
		s.dataRoot = path
	}
}
