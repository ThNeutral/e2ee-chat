package client

import (
	"chat/shared"
	"fmt"
	"net"
	"net/http"
)

type Config struct {
	ServerAddr net.Addr
	HTTPClient *http.Client

	GUI GUI
}

type Client struct {
	serverAddr net.Addr
	httpClient *http.Client

	gui GUI
}

func New(cfg Config) (*Client, error) {
	eb := shared.NewErrorBuilder().Msg("failed to initialize client")

	if cfg.HTTPClient == nil {
		return nil, eb.Causef("http client not passed").Err()
	}

	if cfg.ServerAddr == nil {
		return nil, eb.Causef("server addr not passed").Err()
	}

	if cfg.GUI == nil {
		return nil, eb.Causef("gui not passed").Err()
	}

	cl := &Client{
		serverAddr: cfg.ServerAddr,
		httpClient: cfg.HTTPClient,

		gui: cfg.GUI,
	}

	return cl, nil
}

func (c *Client) Run() error {
	eb := shared.NewErrorBuilder().Msg("failed to run client")

	err := c.gui.Init()
	if err != nil {
		return eb.Cause(err).Err()
	}
	defer shared.CloseWithEB(c.gui, eb)

	err = c.gui.Run()
	if err != nil {
		return eb.Cause(err).Err()
	}

	return nil
}

func (c *Client) getServerURL(endpoint string) string {
	return fmt.Sprintf("http://%v%v", c.serverAddr.String(), endpoint)
}
