package server

import "github.com/coder/websocket"

type Hub interface {
	AddConnection(*websocket.Conn) error
}
