package raylib

import (
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
	root    *RectangleComponent
}

func New(cfg Config) *Raylib {
	if cfg.WindowName == "" {
		cfg.WindowName = "WINDOW"
	}

	if cfg.TargetFramerate == 0 {
		cfg.TargetFramerate = 60
	}

	root := NewRectangleComponent(rl.RectangleInt32{
		X:      0,
		Y:      0,
		Width:  cfg.Size.X,
		Height: cfg.Size.Y,
	}, cfg.BackgroundColor)

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

func (r *Raylib) SetRootComponent(component Component) {
	r.root.Children = []Component{component}
	fmt.Println(r.root.Children)
}
