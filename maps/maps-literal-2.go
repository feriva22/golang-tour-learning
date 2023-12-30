package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{ //this will auto determine the type of value is Vertex struct
	"Bell Labs": {40.68433, -74.39967}, //you can directly parsing the value
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
