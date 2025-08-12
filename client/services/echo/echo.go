package echo

import (
	"chat/shared"
	"context"
)

type Config struct {
	EchoRepository EchoRepository
}
type Echo struct {
	echoRepository EchoRepository
}

type EchoRepository interface {
	Echo(ctx context.Context, payload string) (string, error)
}

func New(cfg Config) (*Echo, error) {
	eb := shared.B().Msg("failed to initialize echo service")

	if cfg.EchoRepository == nil {
		return nil, eb.Causef("echo repository not passed").Err()
	}

	return &Echo{
		echoRepository: cfg.EchoRepository,
	}, nil
}

func (r *Echo) Echo(ctx context.Context, payload string) (string, error) {
	return r.echoRepository.Echo(ctx, payload)
}
