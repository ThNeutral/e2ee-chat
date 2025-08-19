package components

import (
	"chat/client/entities"
	"chat/shared/rlutils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CircleComponent struct {
	rlutils.Circle
	BaseComponent
}

func NewCircle(circle rlutils.Circle, color rl.Color) *CircleComponent {
	return &CircleComponent{
		Circle: circle,
		BaseComponent: BaseComponent{
			Color:         color,
			ChildrenField: []entities.Component{},
		},
	}
}

func (c *CircleComponent) Type() entities.ComponentType {
	return entities.ComponentTypeCircle
}

func (c *CircleComponent) Contains(point rlutils.Vector2) bool {
	return rlutils.Circle_Contains(c.Circle, point)
}
