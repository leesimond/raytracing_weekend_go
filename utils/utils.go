package utils

import (
	"math"
	"math/rand"
)

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Random(min float64, max float64) float64 {
	// Returns a random real in [min, max)
	return min + (max-min)*rand.Float64()
}
