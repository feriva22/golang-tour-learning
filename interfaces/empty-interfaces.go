package main

import "fmt"

func main() {
	//this will be called , and no need to defined in struct
	var i interface{}
	describe(i)

	i = 42 // example use is directly use any type of Values
	describe(i)

	i = "hello"
	describe(i)
}

//but must be defined if the parameter is empty interface
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
