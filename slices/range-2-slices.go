package main

import "fmt"

func main() {
	//create empty slice with capacity 10
	pow := make([]int, 10)

	//you can skip to passing the value by using only one variable i , without second variable
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	//you can skip the index variable by passing _ in the position
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
