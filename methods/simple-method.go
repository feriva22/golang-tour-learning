package main

import (
	"fmt"
	"math"
	"time"
)

// in Golang , not have classes, but you can still define it as struct
type Vertex struct {
	X, Y float64
}

type Car struct {
	brand         string
	series        string
	engine_type   string
	tires_count   int
	release_year  int
	release_price float64
}

//when your "class" want have methode, you can create function as method with a special receiver argument
//The receiver appears in its own argument list between the func keyword and the method name.
//In this example, the Abs method has a receiver of type Vertex named v.

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (c Car) GapYearRelease() int {
	//the formula is every year between current year and release year is make the price drop 2%
	current_year := time.Now().Year()
	gap_year := current_year - c.release_year

	return gap_year
}

func (c Car) PredictCurrentPrice() float64 {

	gap_year := c.GapYearRelease()

	//price every year gap
	price_drop_every_year := 0.02 * c.release_price

	return c.release_price - (price_drop_every_year * float64(gap_year))

}

func main() {
	v := Vertex{3, 4}
	//call method directly from struct
	fmt.Println(v.Abs())

	//another example of creating "objects"
	car_1 := Car{brand: "Toyota", series: "Avanza", engine_type: "bensin", tires_count: 4, release_year: 2014, release_price: 150000000.0}
	car_2 := Car{brand: "Toyota", series: "Kijang Inova", engine_type: "diesel", tires_count: 4, release_year: 2005, release_price: 100000000.0}

	fmt.Printf("Gap Year of Car 1 : %d Year \n", car_1.GapYearRelease())
	fmt.Printf("Release price of Car 1 : %f \n", car_1.release_price)
	fmt.Printf("The current price of Car 1 : %f \n", car_1.PredictCurrentPrice())
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("Gap Year of Car 2 : %d Year \n", car_2.GapYearRelease())
	fmt.Printf("Release price of Car 2 : %f \n", car_2.release_price)
	fmt.Printf("The current price of Car 2 : %f \n", car_2.PredictCurrentPrice())
}
