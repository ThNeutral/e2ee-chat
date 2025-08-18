package server

import (
	"context"
	"fmt"
)

type EchoRequest struct {
	Value string `json:"value"`
}

func (r EchoRequest) Validate() error {
	if r.Value == "" {
		return fmt.Errorf("value cannot be empty")
	}

	return nil
}

type EchoResponse struct {
	Value string `json:"value"`
}

func (s *Server) HandleEcho(ctx context.Context, req EchoRequest) (any, error) {
	return EchoResponse{
		Value: req.Value,
	}, nil
}
