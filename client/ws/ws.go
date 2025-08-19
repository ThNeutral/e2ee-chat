package ws

import (
	"context"
	"net/url"
	"time"

	"github.com/coder/websocket"
)

type Config struct {
	WSEndpoint *url.URL
}
type Websocket struct {
	wsEndpoint *url.URL

	conn            *websocket.Conn
	lastMessageTime time.Time
}

func New(cfg Config) *Websocket {
	return &Websocket{
		wsEndpoint: cfg.WSEndpoint,
	}
}

func (ws *Websocket) Connect() error {
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

	ws.conn = conn
	go ws.reader()
	go ws.pingLoop()

	return nil
}
