package server

import (
	"chat/shared"
)

type ServerConfig struct {
	Hub Hub
}

type Server struct {
	hub Hub
}

func New(cfg ServerConfig) (*Server, error) {
	eb := shared.NewErrorBuilder().Msg("failed to initialize server")

	if cfg.Hub == nil {
		return nil, eb.Causef("hub was not passed").Err()
	}

	s := &Server{
		hub: cfg.Hub,
	}

	return s, nil
}
