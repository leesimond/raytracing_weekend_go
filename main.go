package main

import (
	"raytracing_weekend_go/camera"
	"raytracing_weekend_go/colour"
	"raytracing_weekend_go/hittable"
	"raytracing_weekend_go/shape"
	"raytracing_weekend_go/vector"
)

func main() {
	var world hittable.HittableList

	materialGround := &hittable.Lambertian{Albedo: colour.Colour{X: 0.8, Y: 0.8}}
	materialCentre := &hittable.Lambertian{Albedo: colour.Colour{X: 0.1, Y: 0.2, Z: 0.5}}
	materialLeft := &hittable.Dielectric{IndexRefraction: 1.5}
	materialRight := hittable.NewMetal(colour.Colour{X: 0.8, Y: 0.6, Z: 0.2}, 0.0)

	world.Add(&shape.Sphere{Centre: vector.Point3{Y: -100.5, Z: -1.0}, Radius: 100, Material: materialGround})
	world.Add(&shape.Sphere{Centre: vector.Point3{Z: -1.0}, Radius: 0.5, Material: materialCentre})
	world.Add(&shape.Sphere{Centre: vector.Point3{X: -1.0, Z: -1.0}, Radius: 0.5, Material: materialLeft})
	world.Add(&shape.Sphere{Centre: vector.Point3{X: -1.0, Z: -1.0}, Radius: -0.4, Material: materialLeft})
	world.Add(&shape.Sphere{Centre: vector.Point3{X: 1.0, Z: -1.0}, Radius: 0.5, Material: materialRight})

	cam := camera.New()
	cam.AspectRatio = 16.0 / 9.0
	cam.ImageWidth = 400
	cam.SamplesPerPixel = 100
	cam.MaxDepth = 50

	cam.VFov = 20
	cam.LookFrom = vector.Point3{X: -2, Y: 2, Z: 1}
	cam.LookAt = vector.Point3{Z: -1}
	cam.VUp = vector.Vec3{Y: 1}

	cam.Render(&world)
}
