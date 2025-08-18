package hub

import (
	"context"
	"log"
	"time"

	"github.com/coder/websocket"
)

func (h *Hub) HandleSignal() {
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	for range timer.C {
		for conn := range h.conns {
			go func(c *websocket.Conn) {
				err := c.Write(
					context.Background(),
					websocket.MessageText,
					[]byte("this message is repeated every 5 seconds"),
				)
				if err != nil {
					log.Printf("failed to send message: %v\n", err)
				}
			}(conn)
		}
	}
}
