package runner

import (
	"chat/raylib/entities"
	"chat/shared/rlutils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Runner) draw() {
	rl.BeginDrawing()

	rl.ClearBackground(r.root.BackgroundColor)

	r.drawWidget(nil, r.root)

	rl.EndDrawing()
}

func (r *Runner) drawWidget(parentRect *rl.RectangleInt32, widget *entities.RectangleWidget) {
	// if r.focused == widget {
	// 	rlutils.DrawBorder(widget.RectangleInt32, widget.FocusBorderSize, widget.FocusBorderColor)
	// }

	actualRect := widget.RectangleInt32
	if parentRect != nil {
		actualRect.X += parentRect.X
		actualRect.Y += parentRect.Y
	}

	rl.DrawRectangle(actualRect.X, actualRect.Y, actualRect.Width, actualRect.Height, widget.BackgroundColor)

	if widget.Text != "" {
		rlutils.DrawCentralizedText(widget.RectangleInt32, widget.Text, widget.FontSize, widget.TextColor)
	}

	for _, child := range widget.Children {
		r.drawWidget(&actualRect, child)
	}
}
