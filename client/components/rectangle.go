package components

import (
	"chat/client/entities"
	"chat/shared/rlutils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectangleComponent struct {
	rl.RectangleInt32
	BaseComponent
}

func NewRectangle(rect rl.RectangleInt32, color rl.Color) *RectangleComponent {
	return &RectangleComponent{
		RectangleInt32: rect,
		BaseComponent: BaseComponent{
			Color:         color,
			ChildrenField: []entities.Component{},
		},
	}
}

func (rect *RectangleComponent) Type() entities.ComponentType {
	return entities.ComponentTypeRectangle
}

func (rect *RectangleComponent) Contains(point rlutils.Vector2) bool {
	return rlutils.Rect_Contains(rect.RectangleInt32, point)
}
