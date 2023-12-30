package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] //this will get {John , Paul}
	b := names[1:3] //this will get {Paul, George}
	fmt.Println(a, b)

	//lets get the example slice is pointer of array
	b[0] = "XXX" // this will also change element in 'names' variable , from Paul to XXX
	fmt.Println(a, b)
	fmt.Println(names)
}
