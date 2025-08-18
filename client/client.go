package client

import "fmt"

type Config struct {
	GUI GUI
}

type Client struct {
	gui GUI
}

func New(cfg Config) *Client {
	return &Client{
		gui: cfg.GUI,
	}
}

func (c *Client) Run() error {
	err := c.gui.Init()
	if err != nil {
		return fmt.Errorf("failed to init gui: %w", err)
	}

	err = c.gui.Run()
	if err != nil {
		return fmt.Errorf("failed to run gui: %w", err)
	}

	err = c.gui.Close()
	if err != nil {
		return fmt.Errorf("failed to close gui: %w", err)
	}

	return nil
}
