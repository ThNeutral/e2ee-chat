package ws

import (
	"chat/client/entities"
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

	onConnectHandler    entities.OnConnectHandler
	onDisconnectHandler entities.OnDisconnectHandler
}

func New(cfg Config) *Websocket {
	return &Websocket{
		wsEndpoint: cfg.WSEndpoint,
	}
}
