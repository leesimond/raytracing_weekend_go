package main

import (
	"bufio"
	"fmt"
	"os"
)

// "raytracing_weekend_go/vector"

var imageWidth int = 256
var imageHeight int = 256

func main() {
	f := bufio.NewWriter(os.Stderr)
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)
	for j := 0; j < imageHeight; j++ {
		progress := fmt.Sprintf("\rScanlines remaining: %d ", imageHeight-j)
		f.WriteString(progress)
		f.Flush()
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / float64(imageWidth-1)
			g := float64(j) / float64(imageHeight-1)
			b := 0

			// var ir int = int(255.999 * r)
			// var ig int = int(255.999 * g)
			ir := int(255 * r)
			ig := int(255 * g)

			fmt.Printf("%d %d %d\n", ir, ig, b)
		}
	}
	f.Write([]byte("\rDone.                 \n"))
	f.Flush()
}
