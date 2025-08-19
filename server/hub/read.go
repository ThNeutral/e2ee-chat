package hub

import (
	"context"
	"fmt"
	"log"

	"github.com/coder/websocket"
)

func (h *Hub) readLoop(c *websocket.Conn) {
	for {
		messageType, payload, err := c.Read(context.TODO())
		if err != nil {
			log.Printf("failed to read message: %v", err)
			h.finalize(c)
			return
		}

		fmt.Printf("type: %v; paylaod: %v\n", messageType, string(payload))
	}
}
