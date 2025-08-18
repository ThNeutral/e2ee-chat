package ws

import (
	"context"
	"net/url"

	"github.com/coder/websocket"
)

type Config struct {
	WSEndpoint url.URL
}
type Websocket struct {
	wsEndpoint url.URL
	conn       *websocket.Conn
}

func New(cfg Config) *Websocket {
	return &Websocket{
		wsEndpoint: cfg.WSEndpoint,
	}
}

func (ws *Websocket) Connect() error {
	conn, _, err := websocket.Dial(context.Background(), ws.wsEndpoint.String(), nil)
	if err != nil {
		return err
	}

	ws.conn = conn
	go ws.reader()

	return nil
}
