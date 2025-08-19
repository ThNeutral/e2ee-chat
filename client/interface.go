package client

import (
	"chat/client/entities"
)

type GUI interface {
	Init() error
	Close() error
	Run() error

	Root() entities.Component
}

type Websocket interface {
	Connect() error
	Disconnect(reason string) error
	IsConnected() bool

	SetOnConnectHandler(entities.OnConnectHandler)
	SetOnDisconnectHandler(entities.OnDisconnectHandler)
}
