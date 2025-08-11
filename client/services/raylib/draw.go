package raylib

import (
	"chat/client/entities"
	"chat/shared/rlutils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Raylib) draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.DarkGray)

	for _, widget := range r.widgets {
		r.drawWidget(widget)
	}

	rl.EndDrawing()
}

func (r *Raylib) drawWidget(widget *entities.RectangleWidget) {
	if r.focused == widget {
		rlutils.DrawBorder(widget.RectangleInt32, widget.FocusBorderSize, widget.FocusBorderColor)
	}

	rl.DrawRectangle(widget.X, widget.Y, widget.Width, widget.Height, widget.BackgroundColor)

	if widget.Text != "" {
		rlutils.DrawCentralizedText(widget.RectangleInt32, widget.Text, widget.FontSize, widget.TextColor)
	}

	return
}
