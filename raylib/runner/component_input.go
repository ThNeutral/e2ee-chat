package runner

import (
	"chat/raylib/entities"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type inputComponentParams struct {
	rl.RectangleInt32

	onInput entities.InputEventHandler
}

func (r *Runner) inputComponent(params inputComponentParams) *entities.RectangleWidget {
	input := getRectangleWidget(params.X, params.Y, params.Width, params.Height)

	input.Focusable = true
	input.OnInput = func(event *entities.InputEvent) {
		if params.onInput != nil {
			params.onInput(event)
			return
		}

		inputChangeHandler(input, event.Text)
	}

	input.BackgroundColor = color.RGBA{255, 255, 255, 255}

	return input
}

func inputChangeHandler(this *entities.RectangleWidget, text []rune) {
	for _, char := range text {
		if char == 127 {
			length := len(this.Text)
			if length != 0 {
				this.Text = this.Text[:length-1]
			}
			continue
		}

		this.Text += string(char)
	}
}
