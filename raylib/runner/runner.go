package runner

import (
	"chat/raylib/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
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

	r := &Runner{
		windowConfig: cfg.WindowConfig,

		root:    root,
		focused: nil,

		running: false,
	}

	r.rootComponent()

	r.root.Children = append(
		r.root.Children,
		r.inputComponent(inputComponentParams{
			RectangleInt32: rl.RectangleInt32{
				X:      200,
				Y:      300,
				Width:  100,
				Height: 50,
			},
			onInput: nil,
		}),
	)

	return r
}
