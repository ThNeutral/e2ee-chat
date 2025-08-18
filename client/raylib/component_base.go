package raylib

import rl "github.com/gen2brain/raylib-go/raylib"

type OnClickHandler func(c Component)

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
