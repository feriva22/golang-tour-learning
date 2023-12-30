package main

import "fmt"

// create interface creature
// we define interaface for object that is creature
type Creature interface {
	Breathing() string
	Food() string
	IsMoving() bool
}

func main() {
	var c Creature

	human_1 := Human{human_race: "chinese", maximum_live: "70 Year"}
	fmt.Println(human_1)

	//call human_1 implement as Creature
	c = &human_1 // a human_1 implements Creature

	fmt.Println(c.Breathing()) //will return as Human

	//call indirectly the Creature when defining Plant object
	//but you not use Pointer for it
	var plant_1 Creature = Plant{color: "Yellow", maximum_live: "1 Year"}
	fmt.Println(plant_1)

	fmt.Println(plant_1.Breathing()) //will return as plant

}

// defining "class" and the interface Creature methods
type Human struct {
	human_race   string
	maximum_live string `default:"80 Year"`
}

func (h *Human) Breathing() string {
	return "Oxygen"
}

func (h *Human) Food() string {
	return "Creature"
}

func (h *Human) IsMoving() bool {
	return true
}

// defining "class" and the interface Creature methods
type Plant struct {
	color        string
	maximum_live string `default:"2 Year"`
}

func (p Plant) Breathing() string {
	return "NO"
}

func (p Plant) Food() string {
	return "Photosyntesis"
}

func (p Plant) IsMoving() bool {
	return false
}
