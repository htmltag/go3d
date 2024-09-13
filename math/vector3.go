package math

import (
	"math"
)

type Vector3 struct {
    X, Y, Z float64
}

func NewVector3(x, y, z float64) *Vector3 {
    return &Vector3{X: x, Y: y, Z: z}
}

func (v *Vector3) Add(other *Vector3) *Vector3 {
    return NewVector3(v.X+other.X, v.Y+other.Y, v.Z+other.Z)
}

func (v *Vector3) Subtract(other *Vector3) *Vector3 {
    return NewVector3(v.X-other.X, v.Y-other.Y, v.Z-other.Z)
}

func (v *Vector3) Multiply(scalar float64) *Vector3 {
    return NewVector3(v.X*scalar, v.Y*scalar, v.Z*scalar)
}

func (v *Vector3) Dot(other *Vector3) float64 {
    return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v *Vector3) Cross(other *Vector3) *Vector3 {
    return NewVector3(
        v.Y*other.Z - v.Z*other.Y,
        v.Z*other.X - v.X*other.Z,
        v.X*other.Y - v.Y*other.X,
    )
}

func (v *Vector3) Magnitude() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector3) Normalize() *Vector3 {
    mag := v.Magnitude()
    return NewVector3(v.X/mag, v.Y/mag, v.Z/mag)
}
