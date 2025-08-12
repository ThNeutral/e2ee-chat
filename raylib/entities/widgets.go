package entities

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type WindowConfig struct {
	Width           int32
	Height          int32
	Title           string
	BackgroundColor color.RGBA
}

func (cfg WindowConfig) ToRectangle() rl.RectangleInt32 {
	return rl.RectangleInt32{
		X:      0,
		Y:      0,
		Width:  cfg.Width,
		Height: cfg.Height,
	}
}

type WidgetType int

const (
	ButtonWidgetType WidgetType = iota
)

type RectangleWidget struct {
	rl.RectangleInt32

	BackgroundColor color.RGBA
	TextColor       color.RGBA

	FontSize int32
	Text     string

	Focusable        bool
	FocusBorderColor color.RGBA
	FocusBorderSize  int32

	OnClick ClickEventHandler
	OnFocus FocusEventHandler
	OnInput InputEventHandler

	Children []*RectangleWidget
}
