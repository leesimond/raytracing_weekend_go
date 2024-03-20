package hittable

import (
	"raytracing_weekend_go/colour"
	"raytracing_weekend_go/ray"
	"raytracing_weekend_go/vector"
)

type Material interface {
	Scatter(rIn *ray.Ray, rec *HitRecord) (bool, colour.Colour, ray.Ray)
}

type Lambertian struct {
	Albedo colour.Colour
}

type Metal struct {
	Albedo colour.Colour
}

func (l *Lambertian) Scatter(rIn *ray.Ray, rec *HitRecord) (bool, colour.Colour, ray.Ray) {
	scatterDirection := rec.Normal.Add(vector.RandomUnitVector())

	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection = rec.Normal
	}
	scattered := ray.Ray{Origin: rec.P, Direction: scatterDirection}
	attenuation := l.Albedo
	return true, attenuation, scattered
}
func (m *Metal) Scatter(rIn *ray.Ray, rec *HitRecord) (bool, colour.Colour, ray.Ray) {
	reflected := vector.Reflect(vector.UnitVector(rIn.Direction), rec.Normal)
	scattered := ray.Ray{Origin: rec.P, Direction: reflected}
	attenuation := m.Albedo
	return true, attenuation, scattered
}
