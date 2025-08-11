package raylib

import (
	"chat/client/entities"
	"chat/shared/rlutils"
	"log"

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

func (r *Raylib) drawWidget(widget Widget) {
	rectangle, ok := widget.(*entities.RectangleWidget)
	if ok {
		if r.focused == rectangle {
			rlutils.DrawBorder(rectangle.RectangleInt32, rectangle.FocusBorderSize, rectangle.FocusBorderColor)
		}

		rl.DrawRectangle(rectangle.X, rectangle.Y, rectangle.Width, rectangle.Height, rectangle.BackgroundColor)

		if rectangle.Text != "" {
			rlutils.DrawCentralizedText(rectangle.RectangleInt32, rectangle.Text, rectangle.FontSize, rectangle.TextColor)
		}

		return
	}

	log.Println("drawWidget: unhandled widget type")
}
