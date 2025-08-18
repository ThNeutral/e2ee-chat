package raylib

import (
	"chat/raylib/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type inputComponentProps struct {
	baseRectangleWidgetProps

	OnInput entities.InputEventHandler
}

func inputComponent(props inputComponentProps) *entities.RectangleWidget {
	input := baseRectangleWidget(baseRectangleWidgetProps{
		Position: props.Position,
	})

	input.BackgroundColor = rl.White

	input.Focusable = true
	input.OnInput = func(event *entities.InputEvent) {
		if props.OnInput != nil {
			props.OnInput(event)
		}

		for _, char := range event.Text {
			if char == 127 {
				if len(input.Text) != 0 {
					input.Text = input.Text[:len(input.Text)-1]
				}
				continue
			}
			input.Text += string(char)
		}
	}

	return input
}
