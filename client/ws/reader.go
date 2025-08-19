package ws

import (
	"context"
	"fmt"
)

func (ws *Websocket) reader() {
	for {
		if !ws.IsConnected() {
			return
		}

		messageType, payload, err := ws.conn.Read(context.TODO())
		if err != nil {
			fmt.Println(err)
			ws.finalize()
			continue
		}

		fmt.Printf("Received message of type %v: %v", messageType, string(payload))
	}
}
