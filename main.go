package main

import (
	"fmt"
	"raytracing_weekend_go/vector"
)

func main() {
	fmt.Println("hello")
	a := vector.Vec3{X: 1, Y: 2, Z: 3}
	fmt.Println(a)
	a.Negate()
	fmt.Println(a)
	a.Negate()
	fmt.Println(a)
	// b := a.Add(&vector.Vec3{X: 4, Y: 5, Z: 6})
	a.Add(vector.Vec3{X: 4, Y: 5, Z: 6})
	fmt.Println(a)
	// fmt.Println(b)
	a.Multiply(vector.Vec3{X: 7, Y: 8, Z: 9})
	fmt.Println(a)
}
