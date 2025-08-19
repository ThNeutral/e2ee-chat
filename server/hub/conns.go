package hub

import (
	"context"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

func (h *Hub) Accept(w http.ResponseWriter, r *http.Request) error {
	var conn *websocket.Conn
	var err error
	conn, err = websocket.Accept(w, r, &websocket.AcceptOptions{
		OnPingReceived: func(ctx context.Context, payload []byte) bool {
			h.updateLastMessageTime(conn)
			return true
		},
		OnPongReceived: func(ctx context.Context, payload []byte) {
			h.updateLastMessageTime(conn)
		},
	})
	if err != nil {
		return err
	}

	status := &connectionStatus{
		LastMessageTime: time.Now(),
	}

	h.connsMutex.Lock()
	h.conns[conn] = status
	h.connsMutex.Unlock()

	go h.readLoop(conn)

	return nil
}

func (h *Hub) finalize(conn *websocket.Conn) error {
	h.connsMutex.Lock()
	delete(h.conns, conn)
	h.connsMutex.Unlock()

	conn.CloseNow()

	return nil
}
