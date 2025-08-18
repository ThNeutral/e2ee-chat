package raylib

import rl "github.com/gen2brain/raylib-go/raylib"

type RectangleComponent struct {
	rl.RectangleInt32
	BaseComponent
}

func NewRectangleComponent(rect rl.RectangleInt32, color rl.Color) *RectangleComponent {
	return &RectangleComponent{
		RectangleInt32: rect,
		BaseComponent: BaseComponent{
			Color:    color,
			Children: []Component{},
		},
	}
}

func (rect *RectangleComponent) Type() ComponentType {
	return ComponentTypeRectangle
}
