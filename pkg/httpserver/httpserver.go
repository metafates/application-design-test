// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func (s *Server) Start() {
	s.notify <- s.server.ListenAndServe()
	close(s.notify)
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
