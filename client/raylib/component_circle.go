package raylib

import (
	"chat/client/raylib/entities"
	"chat/shared/rlutils"
	"fmt"

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
			Color:         color,
			ChildrenField: []Component{},
		},
	}
}

func (c *CircleComponent) Type() ComponentType {
	return ComponentTypeCircle
}

func (c *CircleComponent) Contains(point rlutils.Vector2) bool {
	fmt.Printf("distance - %v, radius - %v\n", rlutils.V2_Distance(c.Center, point), c.Radius)

	return rlutils.V2_Distance(c.Center, point) <= c.Radius
}
