package rlutils

import rl "github.com/gen2brain/raylib-go/raylib"

func GetKeyboardInputForLastFrame() []rune {
	chars := make([]rune, 0, 5)

	if rl.IsKeyPressed(rl.KeyBackspace) {
		chars = append(chars, 127)
	}

	char := rl.GetCharPressed()
	for char > 0 {
		if 32 <= char && char <= 125 {
			chars = append(chars, char)
		}

		char = rl.GetCharPressed()
	}

	return chars
}
