package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// z := 1.0
	z := x / 2
	zz := 0.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-zz) < 1e-10 {
			break
		}
		zz = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
