package entities

import "chat/shared/rlutils"

type ComponentType int

const (
	ComponentTypeCircle ComponentType = iota
	ComponentTypeRectangle
)

type Component interface {
	Type() ComponentType
	OnClick() OnClickHandler
	Contains(rlutils.Vector2) bool
	Children() []Component
	AddChild(Component)
}

type OnClickHandler func()
