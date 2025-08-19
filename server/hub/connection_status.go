package hub

import (
	"time"

	"github.com/coder/websocket"
)

type connectionStatus struct {
	LastMessageTime time.Time
}

func (h *Hub) updateLastMessageTime(c *websocket.Conn) {
	h.connsMutex.Lock()
	h.conns[c].LastMessageTime = time.Now()
	h.connsMutex.Unlock()
}
