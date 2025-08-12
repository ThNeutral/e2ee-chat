package runner

import (
	"chat/raylib/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Runner) AddRectangleWidget(x int32, y int32, width int32, height int32) *entities.RectangleWidget {
	button := &entities.RectangleWidget{
		RectangleInt32: rl.RectangleInt32{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		},

		BackgroundColor: rl.LightGray,
		TextColor:       rl.Black,

		Text:     "",
		FontSize: 16,

		FocusBorderColor: rl.Black,
		FocusBorderSize:  2,

		OnClick:  nil,
		OnFocus:  nil,
		OnChange: nil,
	}

	r.widgets = append(r.widgets, button)

	return button
}
