package runner

import (
	"chat/raylib/entities"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (r *Runner) rootComponent() {
	text := getRectangleWidget(220, 100, 100, 100)
	text.BackgroundColor = color.RGBA{}

	r.root.Children = append(r.root.Children, text)

	r.root.OnClick = func(event *entities.ClickEvent) {
		text.Text = ""
	}

	box := getRectangleWidget(100, 100, 100, 100)
	box.BackgroundColor = rl.LightGray
	box.OnClick = func(event *entities.ClickEvent) {
		text.Text = "box clicked"
	}

	r.root.Children = append(r.root.Children, box)

	box2 := getRectangleWidget(50, 50, 50, 50)
	box2.BackgroundColor = rl.SkyBlue
	box2.OnClick = func(event *entities.ClickEvent) {
		text.Text = "box2 clicked"
	}

	box.Children = append(box.Children, box2)

	box3 := getRectangleWidget(25, 25, 25, 25)
	box.BackgroundColor = rl.Brown

	box2.Children = append(box2.Children, box3)
}
