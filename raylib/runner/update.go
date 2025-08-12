package runner

import (
	"chat/raylib/entities"
	"chat/shared/rlutils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Runner) update() {
	// if input := rlutils.GetInputForLastFrame(); len(input) != 0 {
	// 	r.handleInput(input)
	// }

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		r.handleMouseClick()
	}
}

func (r *Runner) handleMouseClick() {
	mousePosition := rlutils.GetMousePosition()

	handlers := buildClickHandlerChain(nil, r.root, mousePosition)

	for _, handler := range handlers {
		if handler == nil {
			break
		}

		event := &entities.ClickEvent{
			Event: entities.Event{
				ShouldPropagate: false,
			},
		}

		handler(event)

		if !event.ShouldPropagate {
			break
		}
	}
}

func buildClickHandlerChain(parentRect *rl.RectangleInt32, widget *entities.RectangleWidget, mousePosition rlutils.Vector2) []entities.ClickEventHandler {
	actualRect := widget.RectangleInt32
	if parentRect != nil {
		actualRect.X += parentRect.X
		actualRect.Y += parentRect.Y
	}

	if !containsRect(actualRect, mousePosition) {
		return []entities.ClickEventHandler{}
	}

	for _, child := range widget.Children {
		prev := buildClickHandlerChain(&actualRect, child, mousePosition)
		if len(prev) != 0 {
			return append(prev, widget.OnClick)
		}
	}

	return []entities.ClickEventHandler{widget.OnClick}
}

func containsRect(rect rl.RectangleInt32, point rlutils.Vector2) bool {
	return point.X >= rect.X && point.X <= rect.X+rect.Width && point.Y >= rect.Y && point.Y <= rect.Y+rect.Height
}
