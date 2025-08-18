package hub

import "github.com/coder/websocket"

type Config struct{}
type Hub struct {
	conns map[*websocket.Conn]bool
}

func New(cfg Config) *Hub {
	hub := &Hub{
		conns: map[*websocket.Conn]bool{},
	}

	go hub.HandleSignal()

	return hub
}
