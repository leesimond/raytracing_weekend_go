package shape

import (
	"math"
	"raytracing_weekend_go/hittable"
	"raytracing_weekend_go/ray"
	"raytracing_weekend_go/vector"
)

type Sphere struct {
	Centre vector.Point3
	Radius float64
}

func (s *Sphere) Hit(r *ray.Ray, rayTmin float64, rayTmax float64, rec *hittable.HitRecord) bool {
	oc := r.Origin.Subtract(s.Centre)
	a := r.Direction.LengthSquared()
	halfB := oc.Dot(r.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false
	}
	sqrtD := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (-halfB - sqrtD) / a
	if root <= rayTmin || rayTmax <= root {
		root = (-halfB + sqrtD) / a
		if root <= rayTmin || rayTmax <= root {
			return false
		}
	}

	rec.T = root
	rec.P = r.At(rec.T)
	outwardNormal := rec.P.Subtract(s.Centre)
	outwardNormal.DivideScalarAssign(s.Radius)
	rec.SetFaceNormal(r, outwardNormal)

	return true
}
