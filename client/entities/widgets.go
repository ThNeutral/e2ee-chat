package entities

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type WindowConfig struct {
	Width  int32
	Height int32
	Title  string
}

type WidgetType int

const (
	ButtonWidgetType WidgetType = iota
)

type ClickEventHandler func(this *RectangleWidget)
type FocusEventHandler func(this *RectangleWidget, focused bool)
type ChangeEventHandler func(this *RectangleWidget, text []rune)

type RectangleWidget struct {
	rl.RectangleInt32

	BackgroundColor color.RGBA
	TextColor       color.RGBA

	FontSize int32
	Text     string

	FocusBorderColor color.RGBA
	FocusBorderSize  int32

	OnClick  ClickEventHandler
	OnFocus  FocusEventHandler
	OnChange ChangeEventHandler
}
