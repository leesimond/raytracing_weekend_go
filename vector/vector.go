package vector

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v *Vec3) Negate() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

func (v *Vec3) AddAssign(other Vec3) {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
}

func (v *Vec3) MultiplyScalarAssign(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
}

func (v *Vec3) DivideScalarAssign(scalar float64) {
	v.MultiplyScalarAssign(1 / scalar)
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Useful for geometric clarity in the code
type Point3 = Vec3

func (v Vec3) String() string {
	return fmt.Sprintf("X: %v Y:%v Z:%v\n", v.X, v.Y, v.Z)
}

func (v *Vec3) Add(other Vec3) Vec3 {
	return Vec3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

func (v *Vec3) Subtract(other Vec3) Vec3 {
	return Vec3{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

func (v *Vec3) Multiply(other Vec3) Vec3 {
	return Vec3{X: v.X * other.X, Y: v.Y * other.Y, Z: v.Z * other.Z}
}

func (v *Vec3) MultiplyScalar(scalar float64) Vec3 {
	return Vec3{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

func (v *Vec3) DivideScalar(scalar float64) Vec3 {
	return v.MultiplyScalar(1 / scalar)
}

func (v *Vec3) Dot(other Vec3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v *Vec3) Cross(other Vec3) Vec3 {
	return Vec3{X: v.Y*other.Z - v.Z*other.Y, Y: v.Z*other.X - v.X*other.Z, Z: v.X*other.Y - v.Y*other.X}
}

func (v *Vec3) UnitVector() Vec3 {
	return v.DivideScalar(v.Length())
}
