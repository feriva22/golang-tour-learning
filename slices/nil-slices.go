package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	//nil slices is zero length and zero capacity
	if s == nil {
		fmt.Println("nil!")
	}
}
