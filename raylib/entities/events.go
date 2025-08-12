package entities

type ClickEvent struct {
	ShouldPropagate bool
}
type ClickEventHandler func(event *ClickEvent)

type FocusEvent struct{}
type FocusEventHandler func(event *FocusEvent)

type InputEvent struct {
	Text []rune
}
type InputEventHandler func(event *InputEvent)
