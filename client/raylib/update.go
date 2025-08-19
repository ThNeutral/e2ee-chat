package raylib

import (
	"chat/client/entities"
	"chat/shared/rlutils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Raylib) update() {
	input := rlutils.GetKeyboardInputForLastFrame()
	if len(input) != 0 {
		r.handleKeyboardInput(input)
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		r.handleMouseClick(rlutils.GetMousePosition())
	}
}

func (r *Raylib) handleKeyboardInput(frameInput []rune) {
	if r.focused == nil {
		return
	}

	r.focused.OnInput()(r.focused, frameInput)
}

func (r *Raylib) handleMouseClick(point rlutils.Vector2) {
	chain := r.buildComponentChain(r.root, point)

	for _, next := range chain {
		if handler := next.OnInput(); handler != nil {
			r.focused = next
		}

		if handler := next.OnClick(); handler != nil {
			handler()
			break
		}
	}
}

func (r *Raylib) buildComponentChain(c entities.Component, point rlutils.Vector2) []entities.Component {
	if !c.Contains(point) {
		return []entities.Component{}
	}

	var chain []entities.Component
	for _, child := range c.Children() {
		ch := r.buildComponentChain(child, point)
		if len(ch) != 0 {
			chain = ch
			break
		}
	}

	return append(chain, c)
}
