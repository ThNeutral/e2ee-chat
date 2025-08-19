package rlutils

import rl "github.com/gen2brain/raylib-go/raylib"

func Rect_Contains(rect rl.RectangleInt32, point Vector2) bool {
	return point.X >= rect.X && point.X <= rect.X+rect.Width && point.Y >= rect.Y && point.Y <= rect.Y+rect.Height
}
