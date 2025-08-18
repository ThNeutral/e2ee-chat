package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type RootComponentProps struct {
	WindowSize rl.RectangleInt32
}

type RootComponent func(RootComponentProps) *RectangleWidget
