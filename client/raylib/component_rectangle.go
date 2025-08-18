package raylib

import (
	"chat/shared/rlutils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectangleComponent struct {
	rl.RectangleInt32
	BaseComponent
}

func NewRectangleComponent(rect rl.RectangleInt32, color rl.Color) *RectangleComponent {
	return &RectangleComponent{
		RectangleInt32: rect,
		BaseComponent: BaseComponent{
			Color:         color,
			ChildrenField: []Component{},
		},
	}
}

func (rect *RectangleComponent) Type() ComponentType {
	return ComponentTypeRectangle
}

func (rect *RectangleComponent) Contains(point rlutils.Vector2) bool {
	return point.X >= rect.X && point.X <= rect.X+rect.Width && point.Y >= rect.Y && point.Y <= rect.Y+rect.Height
}
