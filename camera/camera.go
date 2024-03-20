package camera

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"raytracing_weekend_go/colour"
	"raytracing_weekend_go/hittable"
	"raytracing_weekend_go/interval"
	"raytracing_weekend_go/ray"
	"raytracing_weekend_go/vector"
)

type Camera struct {
	AspectRatio     float64
	ImageWidth      int
	imageHeight     int
	SamplesPerPixel int
	MaxDepth        int
	centre          vector.Point3
	pixel00Loc      vector.Point3
	pixelDeltaU     vector.Vec3
	pixelDeltaV     vector.Vec3
}

func New() Camera {
	return Camera{AspectRatio: 1, ImageWidth: 100, SamplesPerPixel: 10, MaxDepth: 10}
}

func (c *Camera) Render(world hittable.Hittable) {
	c.initialise()

	w := bufio.NewWriter(os.Stderr)
	fmt.Printf("P3\n%d %d\n255\n", c.ImageWidth, c.imageHeight)
	for j := 0; j < c.imageHeight; j++ {
		progress := fmt.Sprintf("\rScanlines remaining: %d ", c.imageHeight-j)
		w.WriteString(progress)
		w.Flush()
		for i := 0; i < c.ImageWidth; i++ {
			pixelColour := colour.Colour{}
			for sample := 0; sample < c.SamplesPerPixel; sample++ {
				r := c.getRay(i, j)
				pixelColour.AddAssign(rayColour(&r, c.MaxDepth, world))
			}
			colour.WriteColour(os.Stdout, pixelColour, c.SamplesPerPixel)
		}
	}
	w.Write([]byte("\rDone.                 \n"))
	w.Flush()
}

func (c *Camera) initialise() {
	c.imageHeight = int(float64(c.ImageWidth) / c.AspectRatio)
	if c.imageHeight < 1 {
		c.imageHeight = 1
	}

	c.centre = vector.Point3{}

	// Determine viewport dimensions
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(c.ImageWidth) / float64(c.imageHeight))

	// Calculate the vector across the horizontal and down the vertical viewport edges
	viewportU := vector.Vec3{X: viewportWidth}
	viewportV := vector.Vec3{Y: -viewportHeight}

	// Calculate the horizontal and vertical delta vector from pixel to pixel
	c.pixelDeltaU = viewportU.DivideScalar(float64(c.ImageWidth))
	c.pixelDeltaV = viewportV.DivideScalar(float64(c.imageHeight))

	// Calculate the location of the upper left pixel
	viewportUpperLeft := c.centre.Subtract(vector.Vec3{Z: focalLength})
	viewportUpperLeft = viewportUpperLeft.Subtract(viewportU.DivideScalar(2))
	viewportUpperLeft = viewportUpperLeft.Subtract(viewportV.DivideScalar(2))
	c.pixel00Loc = c.pixelDeltaU.Add(c.pixelDeltaV)
	c.pixel00Loc = c.pixel00Loc.MultiplyScalar(0.5)
	c.pixel00Loc = c.pixel00Loc.Add(viewportUpperLeft)
}

func (c *Camera) getRay(i int, j int) ray.Ray {
	pixelDeltaUI := c.pixelDeltaU.MultiplyScalar(float64(i))
	pixelDeltaVJ := c.pixelDeltaV.MultiplyScalar(float64(j))
	pixelCentre := c.pixel00Loc.Add(pixelDeltaUI.Add(pixelDeltaVJ))
	pixelSample := pixelCentre.Add(c.pixelSampleSquare())

	rayOrigin := c.centre
	rayDirection := pixelSample.Subtract(rayOrigin)
	return ray.Ray{Origin: rayOrigin, Direction: rayDirection}
}

func (c *Camera) pixelSampleSquare() vector.Vec3 {
	// Returns a random point in teh square surrounding a pixel at the origin
	px := -0.5 + rand.Float64()
	py := -0.5 + rand.Float64()
	pixelDelaURandom := c.pixelDeltaU.MultiplyScalar(px)
	pixelDelaVRandom := c.pixelDeltaV.MultiplyScalar(py)
	return pixelDelaURandom.Add(pixelDelaVRandom)
}

func rayColour(r *ray.Ray, depth int, world hittable.Hittable) colour.Colour {
	var rec hittable.HitRecord

	// If we've exceeded the ray bounce limit, no more light is gathered
	if depth <= 0 {
		return colour.Colour{}
	}

	if world.Hit(r, &interval.Interval{Min: 0.001, Max: math.Inf(1)}, &rec) {
		var scattered ray.Ray
		var attenuation colour.Colour
		isScattered, attenuation, scattered := rec.Material.Scatter(r, &rec)
		if isScattered {
			return attenuation.Multiply(rayColour(&scattered, depth-1, world))
		}
		return colour.Colour{}
	}

	unitDirection := vector.UnitVector(r.Direction)
	a := 0.5 * (unitDirection.Y + 1)
	colour1 := colour.Colour{X: 1, Y: 1, Z: 1}
	colour1 = colour1.MultiplyScalar(1 - a)
	colour2 := colour.Colour{X: 0.5, Y: 0.7, Z: 1}
	colour2 = colour2.MultiplyScalar(a)
	return colour1.Add(colour2)
}
