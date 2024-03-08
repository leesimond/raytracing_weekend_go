package ray

import "raytracing_weekend_go/vector"

type Ray struct {
	Origin    vector.Point3
	Direction vector.Vec3
}

// Revisit value return instead of pointer
func New(origin vector.Point3, direction vector.Vec3) Ray {
	return Ray{Origin: origin, Direction: direction}
}

func (r *Ray) At(t float64) vector.Point3 {
	return r.Origin.Add(r.Direction.MultiplyScalar(t))
}
