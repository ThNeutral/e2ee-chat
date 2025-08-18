package ws

import (
	"context"
	"fmt"
)

func (ws *Websocket) reader() {
	for {
		messageType, payload, err := ws.conn.Read(context.TODO())
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Received message of type %v: %v", messageType, string(payload))
	}
}
