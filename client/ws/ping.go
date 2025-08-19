package ws

import (
	"context"
	"log"
	"time"
)

const (
	PING_DELAY = 10 * time.Second
)

func (ws *Websocket) pingLoop() {
	ticker := time.NewTicker(PING_DELAY / 2)
	defer ticker.Stop()

	for range ticker.C {
		if !ws.IsConnected() {
			return
		}

		if time.Now().Before(ws.lastMessageTime.Add(PING_DELAY)) {
			return
		}

		err := ws.conn.Ping(context.TODO())
		if err != nil {
			log.Printf("failed to ping: %v", err)
			ws.finalize()
			return
		}
	}
}
