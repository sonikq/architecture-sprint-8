package http

import (
	"context"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(address string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + strings.Split(address, ":")[1],
			Handler:        handler,
			ReadTimeout:    150 * time.Hour,
			WriteTimeout:   150 * time.Hour,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
