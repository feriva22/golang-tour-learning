package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	//convert x and y as float64
	var f float64 = math.Sqrt(float64(x*x + y*y))
	//convert f as unsigned integer
	var z uint = uint(f)

	fmt.Println(x, y, z)
}
