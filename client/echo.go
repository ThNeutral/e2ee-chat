package client

import (
	"context"
)

type EchoRequest struct {
	Value string `json:"value"`
}

type EchoResponse struct {
	Value string `json:"value"`
}

func (c *Client) Echo(ctx context.Context, payload string) (string, error) {
	return c.echo.Echo(ctx, payload)
}
