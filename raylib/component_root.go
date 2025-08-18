package raylib

import (
	"chat/raylib/entities"
	"context"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func rootComponent(client *Raylib) entities.RootComponent {
	return func(props entities.RootComponentProps) *entities.RectangleWidget {
		root := baseRectangleWidget(baseRectangleWidgetProps{
			Position: props.WindowSize,
		})

		input := inputComponent(inputComponentProps{
			baseRectangleWidgetProps: baseRectangleWidgetProps{
				Position: rl.RectangleInt32{
					X:      50,
					Y:      50,
					Width:  100,
					Height: 25,
				},
			},
		})

		button := buttonComponent(buttonComponentProps{
			baseRectangleWidgetProps: baseRectangleWidgetProps{
				Position: rl.RectangleInt32{
					X:      50,
					Y:      100,
					Width:  50,
					Height: 25,
				},
			},
			OnClick: func(this *entities.RectangleWidget, event *entities.ClickEvent) {
				go func() {
					client.echo.Echo(context.Background(), input.Text)
				}()
			},
		})

		root.Children = []*entities.RectangleWidget{
			input,
			button,
		}

		return root
	}
}
