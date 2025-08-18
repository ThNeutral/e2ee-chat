package rlutils

import "math"

type Vector2 struct {
	X int32
	Y int32
}

func V2_DistanceSquared(left, right Vector2) int32 {
	dx := left.X - right.X
	dy := left.Y - right.Y
	return dx*dx + dy*dy
}

func V2_Distance(left, right Vector2) float32 {
	return float32(math.Sqrt(float64(V2_DistanceSquared(left, right))))
}
