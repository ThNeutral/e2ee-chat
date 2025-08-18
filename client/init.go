package client

import (
	"chat/client/raylib"
	"chat/client/raylib/entities"
	"chat/shared/rlutils"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (c *Client) init() {
	root := raylib.NewRectangleComponent(rl.RectangleInt32{
		X:      100,
		Y:      100,
		Width:  400,
		Height: 400,
	}, rl.RayWhite)

	c.gui.SetRootComponent(root)

	circle := raylib.NewCircleComponent(entities.Circle{
		Center: rlutils.Vector2{
			X: 150,
			Y: 150,
		},
		Radius: 25,
	}, rl.Red)

	circle.OnClickField = func(_ raylib.Component) {
		err := c.websocket.Connect()
		if err != nil {
			fmt.Println(err)
		}
	}

	root.ChildrenField = append(root.ChildrenField, circle)
}
