package server

import (
	"net/http"
	"time"
)

type server struct {
	httpServer *http.Server
}

func NewServer() *server {
	return &server{}
}

func (s *server) RunServer(addr string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ":" + addr,
		Handler:      handler,
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}
