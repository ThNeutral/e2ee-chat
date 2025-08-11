package raylib

import (
	"chat/client/entities"
	"chat/shared/rlutils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Raylib) update() {
	if input := rlutils.GetInputForLastFrame(); len(input) != 0 {
		r.handleInput(input)
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		r.handleMouseClick()
	}
}

func (r *Raylib) handleInput(chars []rune) {
	if r.focused == nil {
		return
	}

	if onChange := r.onChange(r.focused); onChange != nil {
		onChange(r.focused, chars)
	}
}

func (r *Raylib) handleMouseClick() {
	point := rlutils.GetMousePosition()

	isClickOnBackground := true
	for _, widget := range r.widgets {
		if r.contains(widget, point) {
			isClickOnBackground = false
			if onClick := r.onClick(widget); onClick != nil {
				onClick(widget)
			}

			onFocus := r.onFocus(widget)
			if widget != r.focused {
				if r.focused != nil {
					if prevOnFocus := r.onFocus(r.focused); prevOnFocus != nil {
						prevOnFocus(widget, false)
					}
				}

				r.focused = widget
				if onFocus != nil {
					onFocus(widget, true)
				}
			}
		}
	}

	if isClickOnBackground {
		if r.focused != nil {
			if onFocus := r.onFocus(r.focused); onFocus != nil {
				onFocus(r.focused, false)
			}
			r.focused = nil
		}
	}
}

func (r *Raylib) contains(widget *entities.RectangleWidget, point rlutils.Vector2) bool {
	return widget.X <= point.X &&
		widget.Y <= point.Y &&
		widget.X+widget.Width >= point.X &&
		widget.Y+widget.Height >= point.Y
}

func (r *Raylib) onClick(widget *entities.RectangleWidget) entities.ClickEventHandler {
	return widget.OnClick
}

func (r *Raylib) onFocus(widget *entities.RectangleWidget) entities.FocusEventHandler {
	return widget.OnFocus
}

func (r *Raylib) onChange(widget *entities.RectangleWidget) entities.ChangeEventHandler {
	return widget.OnChange
}
