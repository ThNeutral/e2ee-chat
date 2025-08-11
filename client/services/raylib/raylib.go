package raylib

import (
	"chat/client/entities"
	"chat/shared"
)

type Config struct {
	WindowConfig entities.WindowConfig
}
type Raylib struct {
	windowConfig entities.WindowConfig

	running bool
	close   chan bool

	widgets []*entities.RectangleWidget

	focused *entities.RectangleWidget
}

func New(cfg Config) (*Raylib, error) {
	eb := shared.NewErrorBuilder().Msg("failed to create raylib")

	if cfg.WindowConfig.Width == 0 {
		return nil, eb.Causef("width not passed").Err()
	}

	if cfg.WindowConfig.Width == 0 {
		return nil, eb.Causef("height not passed").Err()
	}

	if cfg.WindowConfig.Title == "" {
		return nil, eb.Causef("title not passed").Err()
	}

	return &Raylib{
		windowConfig: cfg.WindowConfig,

		running: false,
		close:   make(chan bool),

		widgets: []*entities.RectangleWidget{},
	}, nil
}
