package main

import "fmt"

// struct is same as a collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4 //access the field using dot , and change the field value
	fmt.Println(v.X)
}
