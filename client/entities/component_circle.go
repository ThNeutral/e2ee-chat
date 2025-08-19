package entities

import (
	"chat/shared/rlutils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CircleComponent struct {
	rlutils.Circle
	BaseComponent
}

func NewCircleComponent(circle rlutils.Circle, color rl.Color) *CircleComponent {
	return &CircleComponent{
		Circle: circle,
		BaseComponent: BaseComponent{
			Color:         color,
			ChildrenField: []Component{},
		},
	}
}

func (c *CircleComponent) Type() ComponentType {
	return ComponentTypeCircle
}

func (c *CircleComponent) Contains(point rlutils.Vector2) bool {
	return rlutils.Circle_Contains(c.Circle, point)
}
