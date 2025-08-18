package rlutils

import rl "github.com/gen2brain/raylib-go/raylib"

func GetMousePosition() Vector2 {
	mousePosition := rl.GetMousePosition()

	return Vector2{
		X: int32(mousePosition.X),
		Y: int32(mousePosition.Y),
	}
}
