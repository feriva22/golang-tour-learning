package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// this is function not method, because not have "receiver argument"
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	//struct as the argument of function Abs
	fmt.Println(Abs(v))
}
