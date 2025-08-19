package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type BaseComponent struct {
	Color         rl.Color
	ChildrenField []Component

	OnClickField OnClickHandler
}

func (c *BaseComponent) OnClick() OnClickHandler {
	return c.OnClickField
}

func (c *BaseComponent) Children() []Component {
	return c.ChildrenField
}

func (c *BaseComponent) AddChild(component Component) {
	c.ChildrenField = append(c.ChildrenField, component)
}
