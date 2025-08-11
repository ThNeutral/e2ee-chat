package client

import (
	"chat/client/entities"
	"chat/client/utils"
	"context"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (c *Client) setupInitialLayout() {
	label := c.gui.AddRectangleWidget(200, 100, 200, 150)
	label.BackgroundColor = color.RGBA{0, 0, 0, 0}

	input := c.gui.AddRectangleWidget(200, 250, 200, 50)
	input.BackgroundColor = rl.White
	input.OnChange = utils.InputChangeHandler

	button := c.gui.AddRectangleWidget(200, 320, 200, 50)
	button.BackgroundColor = rl.LightGray
	button.TextColor = rl.Black
	button.Text = "text"
	button.OnClick = buttonClickHandler(c, input, label)
}

func buttonClickHandler(c *Client, input, label *entities.RectangleWidget) entities.ClickEventHandler {
	return func(this *entities.RectangleWidget) {
		this.OnClick = nil

		this.BackgroundColor = rl.Blue
		this.Text = "loading"

		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), c.defaultTimeout)
			defer cancel()

			resp, err := c.echo.Echo(ctx, input.Text)
			if err != nil {
				label.Text = err.Error()
				label.TextColor = rl.Red
			} else {
				label.Text = resp
				label.TextColor = rl.Black
			}

			this.BackgroundColor = rl.LightGray

			this.OnClick = buttonClickHandler(c, input, label)
		}()
	}
}
