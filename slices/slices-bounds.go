package main

import "fmt"

func main() {
	//definindg a slices
	s := []int{2, 3, 5, 7, 11, 13}

	//this type bound to get the slice will get : {3,5,7}
	s = s[1:4]
	fmt.Println(s)

	//this type bound, auto populate low bound with 0 , so will get : {3,5}
	s = s[:2]
	fmt.Println(s)

	//this type bond, auto populate high bound with length of slices (2 items), so will get : {5}
	s = s[1:]
	fmt.Println(s)
}
