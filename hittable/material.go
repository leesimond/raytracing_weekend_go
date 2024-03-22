package hittable

import (
	"math"
	"math/rand"
	"raytracing_weekend_go/colour"
	"raytracing_weekend_go/ray"
	"raytracing_weekend_go/vector"
)

type Material interface {
	Scatter(rIn *ray.Ray, rec *HitRecord) (bool, *colour.Colour, *ray.Ray)
}

type Lambertian struct {
	Albedo colour.Colour
}

type Metal struct {
	Albedo colour.Colour
	Fuzz   float64
}

type Dielectric struct {
	IndexRefraction float64
}

func (l *Lambertian) Scatter(rIn *ray.Ray, rec *HitRecord) (bool, *colour.Colour, *ray.Ray) {
	scatterDirection := rec.Normal.Add(vector.RandomUnitVector())

	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection = rec.Normal
	}
	scattered := &ray.Ray{Origin: rec.P, Direction: scatterDirection}
	attenuation := &l.Albedo
	return true, attenuation, scattered
}

func NewMetal(albedo colour.Colour, fuzz float64) *Metal {
	if fuzz > 1 {
		fuzz = 1
	}
	return &Metal{Albedo: albedo, Fuzz: fuzz}
}

func (m *Metal) Scatter(rIn *ray.Ray, rec *HitRecord) (bool, *colour.Colour, *ray.Ray) {
	reflected := vector.Reflect(vector.UnitVector(rIn.Direction), rec.Normal)
	fuzz := vector.RandomUnitVector()
	fuzz = fuzz.MultiplyScalar(m.Fuzz)
	scattered := &ray.Ray{Origin: rec.P, Direction: reflected.Add(fuzz)}
	attenuation := &m.Albedo
	isScattered := scattered.Direction.Dot(rec.Normal) > 0
	return isScattered, attenuation, scattered
}

func (d *Dielectric) Scatter(rIn *ray.Ray, rec *HitRecord) (bool, *colour.Colour, *ray.Ray) {
	attenuation := &colour.Colour{X: 1.0, Y: 1.0, Z: 1.0}
	refractionRatio := d.IndexRefraction
	if rec.frontFace {
		refractionRatio = 1.0 / d.IndexRefraction
	}

	unitDirection := vector.UnitVector(rIn.Direction)
	cosTheta := math.Min(rec.Normal.Dot(unitDirection.Negate()), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1.0
	var direction vector.Vec3

	if cannotRefract || reflectance(cosTheta, refractionRatio) > rand.Float64() {
		direction = vector.Reflect(unitDirection, rec.Normal)
	} else {
		direction = vector.Refract(&unitDirection, &rec.Normal, refractionRatio)
	}

	scattered := &ray.Ray{Origin: rec.P, Direction: direction}
	return true, attenuation, scattered
}

func reflectance(cosine float64, refIndex float64) float64 {
	// Use Schlick's approximation for reflectance
	r0 := (1 - refIndex) / (1 + refIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
