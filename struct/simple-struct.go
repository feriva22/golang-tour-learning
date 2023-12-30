package main

import "fmt"

// struct is same as a collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
