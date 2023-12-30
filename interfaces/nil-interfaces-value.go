package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	//this is will return error, because interface called without struct implemented it
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
