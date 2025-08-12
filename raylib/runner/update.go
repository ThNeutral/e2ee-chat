package runner

import (
	"chat/raylib/entities"
	"chat/shared/rlutils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Runner) update() {
	if input := rlutils.GetInputForLastFrame(); len(input) != 0 {
		r.handleInput(input)
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		r.handleMouseClick()
	}
}

func (r *Runner) handleInput(input []rune) {
	if r.focused != nil && r.focused.OnInput != nil {
		event := &entities.InputEvent{
			Text: input,
		}
		r.focused.OnInput(event)
	}
}

func (r *Runner) handleMouseClick() {
	mousePosition := rlutils.GetMousePosition()

	widgets := buildClickHandlerChain(nil, r.root, mousePosition)

	r.focused = nil
	for _, widget := range widgets {
		if widget.Focusable {
			r.focused = widget
		}

		if widget.OnClick == nil {
			break
		}

		event := &entities.ClickEvent{
			ShouldPropagate: false,
		}

		widget.OnClick(event)

		if !event.ShouldPropagate {
			break
		}
	}
}

func buildClickHandlerChain(parentRect *rl.RectangleInt32, widget *entities.RectangleWidget, mousePosition rlutils.Vector2) []*entities.RectangleWidget {
	actualRect := widget.RectangleInt32
	if parentRect != nil {
		actualRect.X += parentRect.X
		actualRect.Y += parentRect.Y
	}

	if !containsRect(actualRect, mousePosition) {
		return []*entities.RectangleWidget{}
	}

	for _, child := range widget.Children {
		prev := buildClickHandlerChain(&actualRect, child, mousePosition)
		if len(prev) != 0 {
			return append(prev, widget)
		}
	}

	return []*entities.RectangleWidget{widget}
}

func containsRect(rect rl.RectangleInt32, point rlutils.Vector2) bool {
	return point.X >= rect.X && point.X <= rect.X+rect.Width && point.Y >= rect.Y && point.Y <= rect.Y+rect.Height
}
