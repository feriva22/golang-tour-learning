package main

import "fmt"

func main() {
	m := make(map[string]int)

	//Directly append the new value by defining Key
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	//Directly change the value by defining Key
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	//Delete element in map by using Key
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	//Retrieve value with checking are value exist with the key defined
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
