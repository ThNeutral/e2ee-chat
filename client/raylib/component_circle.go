package raylib

import (
	"chat/client/raylib/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CircleComponent struct {
	entities.Circle
	BaseComponent
}

func NewCircleComponent(circle entities.Circle, color rl.Color) *CircleComponent {
	return &CircleComponent{
		Circle: circle,
		BaseComponent: BaseComponent{
			Color:    color,
			Children: []Component{},
		},
	}
}

func (c *CircleComponent) Type() ComponentType {
	return ComponentTypeCircle
}
