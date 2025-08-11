package raylib

import (
	"chat/client/entities"
	"chat/shared/rlutils"
	"log"

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

	widget, ok := r.focused.(*entities.RectangleWidget)
	if !ok {
		log.Println("handleInput: unhandled widget type")
		return
	}

	for _, char := range chars {
		if char == 127 {
			if len(widget.Text) != 0 {
				widget.Text = widget.Text[:len(widget.Text)-1]
			}
			continue
		}

		widget.Text += string(char)
	}
}

func (r *Raylib) handleMouseClick() {
	point := rlutils.GetMousePosition()

	isClickOnBackground := true
	for _, widget := range r.widgets {
		if r.contains(widget, point) {
			isClickOnBackground = false
			if onClick := r.onClick(widget); onClick != nil {
				onClick()
			}

			onFocus := r.onFocus(widget)
			if widget != r.focused {
				if r.focused != nil {
					if prevOnFocus := r.onFocus(r.focused); prevOnFocus != nil {
						prevOnFocus(false)
					}
				}

				r.focused = widget
				if onFocus != nil {
					onFocus(true)
				}
			}
		}
	}

	if isClickOnBackground {
		if r.focused != nil {
			if onFocus := r.onFocus(r.focused); onFocus != nil {
				onFocus(false)
			}
			r.focused = nil
		}
	}
}

func (r *Raylib) contains(widget Widget, point rlutils.Vector2) bool {
	rectangle, ok := widget.(*entities.RectangleWidget)
	if ok {
		return rectangle.X <= point.X &&
			rectangle.Y <= point.Y &&
			rectangle.X+rectangle.Width >= point.X &&
			rectangle.Y+rectangle.Height >= point.Y
	}

	log.Println("contains: unhandled widget type")
	return false
}

func (r *Raylib) onClick(widget Widget) entities.ClickEventHandler {
	button, ok := widget.(*entities.RectangleWidget)
	if ok {
		return button.OnClick
	}

	log.Println("onClick: unhandled widget type")
	return nil
}

func (r *Raylib) onFocus(widget Widget) entities.FocusEventHandler {
	button, ok := widget.(*entities.RectangleWidget)
	if ok {
		return button.OnFocus
	}

	log.Println("on: unhandled widget type")
	return nil
}
