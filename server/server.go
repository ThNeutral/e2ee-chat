package server

import (
	"chat/server/middlewares"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Config struct {
	Port int
	Hub  Hub
}
type Server struct {
	port    int
	handler http.Handler
	hub     Hub
}

func New(cfg Config) *Server {
	router := chi.NewRouter()

	router.Use(middlewares.Logger)
	router.Use(middlewares.ResponseWriterWrapper)

	server := &Server{
		port:    cfg.Port,
		handler: router,
		hub:     cfg.Hub,
	}

	router.Get("/chat", server.HandleConnectChat)

	return server
}

func (s *Server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%v", s.port), s.handler)
}
