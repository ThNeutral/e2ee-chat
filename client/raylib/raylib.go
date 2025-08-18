package raylib

import (
	"chat/client/raylib/entities"
	"chat/shared/rlutils"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Config struct {
	Size            rlutils.Vector2
	BackgroundColor rl.Color
	WindowName      string
	TargetFramerate int
}

type Raylib struct {
	config Config

	running bool
	root    Component
}

func New(cfg Config) *Raylib {
	if cfg.WindowName == "" {
		cfg.WindowName = "WINDOW"
	}

	if cfg.TargetFramerate == 0 {
		cfg.TargetFramerate = 60
	}

	return &Raylib{
		config: cfg,

		running: false,

		root: componentCircle(entities.Circle{
			Center: rlutils.Vector2{
				X: 50,
				Y: 50,
			},
			Radius: 75,
		}, rl.Red),
	}
}

func (r *Raylib) Init() error {
	if r.running {
		return fmt.Errorf("already running")
	}

	rl.InitWindow(r.config.Size.X, r.config.Size.Y, r.config.WindowName)
	r.running = true

	return nil
}

func (r *Raylib) Close() error {
	if !r.running {
		return fmt.Errorf("not running")
	}

	rl.CloseWindow()
	r.running = false

	return nil
}
