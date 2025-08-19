package entities

import (
	"chat/shared/rlutils"
)

type ComponentType int

const (
	ComponentTypeCircle ComponentType = iota
	ComponentTypeRectangle
	ComponentTypeInput
)

type Component interface {
	Type() ComponentType

	Contains(rlutils.Vector2) bool

	Children() []Component
	AddChild(Component)

	Text() string

	OnClick() OnClickHandler
	OnInput() OnInputHandler
}
