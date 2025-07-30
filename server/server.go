package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	handler    http.Handler
	serverPort int
}

func New(serverPort int) *Server {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	s := &Server{
		handler:    r,
		serverPort: serverPort,
	}

	r.Post("/echo", s.handleEcho)

	return s
}

func (s *Server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%v", s.serverPort), s.handler)
}
