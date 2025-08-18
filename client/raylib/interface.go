package raylib

type ComponentType int

const (
	ComponentTypeCircle ComponentType = iota
	ComponentTypeRectangle
)

type Component interface {
	Type() ComponentType
}
