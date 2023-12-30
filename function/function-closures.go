package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	//this is example of returning function closure,
	//variable pos and neg will be become function referenced from function adder
	//but this also save the value from refrenced variable like sum
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		//this will return the adder function statement, but different call name function
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
