package hittable

import (
	"raytracing_weekend_go/colour"
	"raytracing_weekend_go/ray"
	"raytracing_weekend_go/vector"
)

type Material interface {
	Scatter(rIn *ray.Ray, rec *HitRecord, attenuation *colour.Colour, scattered *ray.Ray) bool
}

type Lambertian struct {
	Albedo colour.Colour
}

type Metal struct {
	Albedo colour.Colour
}

func (l *Lambertian) Scatter(rIn *ray.Ray, rec *HitRecord, attenuation *colour.Colour, scattered *ray.Ray) bool {
	scatterDirection := rec.Normal.Add(vector.RandomUnitVector())

	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection = rec.Normal
	}
	scattered = &ray.Ray{Origin: rec.P, Direction: scatterDirection}
	attenuation = &l.Albedo
	return true
}
func (m *Metal) Scatter(rIn *ray.Ray, rec *HitRecord, attenuation *colour.Colour, scattered *ray.Ray) bool {
	reflected := vector.Reflect(vector.UnitVector(rIn.Direction), rec.Normal)
	scattered = &ray.Ray{Origin: rec.P, Direction: reflected}
	attenuation = &m.Albedo
	return true
}
