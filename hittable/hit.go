package hittable

import (
	"raytracing_weekend_go/ray"
	"raytracing_weekend_go/vector"
)

type HitRecord struct {
	P         vector.Point3
	Normal    vector.Vec3
	T         float64
	frontFace bool
}

func (h *HitRecord) SetFaceNormal(r *ray.Ray, outwardNormal vector.Vec3) {
	// Set the hit record normal vector.
	// NOTE: the parameter `outwardNormal` is assumed to have unit length.

	h.frontFace = r.Direction.Dot(outwardNormal) < 0

	if h.frontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Negate()
	}
}

type Hittable interface {
	Hit(r *ray.Ray, rayTmin float64, rayTmax float64, rec *HitRecord) bool
}

type HittableList struct {
	objects []Hittable
}

func (h *HittableList) Clear() {
	h.objects = nil
}

func (h *HittableList) Add(object Hittable) {
	h.objects = append(h.objects, object)
}

func (h *HittableList) Hit(r *ray.Ray, rayTmin float64, rayTmax float64, rec *HitRecord) bool {
	var tempRecord HitRecord
	hitAnything := false
	closestSoFar := rayTmax

	for _, object := range h.objects {
		if object.Hit(r, rayTmin, closestSoFar, &tempRecord) {
			hitAnything = true
			closestSoFar = tempRecord.T
			*rec = tempRecord
		}
	}

	return hitAnything
}
