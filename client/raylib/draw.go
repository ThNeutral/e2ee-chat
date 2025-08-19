package raylib

import (
	"chat/client/components"
	"chat/client/entities"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Raylib) draw() {
	r.drawComponent(r.root)
}

func (r *Raylib) drawComponent(c entities.Component) {
	switch c.Type() {
	case entities.ComponentTypeCircle:
		r.drawCircle(c.(*components.CircleComponent))
	case entities.ComponentTypeRectangle:
		r.drawRectangle(c.(*components.RectangleComponent))
	case entities.ComponentTypeInput:
		r.drawRectangle(c.(*components.InputComponent).RectangleComponent)
	default:
		log.Fatalf("unknown component type: %v\n", c.Type())
	}
}

func (r *Raylib) drawCircle(c *components.CircleComponent) {
	rl.DrawCircle(c.Center.X, c.Center.Y, c.Radius, c.Color)
	r.drawChildren(c)
}

func (r *Raylib) drawRectangle(rect *components.RectangleComponent) {
	rl.DrawRectangle(rect.X, rect.Y, rect.Width, rect.Height, rect.Color)
	r.drawChildren(rect)

	rl.DrawText(rect.Text_, rect.X, rect.Y, 16, rl.Black)
}

func (r *Raylib) drawChildren(c entities.Component) {
	for _, child := range c.Children() {
		r.drawComponent(child)
	}
}
