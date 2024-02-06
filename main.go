package main

import (
	"bufio"
	"fmt"
	"os"
	"raytracing_weekend_go/colour"
	"raytracing_weekend_go/ray"
	"raytracing_weekend_go/vector"
)

func rayColour(r *ray.Ray) colour.Colour {
	return colour.Colour{}
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

	w := bufio.NewWriter(os.Stderr)
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)
	for j := 0; j < imageHeight; j++ {
		progress := fmt.Sprintf("\rScanlines remaining: %d ", imageHeight-j)
		w.WriteString(progress)
		w.Flush()
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / float64(imageWidth-1)
			g := float64(j) / float64(imageHeight-1)
			b := 0.0
			pixelColour := colour.Colour{X: r, Y: g, Z: b}
			colour.WriteColour(os.Stdout, pixelColour)
		}
	}
	w.Write([]byte("\rDone.                 \n"))
	w.Flush()
}
