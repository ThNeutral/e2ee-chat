package rlutils

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawCentralizedText(rect rl.RectangleInt32, text string, fontSize int32, color color.RGBA) {
	textWidth := rl.MeasureText(text, fontSize)

	width := rect.X + rect.Width/2 - textWidth/2
	height := rect.Y + rect.Height/2 - fontSize/2

	rl.DrawText(text, width, height, fontSize, color)
}

func GetInputForLastFrame() []rune {
	var input []rune

	char := rl.GetCharPressed()
	for char > 0 {
		if (char >= 32) && (char <= 125) {
			input = append(input, rune(char))
		}
		char = rl.GetCharPressed()
	}

	if rl.IsKeyPressed(rl.KeyBackspace) {
		input = append(input, 127)
	}

	return input
}
