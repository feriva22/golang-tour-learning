package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false                // boolean type
	MaxInt uint64     = 1<<64 - 1            //unsigned integer with length 64 bit max
	z      complex128 = cmplx.Sqrt(-5 + 12i) // complex128 type is like you get 2 of float64 literals
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
