package main

import (
	"math/rand"
	"raytracing_weekend_go/camera"
	"raytracing_weekend_go/colour"
	"raytracing_weekend_go/hittable"
	"raytracing_weekend_go/shape"
	"raytracing_weekend_go/utils"
	"raytracing_weekend_go/vector"
)

func main() {
	var world hittable.HittableList

	groundMaterial := &hittable.Lambertian{Albedo: colour.Colour{X: 0.5, Y: 0.5, Z: 0.5}}
	world.Add(&shape.Sphere{Centre: vector.Point3{Y: -1000}, Radius: 1000, Material: groundMaterial})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMaterial := rand.Float64()
			centre := vector.Point3{X: float64(a) + 0.9*rand.Float64(), Y: 0.2, Z: float64(b) + 0.9*rand.Float64()}
			centrePointDifference := centre.Subtract(vector.Point3{X: 4, Y: 0.2})
			if centrePointDifference.Length() > 0.9 {
				var sphereMaterial hittable.Material
				if chooseMaterial < 0.8 {
					// diffuse
					var albedo colour.Colour = vector.Random()
					albedo = albedo.Multiply(vector.Random())
					sphereMaterial = &hittable.Lambertian{Albedo: albedo}
					world.Add(&shape.Sphere{Centre: centre, Radius: 0.2, Material: sphereMaterial})
				} else if chooseMaterial < 0.95 {
					// metal
					var albedo colour.Colour = vector.RandomMinMax(0.5, 1)
					fuzz := utils.Random(0, 0.5)
					sphereMaterial = &hittable.Metal{Albedo: albedo, Fuzz: fuzz}
					world.Add(&shape.Sphere{Centre: centre, Radius: 0.2, Material: sphereMaterial})
				} else {
					// glass
					sphereMaterial = &hittable.Dielectric{IndexRefraction: 1.5}
					world.Add(&shape.Sphere{Centre: centre, Radius: 0.2, Material: sphereMaterial})
				}
			}
		}
	}

	material1 := &hittable.Dielectric{IndexRefraction: 1.5}
	world.Add(&shape.Sphere{Centre: vector.Point3{Y: 1}, Radius: 1.0, Material: material1})

	material2 := &hittable.Lambertian{Albedo: colour.Colour{X: 0.4, Y: 0.2, Z: 0.1}}
	world.Add(&shape.Sphere{Centre: vector.Point3{X: -4, Y: 1}, Radius: 1.0, Material: material2})

	material3 := &hittable.Metal{Albedo: colour.Colour{X: 0.7, Y: 0.6, Z: 0.5}, Fuzz: 0.0}
	world.Add(&shape.Sphere{Centre: vector.Point3{X: 4, Y: 1}, Radius: 1.0, Material: material3})

	cam := camera.New()
	cam.AspectRatio = 16.0 / 9.0
	cam.ImageWidth = 1200
	cam.SamplesPerPixel = 10
	cam.MaxDepth = 50

	cam.VFov = 20
	cam.LookFrom = vector.Point3{X: 13, Y: 2, Z: 3}
	cam.LookAt = vector.Point3{}
	cam.VUp = vector.Vec3{Y: 1}

	cam.DefocusAngle = 0.6
	cam.FocusDistance = 10.0

	cam.Render(&world)
}
