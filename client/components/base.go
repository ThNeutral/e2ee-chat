package components

import (
	"chat/client/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BaseComponent struct {
	Color         rl.Color
	ChildrenField []entities.Component

	OnClickField entities.OnClickHandler
}

func (c *BaseComponent) OnClick() entities.OnClickHandler {
	return c.OnClickField
}

func (c *BaseComponent) Children() []entities.Component {
	return c.ChildrenField
}

func (c *BaseComponent) AddChild(component entities.Component) {
	c.ChildrenField = append(c.ChildrenField, component)
}
