package ray

import "raytracing_weekend_go/vector"

type Ray struct {
	origin    vector.Point3
	Direction vector.Vec3
}

func New(origin vector.Point3, direction vector.Vec3) Ray {
	return Ray{origin: origin, Direction: direction}
}

func (r *Ray) At(t float64) vector.Point3 {
	return r.origin.Add(r.Direction.MultiplyScalar(t))
}
