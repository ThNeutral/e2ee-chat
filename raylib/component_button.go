package raylib

import (
	"chat/raylib/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type buttonComponentProps struct {
	baseRectangleWidgetProps

	OnClick entities.ClickEventHandler
}

func buttonComponent(props buttonComponentProps) *entities.RectangleWidget {
	button := baseRectangleWidget(props.baseRectangleWidgetProps)

	button.OnClick = props.OnClick
	button.BackgroundColor = rl.Lime
	button.Text = "echo"

	return button
}
