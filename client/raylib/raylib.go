package raylib

import (
	"chat/client/components"
	"chat/client/entities"
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
	windowName      string
	targetFramerate int

	running bool
	root    *components.RectangleComponent
}

func New(cfg Config) *Raylib {
	if cfg.WindowName == "" {
		cfg.WindowName = "WINDOW"
	}

	if cfg.TargetFramerate == 0 {
		cfg.TargetFramerate = 60
	}

	root := components.NewRectangle(rl.RectangleInt32{
		X:      0,
		Y:      0,
		Width:  cfg.Size.X,
		Height: cfg.Size.Y,
	}, cfg.BackgroundColor)

	root.OnClickField = func() {}

	return &Raylib{
		running: false,
		root:    root,

		windowName:      cfg.WindowName,
		targetFramerate: cfg.TargetFramerate,
	}
}

func (r *Raylib) Init() error {
	if r.running {
		return fmt.Errorf("already running")
	}

	rl.InitWindow(r.root.Width, r.root.Height, r.windowName)
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

func (r *Raylib) Root() entities.Component {
	return r.root
}
