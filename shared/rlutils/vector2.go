package rlutils

import rl "github.com/gen2brain/raylib-go/raylib"

type Vector2 struct {
	X int32
	Y int32
}

func ToFloatVector(v2 Vector2) rl.Vector2 {
	return rl.Vector2{
		X: float32(v2.X),
		Y: float32(v2.Y),
	}
}

func ToIntVector(v2 rl.Vector2) Vector2 {
	return Vector2{
		X: int32(v2.X),
		Y: int32(v2.Y),
	}
}
