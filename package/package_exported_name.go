package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(math.pi) // this is example of "unexported name" , this can't be accessed when exported when start  with un-capitalized letter
	fmt.Println(math.Pi) // this is right access Exported name (start with Capitalized Letter)
}
