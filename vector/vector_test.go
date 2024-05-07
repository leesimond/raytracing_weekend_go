package vector_test

import (
	"raytracing_weekend_go/vector"
	"strconv"
	"testing"
)

func TestVec3Addition(t *testing.T) {
	data := []struct {
		vec1     vector.Vec3
		vec2     vector.Vec3
		expected vector.Vec3
	}{
		{vector.Vec3{1, 2, 3}, vector.Vec3{4, 5, 6}, vector.Vec3{5, 7, 9}},
		{vector.Vec3{1, 2, 3}, vector.Vec3{-4, -5, -6}, vector.Vec3{-3, -3, -3}},
		{vector.Vec3{-1, -2, -3}, vector.Vec3{-4, -5, -6}, vector.Vec3{-5, -7, -9}},
	}
	for i, d := range data {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := d.vec1.Add(d.vec2)
			if result != d.expected {
				t.Errorf("Expected %v, got %v", d.expected, result)
			}
		})
	}
}
