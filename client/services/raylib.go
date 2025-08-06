package services

import (
	"chat/client/entities"
	"chat/client/widgets"
	"chat/shared"
	"chat/shared/rlutils"
	"errors"
	"math/rand/v2"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Widget interface {
	Draw()

	Contains(point rl.Vector2) bool

	HasCallback() bool
	Callback()
}

type RaylibConfig struct {
	WindowConfig entities.WindowConfig
}
type Raylib struct {
	windowConfig entities.WindowConfig

	running bool
	close   chan bool

	widgets map[entities.WidgetType][]Widget
}

func NewRaylib(cfg RaylibConfig) (*Raylib, error) {
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

	widget := &widgets.RectangleButton{
		RectangleInt32: rl.RectangleInt32{
			X:      100,
			Y:      100,
			Width:  100,
			Height: 100,
		},

		LoadingColor: rl.Blue,
		ErrColor:     rl.Black,
		ErrTextColor: rl.Red,
		ErrTextSize:  20,

		IdleColor: rl.RayWhite,

		IsAsync: true,
		CallbackFunc: func() error {
			time.Sleep(5 * time.Second)

			if rand.IntN(2) != 1 {
				return errors.New("test")
			}

			return nil
		},
	}

	return &Raylib{
		windowConfig: cfg.WindowConfig,

		running: false,
		close:   make(chan bool),

		widgets: map[entities.WidgetType][]Widget{
			entities.ButtonWidgetType: {
				widget,
			},
		},
	}, nil
}

func (r *Raylib) Init() error {
	eb := shared.NewErrorBuilder().Msg("failed to init raylib")

	if r.running {
		return eb.Causef("already running").Err()
	}

	rl.InitWindow(r.windowConfig.Width, r.windowConfig.Height, r.windowConfig.Title)
	r.running = true

	rl.SetTargetFPS(60)

	return nil
}

func (r *Raylib) Close() error {
	eb := shared.NewErrorBuilder().Msg("failed to close raylib")

	if !r.running {
		return eb.Causef("not running").Err()
	}

	rl.CloseWindow()
	r.running = false

	return nil
}

func (r *Raylib) Run() error {
	eb := shared.NewErrorBuilder().Msg("failed to start loop")

	if !r.running {
		return eb.Causef("not running").Err()
	}

	for {
		shouldStop := false

		select {
		case v := <-r.close:
			shouldStop = v
		default:
			shouldStop = rl.WindowShouldClose()
		}

		if shouldStop {
			break
		}

		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			r.HandleButtonPress()
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.DarkGray)

		for _, widgets := range r.widgets {
			for _, widget := range widgets {
				widget.Draw()
			}
		}

		rl.EndDrawing()
	}

	return nil
}

func (r *Raylib) HandleButtonPress() {
	point := rlutils.GetMousePosition()

	for _, button := range r.widgets[entities.ButtonWidgetType] {
		if button.HasCallback() && button.Contains(point) {
			button.Callback()
		}
	}
}
