package raylib

import (
	"chat/raylib/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type baseRectangleWidgetProps struct {
	Position rl.RectangleInt32
}

func baseRectangleWidget(props baseRectangleWidgetProps) *entities.RectangleWidget {
	button := &entities.RectangleWidget{
		RectangleInt32: props.Position,

		BackgroundColor: rl.LightGray,
		TextColor:       rl.Black,

		Text:     "",
		FontSize: 16,

		Focusable:        false,
		FocusBorderColor: rl.Black,
		FocusBorderSize:  2,

		OnClick: nil,
		OnFocus: nil,
		OnInput: nil,

		Children: []*entities.RectangleWidget{},
	}

	return button
}
