package main

import "fmt"

// defining is one per variable , you can add the type in the last
var i, j int = 1, 2

func main() {
	//you also not define the type, but the type will be defined automatically from initialized value type
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
