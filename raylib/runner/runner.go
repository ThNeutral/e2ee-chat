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
	r := &Runner{
		windowConfig: cfg.WindowConfig,

		root:    nil,
		focused: nil,

		running: false,
	}

	return r
}

func (r *Runner) SetRootComponent(root entities.RootComponent) {
	r.root = root(entities.RootComponentProps{
		WindowSize: r.windowConfig.ToRectangle(),
	})
}
