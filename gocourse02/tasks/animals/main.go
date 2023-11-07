	// exercise
	// 1. Make a roaring tiger, which extends a cat
	// implement: mewing and roaring
	// 2. Make a trash bin, which have an ability to accept garbage, can be filled or empty

package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Legs int
	Size string
}

type Tiger struct {
	Cat
}

func (c Cat) Meow() {
	fmt.Println("meow")
}

func (t Tiger) Roar() {
	fmt.Println("ROAR!!!")
}

func main() {
	kitty := Cat{
		Legs: 4,
		Size: "small",
	}
	tiger := Tiger{
		Cat{
			Legs: 4,
			Size: "big",
		},
	}
	fmt.Printf("Cat: %+v\n",kitty)
	kitty.Meow()
	fmt.Printf("Tiger: %+v\n",tiger)
	tiger.Roar()
}