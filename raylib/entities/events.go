package entities

type Event struct {
	ShouldPropagate bool
}

type ClickEvent struct {
	Event
}

type ClickEventHandler func(event *ClickEvent)
type FocusEventHandler func(this *RectangleWidget, focused bool)
type ChangeEventHandler func(this *RectangleWidget, text []rune)
