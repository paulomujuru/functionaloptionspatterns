package server

import (
	"log"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
}

func New(host string, port int, timeout time.Duration) *Server {
	s := &Server{
		host:    host,
		port:    port,
		timeout: timeout,
	}
	return s
}

func (s *Server) Start() {
	log.Printf("Server starting %#v/n", s)
}

func (s *Server) Stop() {
	log.Printf("Servers stopped %#v/n", s)
}
