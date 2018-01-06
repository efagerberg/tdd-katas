package physics

import "math"

type Vector3 struct {
	X, Y, Z float64
}

func (s Vector3) Add(other Vector3) Vector3 {
	return Vector3{s.X + other.X, s.Y + other.Y, s.Z + other.Z}
}

func (s Vector3) Sub(other Vector3) Vector3 {
	return Vector3{s.X - other.X, s.Y - other.Y, s.Z - other.Z}
}

func (s Vector3) Mul(factor float64) Vector3 {
	return Vector3{s.X * factor, s.Y * factor, s.Z * factor}
}

func (s Vector3) Div(factor float64) Vector3 {
	return Vector3{s.X / factor, s.Y / factor, s.Z / factor}
}

func (s Vector3) Magnitude() float64 {
	dist := math.Pow(s.X, 2) + math.Pow(s.Y, 2) + math.Pow(s.Z, 2)
	return math.Sqrt(dist)
}

func (s Vector3) Normalize() Vector3 {
	return s.Div(s.Magnitude())
}

func (s Vector3) Dot(other Vector3) float64 {
	return s.X*other.X + s.Y*other.Y + s.Z*other.Z
}
