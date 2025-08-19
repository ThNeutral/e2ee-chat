package client

import (
	"chat/client/components"
	"chat/shared/rlutils"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (c *Client) init() {
	root := c.gui.Root()

	circle := components.NewCircle(rlutils.Circle{
		Center: rlutils.Vector2{
			X: 150,
			Y: 150,
		},
		Radius: 25,
	}, rl.Red)

	circle.OnClickField = func() {
		if c.websocket.IsConnected() {
			err := c.websocket.Disconnect("manual disconnect")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := c.websocket.Connect()
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	c.websocket.SetOnConnectHandler(func() {
		circle.Color = rl.Green
	})

	c.websocket.SetOnDisconnectHandler(func() {
		circle.Color = rl.Red
	})

	root.AddChild(circle)
}
