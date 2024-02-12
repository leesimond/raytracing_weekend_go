package main

import (
	"bufio"
	"fmt"
	"os"
	"raytracing_weekend_go/colour"
	"raytracing_weekend_go/ray"
	"raytracing_weekend_go/vector"
)

func hitSphere(centre *vector.Point3, radius float64, r *ray.Ray) bool {
	oc := r.Origin.Subtract(*centre)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant >= 0
}

func rayColour(r *ray.Ray) colour.Colour {
	if hitSphere(&vector.Point3{Z: -1}, 0.5, r) {
		return colour.Colour{X: 1}
	}

	unitDirection := r.Direction.UnitVector()
	a := 0.5 * (unitDirection.Y + 1)
	colour1 := colour.Colour{X: 1, Y: 1, Z: 1}
	colour1 = colour1.MultiplyScalar(1 - a)
	colour2 := colour.Colour{X: 0.5, Y: 0.7, Z: 1}
	colour2 = colour2.MultiplyScalar(a)
	return colour1.Add(colour2)
}

func main() {
	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}

	// Camera
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCentre := vector.Point3{}

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := vector.Vec3{X: viewportWidth}
	viewportV := vector.Vec3{Y: -viewportHeight}

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixelDeltaU := viewportU.DivideScalar(float64(imageWidth))
	pixelDeltaV := viewportV.DivideScalar(float64(imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := cameraCentre.Subtract(vector.Vec3{Z: focalLength})
	viewportUpperLeft = viewportUpperLeft.Subtract(viewportU.DivideScalar(2))
	viewportUpperLeft = viewportUpperLeft.Subtract(viewportV.DivideScalar(2))
	pixelDelta := pixelDeltaU.Add(pixelDeltaV)
	pixel00Loc := viewportUpperLeft.Add(pixelDelta.MultiplyScalar(0.5))

	w := bufio.NewWriter(os.Stderr)
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)
	for j := 0; j < imageHeight; j++ {
		progress := fmt.Sprintf("\rScanlines remaining: %d ", imageHeight-j)
		w.WriteString(progress)
		w.Flush()
		for i := 0; i < imageWidth; i++ {
			pixelDeltaUI := pixelDeltaU.MultiplyScalar(float64(i))
			pixelDeltaVJ := pixelDeltaV.MultiplyScalar(float64(j))
			pixelCentre := pixel00Loc.Add(pixelDeltaUI.Add(pixelDeltaVJ))
			rayDirection := pixelCentre.Subtract(cameraCentre)
			// Revisit New()
			r := ray.New(cameraCentre, rayDirection)
			pixelColour := rayColour(&r)
			colour.WriteColour(os.Stdout, pixelColour)
		}
	}
	w.Write([]byte("\rDone.                 \n"))
	w.Flush()
}
