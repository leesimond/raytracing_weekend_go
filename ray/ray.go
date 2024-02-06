package ray

import "raytracing_weekend_go/vector"

type Ray struct {
	origin    vector.Point3
	direction vector.Vec3
}

func (r *Ray) At(t float64) vector.Point3 {
	return r.origin.Add(r.direction.MultiplyScalar(t))
}
