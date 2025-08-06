package services

import "github.com/gorilla/websocket"

type HubConfig struct {
	Upgrader websocket.Upgrader
}
type Hub struct {
	upgrader websocket.Upgrader
	conns    map[*websocket.Conn]bool
}

func NewHub(cfg HubConfig) *Hub {
	return &Hub{
		upgrader: cfg.Upgrader,
		conns:    map[*websocket.Conn]bool{},
	}
}
