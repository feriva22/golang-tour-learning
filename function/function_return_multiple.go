package main

import (
	"fmt"
)

func swap(x string, y string) (string, string) { //add 2 param x and y then return int
	return y, x
}

func main() {
	//place the result in variable a and b
	a, b := swap("hello", "world")
	fmt.Printf("The swaped string of hello world is : %s and %s \n", a, b)

}
