package main

import (
	"math"
	"math/rand"
	"raytracing_weekend_go/camera"
	"raytracing_weekend_go/hittable"
	"raytracing_weekend_go/shape"
	"raytracing_weekend_go/vector"
)

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func random(min float64, max float64) float64 {
	// Returns a random real in [min, max)
	return min + (max-min)*rand.Float64()
}

func main() {
	var world hittable.HittableList

	world.Add(&shape.Sphere{Centre: vector.Point3{Z: -1}, Radius: 0.5})
	world.Add(&shape.Sphere{Centre: vector.Point3{Y: -100.5, Z: -1}, Radius: 100})

	cam := camera.New()
	cam.AspectRatio = 16.0 / 9.0
	cam.ImageWidth = 400
	cam.SamplesPerPixel = 100

	cam.Render(&world)
}
