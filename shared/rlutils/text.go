package rlutils

import (
	"image/color"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawCentralizedText(rect rl.RectangleInt32, text string, fontSize int32, color color.RGBA) {
	words := strings.Fields(text)
	var lines []string
	var currentLine string

	for _, word := range words {
		testLine := word
		if currentLine != "" {
			testLine = currentLine + " " + word
		}

		if rl.MeasureText(testLine, fontSize) > int32(rect.Width) {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			currentLine = testLine
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	totalTextHeight := int32(len(lines)) * fontSize
	startY := rect.Y + (rect.Height/2 - totalTextHeight/2)

	for i, line := range lines {
		textWidth := rl.MeasureText(line, fontSize)
		x := rect.X + rect.Width/2 - textWidth/2
		y := startY + int32(i)*fontSize
		rl.DrawText(line, x, y, fontSize, color)
	}
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
