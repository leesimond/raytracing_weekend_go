package vector

import (
	"fmt"
	"math"
	"math/rand"
	"raytracing_weekend_go/utils"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v *Vec3) Negate() Vec3 {
	return Vec3{X: -v.X, Y: -v.Y, Z: -v.Z}
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

func (v *Vec3) NearZero() bool {
	s := 1e-8
	return math.Abs(v.X) < s && math.Abs(v.Y) < s && math.Abs(v.Z) < s
}

func Random() Vec3 {
	return Vec3{X: rand.Float64(), Y: rand.Float64(), Z: rand.Float64()}
}

func RandomMinMax(min float64, max float64) Vec3 {
	return Vec3{X: utils.Random(min, max), Y: utils.Random(min, max), Z: utils.Random(min, max)}
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

func UnitVector(v Vec3) Vec3 {
	return v.DivideScalar(v.Length())
}

func RandomInUnitSphere() Vec3 {
	for {
		p := RandomMinMax(-1, 1)
		if p.LengthSquared() < 1 {
			return p
		}
	}
}

func RandomUnitVector() Vec3 {
	return UnitVector(RandomInUnitSphere())
}

func RandomOnHemisphere(normal *Vec3) Vec3 {
	onUnitSphere := RandomUnitVector()
	if onUnitSphere.Dot(*normal) > 0.0 {
		return onUnitSphere
	} else {
		return onUnitSphere.Negate()
	}
}

func Reflect(v Vec3, n Vec3) Vec3 {
	a := v.Dot(n)
	b := n.MultiplyScalar(a)
	b = b.MultiplyScalar(2)
	return v.Subtract(b)
}

func Refract(uv *Vec3, n *Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := math.Min(n.Dot(uv.Negate()), 1.0)
	rOutPerpendicular := uv.Add(n.MultiplyScalar(cosTheta))
	rOutPerpendicular.MultiplyScalarAssign(etaiOverEtat)
	rOutParallelSqrt := -math.Sqrt(math.Abs(1.0 - rOutPerpendicular.LengthSquared()))
	rOutParallel := n.MultiplyScalar(rOutParallelSqrt)
	return rOutPerpendicular.Add(rOutParallel)
}
