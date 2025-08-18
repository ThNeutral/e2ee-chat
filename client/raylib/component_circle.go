package raylib

import (
	"chat/client/raylib/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CircleComponent struct {
	entities.Circle
	Color rl.Color
}

func componentCircle(circle entities.Circle, color rl.Color) *CircleComponent {
	return &CircleComponent{
		Circle: circle,
		Color:  color,
	}
}

func (c *CircleComponent) Type() ComponentType {
	return ComponentTypeCircle
}
