package main

import "fmt"

func main() {
	var a [2]string //this will create fixed-length of array with 2 element only, can't be resized
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} //define length of array int and initialize the value
	fmt.Println(primes)
}
