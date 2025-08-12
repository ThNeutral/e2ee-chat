package raylib

import (
	"chat/shared"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Raylib) Init() error {
	eb := shared.B().Msg("failed to init raylib")

	if r.running {
		return eb.Causef("already running").Err()
	}

	rl.InitWindow(r.windowConfig.Width, r.windowConfig.Height, r.windowConfig.Title)
	r.running = true

	rl.SetTargetFPS(60)

	return nil
}

func (r *Raylib) Close() error {
	eb := shared.B().Msg("failed to close raylib")

	if !r.running {
		return eb.Causef("not running").Err()
	}

	rl.CloseWindow()
	r.running = false

	return nil
}

func (r *Raylib) Run() error {
	eb := shared.B().Msg("failed to start loop")

	if !r.running {
		return eb.Causef("not running").Err()
	}

	for !r.shouldClose() {
		r.update()

		r.draw()
	}

	return nil
}

func (r *Raylib) shouldClose() bool {
	shouldStop := false

	select {
	case v := <-r.close:
		shouldStop = v
	default:
		shouldStop = rl.WindowShouldClose()
	}

	return shouldStop
}
