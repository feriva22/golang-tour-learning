package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	//we can parsing the function as argument from another function
	fmt.Println(compute(hypot))

	//we can also parsing the function from another package as argument
	fmt.Println(compute(math.Pow))
}
