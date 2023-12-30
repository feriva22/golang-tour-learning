package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	v := Vertex{3, 4}
	//this is normal calling of the vertex struct
	v.Scale(2)
	//this is normal calling of the vertex struct pointer
	ScaleFunc(&v, 10)

	//this is indirection , p is the pointer of Vertect struct
	p := &Vertex{4, 3}
	p.Scale(3)      // this will automatically read p variable as pointer or not
	ScaleFunc(p, 8) //you not need to add '&' because p variable is pointer already

	fmt.Println(v, p)
}
