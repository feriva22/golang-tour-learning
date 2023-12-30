package main

import "fmt"

func main() {
	sum := 0
	// in golang, only 'for' you can used for looping
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
