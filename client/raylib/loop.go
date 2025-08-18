package raylib

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Raylib) Run() error {
	if !r.running {
		return fmt.Errorf("not running")
	}

	for !r.shouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(r.config.BackgroundColor)

		r.draw()

		rl.EndDrawing()
	}

	return nil
}

func (r *Raylib) shouldClose() bool {
	return rl.WindowShouldClose()
}
