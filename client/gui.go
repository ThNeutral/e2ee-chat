package client

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (c *Client) setupInitialLayout() {
	input := c.gui.AddRectangleWidget(200, 100, 200, 50)

	input.BackgroundColor = rl.White

	button := c.gui.AddRectangleWidget(200, 200, 200, 50)

	button.BackgroundColor = rl.LightGray
	button.TextColor = rl.Black

	button.Text = "text"

	button.OnClick = func() {
		log.Println(input.Text)
	}
}
