package hub

import (
	"context"
	"time"

	"github.com/coder/websocket"
)

const (
	PING_DELAY = 10 * time.Second
)

func (h *Hub) pingLoop() {
	ticker := time.NewTicker(PING_DELAY / 2)
	defer ticker.Stop()

	for range ticker.C {
		h.connsMutex.RLock()
		for conn, connStatus := range h.conns {
			go h.handlePing(conn, connStatus)
		}
		h.connsMutex.RUnlock()
	}
}

func (h *Hub) handlePing(c *websocket.Conn, status *connectionStatus) {
	if time.Now().After(status.LastMessageTime.Add(PING_DELAY * 3)) {
		h.finalize(c)
		return
	}

	if time.Now().Before(status.LastMessageTime.Add(PING_DELAY)) {
		return
	}

	err := c.Ping(context.TODO())
	if err != nil {
		h.finalize(c)
		return
	}
}
