package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	//i is index and v is the value at the index
	//its same as foreach(pow as $i => $v) in php
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
