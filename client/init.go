package client

import (
	"chat/client/entities"
	"chat/shared/rlutils"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (c *Client) init() {
	root := c.gui.GetRootComponent()

	circle := entities.NewCircleComponent(rlutils.Circle{
		Center: rlutils.Vector2{
			X: 150,
			Y: 150,
		},
		Radius: 25,
	}, rl.Red)

	circle.OnClickField = func() {
		err := c.websocket.Connect()
		if err != nil {
			fmt.Println(err)
		}
	}

	root.AddChild(circle)
}
