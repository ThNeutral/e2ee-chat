package raylib

import (
	"chat/raylib/entities"
	"context"
)

type Runner interface {
	Init() error
	Close() error
	Run() error
}

type GUI interface {
	SetRootComponent(entities.RootComponent)
}

type Echo interface {
	Echo(ctx context.Context, payload string) (string, error)
}
