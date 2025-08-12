package runner

import "chat/raylib/entities"

type Config struct {
	WindowConfig entities.WindowConfig
}

type Runner struct {
	windowConfig entities.WindowConfig

	widgets []*entities.RectangleWidget
	focused *entities.RectangleWidget

	running bool
}

func New(cfg Config) *Runner {
	return &Runner{
		windowConfig: cfg.WindowConfig,
		widgets:      []*entities.RectangleWidget{},
		focused:      nil,
		running:      false,
	}
}
