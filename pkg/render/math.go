package render

import (
	"math"
)

type Vec2f struct {
	X float64
	Y float64
}

func (v1 Vec2f) Dist(v2 Vec2f) float64 {
	x, y := v1.X-v2.X, v1.Y-v2.Y
	return math.Sqrt(x*x + y*y)
}

func (v1 Vec2f) Add(v2 Vec2f) Vec2f {
	return Vec2f{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 Vec2f) Sub(v2 Vec2f) Vec2f {
	return Vec2f{v1.X - v2.X, v1.Y - v2.Y}
}

func (v1 Vec2f) Mul(scalar float64) Vec2f {
	return Vec2f{v1.X * scalar, v1.Y * scalar}
}

func (v1 Vec2f) Div(scalar float64) Vec2f {
	return Vec2f{v1.X / scalar, v1.Y / scalar}
}

func (v1 Vec2f) Lerp(v2 Vec2f, t float64) Vec2f {
	return v1.Add(v2.Sub(v1).Mul(t))
}

func (v1 Vec2f) Gamma(pow float64) Vec2f {
	return Vec2f{math.Pow(v1.X, pow), math.Pow(v1.Y, pow)}
}

func (v1 Vec2f) Fminf(v2 Vec2f) Vec2f {
	return Vec2f{min(v1.X, v2.X), min(v1.Y, v2.Y)}
}

func (v1 Vec2f) Fmaxf(v2 Vec2f) Vec2f {
	return Vec2f{max(v1.X, v2.X), max(v1.Y, v2.Y)}
}

type Vec4f struct {
	X float64
	Y float64
	Z float64
	A float64
}

func (v1 Vec4f) Dist(v2 Vec4f) float64 {
	x, y, z, a := v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z, v1.A-v2.A
	return math.Sqrt(x*x + y*y + z*z + a*a)
}

func (v1 Vec4f) Add(v2 Vec4f) Vec4f {
	return Vec4f{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z, v1.A + v2.A}
}

func (v1 Vec4f) Sub(v2 Vec4f) Vec4f {
	return Vec4f{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z, v1.A - v2.A}
}

func (v Vec4f) Mul(scalar float64) Vec4f {
	return Vec4f{v.X * scalar, v.Y * scalar, v.Z * scalar, v.A * scalar}
}

func (v Vec4f) Div(scalar float64) Vec4f {
	return Vec4f{v.X / scalar, v.Y / scalar, v.Z / scalar, v.A / scalar}
}

func (v1 Vec4f) Lerp(v2 Vec4f, t float64) Vec4f {
	return v1.Add(v2.Sub(v1).Mul(t))
}

func (v1 Vec4f) Gamma(pow float64) Vec4f {
	return Vec4f{math.Pow(v1.X, pow), math.Pow(v1.Y, pow), math.Pow(v1.Z, pow), math.Pow(v1.A, pow)}
}

func (v1 Vec4f) Fminf(v2 Vec4f) Vec4f {
	return Vec4f{min(v1.X, v2.X), min(v1.Y, v2.Y), min(v1.Z, v2.Z), min(v1.A, v2.A)}
}

func (v1 Vec4f) Fmaxf(v2 Vec4f) Vec4f {
	return Vec4f{max(v1.X, v2.X), max(v1.Y, v2.Y), max(v1.Z, v2.Z), max(v1.A, v2.A)}
}
