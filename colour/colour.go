package colour

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"raytracing_weekend_go/interval"
	"raytracing_weekend_go/vector"
)

type Colour = vector.Vec3

func linearToGamma(linearComponent float64) float64 {
	return math.Sqrt(linearComponent)
}

// Write the translated [0,255] value of each color component.
func WriteColour(file *os.File, pixelColour Colour, samplesPerPixel int) {
	w := bufio.NewWriter(file)
	r := pixelColour.X
	g := pixelColour.Y
	b := pixelColour.Z

	// Divide the colour by the number of samples
	scale := 1.0 / float64(samplesPerPixel)
	r *= scale
	g *= scale
	b *= scale

	// Apply the linear to gamma transform
	r = linearToGamma(r)
	g = linearToGamma(g)
	b = linearToGamma(b)

	// Write the translated [0,255] value of each colour component
	intensity := interval.Interval{Min: 0.0, Max: 0.999}
	colour := fmt.Sprintf(
		"%d %d %d\n",
		int(255*intensity.Clamp(r)),
		int(255*intensity.Clamp(g)),
		int(255*intensity.Clamp(b)),
	)
	w.WriteString(colour)
	w.Flush()
}
