package runner

import (
	"chat/shared/errs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Runner) Init() error {
	eb := errs.B().Msg("failed to init raylib")

	if r.running {
		return eb.Causef("already running").Err()
	}

	rl.InitWindow(r.windowConfig.Width, r.windowConfig.Height, r.windowConfig.Title)
	r.running = true

	rl.SetTargetFPS(60)

	return nil
}

func (r *Runner) Close() error {
	eb := errs.B().Msg("failed to close raylib")

	if !r.running {
		return eb.Causef("not running").Err()
	}

	rl.CloseWindow()
	r.running = false

	return nil
}

func (r *Runner) Run() error {
	eb := errs.B().Msg("failed to start loop")

	if !r.running {
		return eb.Causef("not running").Err()
	}

	for !r.shouldClose() {

		r.update()

		r.draw()
	}

	return nil
}

func (r *Runner) shouldClose() bool {
	shouldStop := false
	shouldStop = rl.WindowShouldClose()

	return shouldStop
}
