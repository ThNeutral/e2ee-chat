package hub

import (
	"sync"

	"github.com/coder/websocket"
)

type Config struct{}
type Hub struct {
	connsMutex sync.RWMutex
	conns      map[*websocket.Conn]*connectionStatus
}

func New(cfg Config) *Hub {
	hub := &Hub{
		conns: map[*websocket.Conn]*connectionStatus{},
	}

	go hub.pingLoop()

	return hub
}
