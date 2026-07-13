package server

import "net/http"

type Server struct {
	mux      *http.ServeMux
	addr     string
	handlers map[string]http.HandlerFunc
}

func New(addr string) *Server {
	mux := http.NewServeMux()
	return &Server{
		mux:      mux,
		addr:     addr,
		handlers: make(map[string]http.HandlerFunc),
	}
}

func (s *Server) AddHandler(path string, handler http.HandlerFunc) {
	s.handlers[path] = handler
}

func (s *Server) Start() error {
	for path, handler := range s.handlers {
		s.mux.HandleFunc(path, handler)
	}
	return http.ListenAndServe(s.addr, s.mux)
}
