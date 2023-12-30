package main

import (
	"fmt"
)

func add(x int, y int) int { //add 2 param x and y then return int
	return x + y
}

func main() {
	fmt.Printf("The sum is %d OK.\n", add(10, 10)) //call add() and format as digit %d

}
