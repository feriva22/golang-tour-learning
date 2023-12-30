package main

import "fmt"

// declare variable in package level
var c, python, java bool

func main() {
	//declare variable in function level
	var i int //default value will be used when not defined
	fmt.Println(i, c, python, java)
}
