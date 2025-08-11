package rlutils

import rl "github.com/gen2brain/raylib-go/raylib"

func GetMousePosition() Vector2 {
	return Vector2{
		X: rl.GetMouseX(),
		Y: rl.GetMouseY(),
	}
}
