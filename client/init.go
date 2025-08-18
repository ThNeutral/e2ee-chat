package client

import (
	"chat/client/raylib"
	"chat/client/raylib/entities"
	"chat/shared/rlutils"
	"image/color"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (c *Client) init() {
	root := raylib.NewRectangleComponent(rl.RectangleInt32{
		X:      100,
		Y:      100,
		Width:  400,
		Height: 400,
	}, rl.RayWhite)

	root.OnClickField = func(c raylib.Component) {
		root.Color = color.RGBA{
			R: uint8(rand.IntN(255)),
			G: uint8(rand.IntN(255)),
			B: uint8(rand.IntN(255)),
			A: 255,
		}
	}

	c.gui.SetRootComponent(root)

	circle := raylib.NewCircleComponent(entities.Circle{
		Center: rlutils.Vector2{
			X: 150,
			Y: 150,
		},
		Radius: 25,
	}, rl.Red)

	circle.OnClickField = func(c raylib.Component) {
		circle.Color = color.RGBA{
			R: uint8(rand.IntN(255)),
			G: uint8(rand.IntN(255)),
			B: uint8(rand.IntN(255)),
			A: 255,
		}
	}

	root.ChildrenField = append(root.ChildrenField, circle)
}
