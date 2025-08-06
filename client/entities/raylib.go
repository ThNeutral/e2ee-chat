package entities

type WindowConfig struct {
	Width  int32
	Height int32
	Title  string
}

type WidgetType int

const (
	ButtonWidgetType WidgetType = iota
)
