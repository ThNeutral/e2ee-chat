package rlutils

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawBorder(rect rl.RectangleInt32, size int32, color color.RGBA) {
	borderRect := rl.RectangleInt32{
		X:      rect.X - size,
		Y:      rect.Y - size,
		Width:  rect.Width + size*2,
		Height: rect.Height + size*2,
	}

	rl.DrawRectangle(borderRect.X, borderRect.Y, borderRect.Width, borderRect.Height, color)
}
