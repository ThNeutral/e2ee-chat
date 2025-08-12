package client

import (
	"chat/shared"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Config struct {
	ServerAddr net.Addr
	HTTPClient *http.Client

	DefaultTimeout time.Duration

	Echo Echo
}

type Client struct {
	serverAddr net.Addr
	httpClient *http.Client

	defaultTimeout time.Duration

	echo Echo
}

func New(cfg Config) (*Client, error) {
	eb := shared.B().Msg("failed to initialize client")

	if cfg.HTTPClient == nil {
		return nil, eb.Causef("http client not passed").Err()
	}

	if cfg.ServerAddr == nil {
		return nil, eb.Causef("server addr not passed").Err()
	}

	if cfg.Echo == nil {
		return nil, eb.Causef("echo not passed").Err()
	}

	if cfg.DefaultTimeout == 0 {
		cfg.DefaultTimeout = 15 * time.Second
	}

	cl := &Client{
		serverAddr: cfg.ServerAddr,
		httpClient: cfg.HTTPClient,

		defaultTimeout: cfg.DefaultTimeout,
		echo:           cfg.Echo,
	}

	return cl, nil
}

func (c *Client) getServerURL(endpoint string) string {
	return fmt.Sprintf("http://%v%v", c.serverAddr.String(), endpoint)
}
