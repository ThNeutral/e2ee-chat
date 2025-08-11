package client

import "chat/client/entities"

type GUI interface {
	Init() error
	Close() error
	Run() error

	AddRectangleWidget(x int32, y int32, width int32, height int32) *entities.RectangleWidget
}
