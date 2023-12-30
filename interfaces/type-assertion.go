package main

import "fmt"

func main() {
	//this is defined if value must string
	var i interface{} = "hello"

	//assert are i is string value ?
	s := i.(string)
	fmt.Println(s)

	//assert are i is string value also check is true or false ?
	s, ok := i.(string)
	fmt.Println(s, ok)

	//assert are i is string value also check is true or false ?
	f, ok := i.(float64)
	fmt.Println(f, ok)

	//assert are i is float64 , but will return error , because access as float64 but interface is string value
	//f = i.(float64) // panic
	//fmt.Println(f)


	//you can also use switch for check type of interface value
	do(21)
	do("hello")
	do(true)
}


func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
