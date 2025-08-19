package rlutils

type Circle struct {
	Center Vector2
	Radius float32
}

func Circle_Contains(c Circle, point Vector2) bool {
	return V2_Distance(c.Center, point) <= c.Radius
}
