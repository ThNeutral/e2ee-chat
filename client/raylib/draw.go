package raylib

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Raylib) draw() {
	r.drawComponent(r.root)
}

func (r *Raylib) drawComponent(c Component) {
	switch c.Type() {
	case ComponentTypeCircle:
		r.drawCircle(c.(*CircleComponent))
	case ComponentTypeRectangle:

	default:
		log.Fatalf("unknown component type: %v\n", c.Type())
	}
}

func (r *Raylib) drawCircle(c *CircleComponent) {
	rl.DrawCircle(c.Center.X, c.Center.Y, c.Radius, c.Color)
}
