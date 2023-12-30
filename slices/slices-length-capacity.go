package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4] //in before access get zero items, but this will get item from underlying item if zero items
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// Extend the access of data higher than the capacity of underlying array, will triger error
	s = s[:8]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
