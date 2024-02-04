package colour

import (
	"bufio"
	"fmt"
	"os"
	"raytracing_weekend_go/vector"
)

type Colour = vector.Vec3

// Write the translated [0,255] value of each color component.
func WriteColour(file *os.File, pixelColour Colour) {
	w := bufio.NewWriter(file)
	colour := fmt.Sprintf(
		"%d %d %d\n",
		int(255*pixelColour.X),
		int(255*pixelColour.Y),
		int(255*pixelColour.Z),
	)
	w.WriteString(colour)
	w.Flush()
}
