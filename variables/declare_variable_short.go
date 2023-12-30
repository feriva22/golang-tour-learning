package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3                                //this is short variable defined with implicit type (auto detect type by value init)
	c, python, java := true, false, "no!" // := is not available in package level

	fmt.Println(i, j, k, c, python, java)
}
