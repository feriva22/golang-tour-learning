package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} //define array with fixed length

	//slice defined with []T , is a slice with type T
	var s []int = primes[1:4] //slice the array to get value start with index 1 to 4 position element, {3,5,7}
	fmt.Println(s)
}
