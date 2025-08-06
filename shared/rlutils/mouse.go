package rlutils

import rl "github.com/gen2brain/raylib-go/raylib"

func GetMousePosition() rl.Vector2 {
	return rl.Vector2{
		X: float32(rl.GetMouseX()),
		Y: float32(rl.GetMouseY()),
	}
}
