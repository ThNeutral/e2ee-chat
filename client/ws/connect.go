package ws

import (
	"context"
	"fmt"
	"time"

	"github.com/coder/websocket"
)

func (ws *Websocket) Connect() error {
	if ws.IsConnected() {
		return fmt.Errorf("already connected")
	}

	conn, _, err := websocket.Dial(
		context.Background(),
		ws.wsEndpoint.String(),
		&websocket.DialOptions{
			OnPingReceived: func(ctx context.Context, payload []byte) bool {
				ws.lastMessageTime = time.Now()
				return false
			},
			OnPongReceived: func(ctx context.Context, payload []byte) {
				ws.lastMessageTime = time.Now()
			},
		},
	)
	if err != nil {
		return err
	}

	if ws.onConnectHandler != nil {
		ws.onConnectHandler()
	}

	ws.conn = conn
	go ws.reader()
	go ws.pingLoop()

	return nil
}

func (ws *Websocket) Disconnect(reason string) error {
	if !ws.IsConnected() {
		return fmt.Errorf("not connected")
	}

	err := ws.conn.Close(websocket.StatusGoingAway, reason)
	if err != nil {
		return err
	}

	if ws.onDisconnectHandler != nil {
		ws.onDisconnectHandler()
	}

	ws.conn = nil

	return nil
}

func (ws *Websocket) IsConnected() bool {
	return ws.conn != nil
}
