package main

import (
	"bufio"
	"fmt"
	"os"
	"raytracing_weekend_go/colour"
)

var imageWidth int = 256
var imageHeight int = 256

func main() {
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
