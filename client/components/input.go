package components

import (
	"chat/client/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type InputComponent struct {
	*RectangleComponent

	OnInput_ entities.OnInputHandler
}

func (input *InputComponent) OnInput() entities.OnInputHandler {
	return input.OnInput_
}

func (input *InputComponent) Type() entities.ComponentType {
	return entities.ComponentTypeInput
}

func NewInput(rect rl.RectangleInt32, color rl.Color, onInput entities.OnInputHandler) *InputComponent {
	handler := func(component entities.Component, chars []rune) {
		input := component.(*InputComponent)

		for _, char := range chars {
			if char == 127 {
				if len(input.Text_) != 0 {
					input.Text_ = input.Text_[0 : len(input.Text_)-1]
				}
				continue
			}

			input.Text_ += string(char)
		}

		if onInput != nil {
			onInput(component, chars)
		}
	}

	input := &InputComponent{
		RectangleComponent: NewRectangle(rect, color),
		OnInput_:           handler,
	}

	return input
}
