package hittable

import (
	"raytracing_weekend_go/interval"
	"raytracing_weekend_go/ray"
	"raytracing_weekend_go/vector"
)

type HitRecord struct {
	P         vector.Point3
	Normal    vector.Vec3
	T         float64
	Material  Material
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
	Hit(r *ray.Ray, rayT *interval.Interval, rec *HitRecord) bool
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

func (h *HittableList) Hit(r *ray.Ray, rayT *interval.Interval, rec *HitRecord) bool {
	var tempRecord HitRecord
	hitAnything := false
	closestSoFar := rayT.Max

	for _, object := range h.objects {
		if object.Hit(r, &interval.Interval{Min: rayT.Min, Max: closestSoFar}, &tempRecord) {
			hitAnything = true
			closestSoFar = tempRecord.T
			*rec = tempRecord
		}
	}

	return hitAnything
}
