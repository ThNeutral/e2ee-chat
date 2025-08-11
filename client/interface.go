package client

import (
	"chat/client/entities"
	"context"
)

type GUI interface {
	Init() error
	Close() error
	Run() error

	AddRectangleWidget(x int32, y int32, width int32, height int32) *entities.RectangleWidget
}

type Echo interface {
	Echo(ctx context.Context, payload string) (string, error)
}
