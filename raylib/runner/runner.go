package runner

import (
	"chat/raylib/entities"
)

type Config struct {
	WindowConfig entities.WindowConfig
}

type Runner struct {
	windowConfig entities.WindowConfig

	root    *entities.RectangleWidget
	focused *entities.RectangleWidget

	running bool
}

func New(cfg Config) *Runner {
	rect := cfg.WindowConfig.ToRectangle()

	root := getRectangleWidget(rect.X, rect.Y, rect.Width, rect.Height)
	root.BackgroundColor = cfg.WindowConfig.BackgroundColor

	return &Runner{
		windowConfig: cfg.WindowConfig,

		root:    root,
		focused: nil,

		running: false,
	}
}
