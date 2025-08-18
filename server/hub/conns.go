package hub

import "github.com/coder/websocket"

func (h *Hub) AddConnection(conn *websocket.Conn) error {
	h.conns[conn] = true

	return nil
}
