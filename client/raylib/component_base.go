package raylib

import rl "github.com/gen2brain/raylib-go/raylib"

type BaseComponent struct {
	Color    rl.Color
	Children []Component
}
