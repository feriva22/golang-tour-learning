package main

import "fmt"

func main() {
	fmt.Println("counting")

	//this will executed reversed, because defer only executed until the surrounding function returns
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
