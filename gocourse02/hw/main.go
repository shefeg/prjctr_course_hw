package main

import (
	"fmt"
)

type Zookeeper struct {
	Name string
}

type Animal struct {
	Name string
	Size int
	Food string
}

type Cage struct {
	Size int
}

type Eagle struct {
	Animal
}

type Tiger struct {
	Animal
}

type Cow struct {
	Animal
}

type Monkey struct {
	Animal
}

type Turtle struct {
	Animal
}

func (z Zookeeper) CatchEagle(eagle Eagle, cage Cage, food string) {
	if eagle.Animal.Food == food && eagle.Animal.Size == cage.Size{
		fmt.Println(eagle.Animal.Name, "cought in a cage successfully!")
	} else {
		fmt.Println(eagle.Animal.Name, "catching to a cage failed :(")
	}
}

func (z Zookeeper) CatchTiger(tiger Tiger, cage Cage, food string) {
	if tiger.Animal.Food == food && tiger.Animal.Size == cage.Size{
		fmt.Println(tiger.Animal.Name, "cought in a cage successfully!")
	} else {
		fmt.Println(tiger.Animal.Name, "catching to a cage failed :(")
	}
}

func (z Zookeeper) CatchCow(cow Cow, cage Cage, food string) {
	if cow.Animal.Food == food && cow.Animal.Size == cage.Size{
		fmt.Println(cow.Animal.Name, "cought in a cage successfully!")
	} else {
		fmt.Println(cow.Animal.Name, "catching to a cage failed :(")
	}
}

func (z Zookeeper) CatchMonkey(monkey Monkey, cage Cage, food string) {
	if monkey.Animal.Food == food && monkey.Animal.Size == cage.Size{
		fmt.Println(monkey.Animal.Name, "cought in a cage successfully!")
	} else {
		fmt.Println(monkey.Animal.Name, "catching to a cage failed :(")
	}
}

func (z Zookeeper) CatchTurtle(turtle Turtle, cage Cage, food string) {
	if turtle.Animal.Food == food && turtle.Animal.Size == cage.Size{
		fmt.Println(turtle.Animal.Name, "cought in a cage successfully!")
	} else {
		fmt.Println(turtle.Animal.Name, "catching to a cage failed :(")
	}
}


func main() {
	zookeeper := Zookeeper{"John"}

	cage_eagle := Cage{10}
	cage_tiger := Cage{5}
	cage_cow := Cage{7}
	cage_monkey := Cage{4}
	cage_turtle := Cage{2}

	big_eagle := Eagle{Animal{"Eagle", 10, "meat"}}
	white_tiger := Tiger{Animal{"Tiger", 5,"meat"}}
	grey_cow := Cow{Animal{"Cow", 5, "grass"}}
	little_monkey := Monkey{Animal{"Monkey", 4, "banana"}}
	sea_turtle := Turtle{Animal{"Turtle", 2, "fish"}}

	zookeeper.CatchEagle(big_eagle,cage_eagle,"meat")
	zookeeper.CatchTiger(white_tiger,cage_tiger,"fish")
	zookeeper.CatchTiger(white_tiger,cage_tiger,"meat")
	zookeeper.CatchCow(grey_cow,cage_turtle,"fish")
	zookeeper.CatchCow(grey_cow,cage_cow,"grass")
	zookeeper.CatchMonkey(little_monkey,cage_monkey,"banana")
	zookeeper.CatchTurtle(sea_turtle,cage_turtle,"fish")
}