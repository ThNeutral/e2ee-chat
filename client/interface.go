package client

import "chat/client/raylib"

type GUI interface {
	Init() error
	Close() error
	Run() error

	SetRootComponent(raylib.Component)
}

type Websocket interface {
	Connect() error
}
