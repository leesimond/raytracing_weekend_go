package hittable

import (
	"raytracing_weekend_go/ray"
	"raytracing_weekend_go/vector"
)

type Hit struct {
	P      vector.Point3
	Normal vector.Vec3
	T      float64
}

type Hittable interface {
	Hit(r ray.Ray, rayTmin float64, rayTmax float64, rec *Hit) bool
}
