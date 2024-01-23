package vector

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

// func (v *Vec3) Add(other *Vec3) {
func (v *Vec3) Add(other Vec3) {
	// return Vec3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
}

func (v *Vec3) Multiply(other Vec3) {
	v.X *= other.X
	v.Y *= other.Y
	v.Z *= other.Z
}

// func (v Vec3) String() string {
// 	return fmt.Sprintf("Vec3: X: %v Y:%v Z:%v\n", v.X, v.Y, v.Z)
// }
