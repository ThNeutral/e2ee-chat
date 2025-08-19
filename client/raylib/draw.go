package raylib

import (
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
		r.drawCircle(c.(*entities.CircleComponent))
	case entities.ComponentTypeRectangle:
		r.drawRectangle(c.(*entities.RectangleComponent))
	default:
		log.Fatalf("unknown component type: %v\n", c.Type())
	}
}

func (r *Raylib) drawCircle(c *entities.CircleComponent) {
	rl.DrawCircle(c.Center.X, c.Center.Y, c.Radius, c.Color)
	r.drawChildren(c)
}

func (r *Raylib) drawRectangle(rect *entities.RectangleComponent) {
	rl.DrawRectangle(rect.X, rect.Y, rect.Width, rect.Height, rect.Color)
	r.drawChildren(rect)
}

func (r *Raylib) drawChildren(c entities.Component) {
	for _, child := range c.Children() {
		r.drawComponent(child)
	}
}
