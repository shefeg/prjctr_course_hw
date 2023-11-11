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
	Status string
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

func (a *Animal) SetStatus(status string) {
	a.Status = status
	fmt.Println(a.Name, "is", a.Status)
}

func (z Zookeeper) CatchEagle(eagle *Eagle, cage Cage, food string) {
	if eagle.Animal.Food == food && eagle.Animal.Size == cage.Size{
		fmt.Println(eagle.Animal.Name, "cought in a cage successfully!")
		eagle.SetStatus("cought")
	} else {
		fmt.Println(eagle.Animal.Name, "catching to a cage failed :(")
		eagle.SetStatus("free")
	}
}

func (z Zookeeper) CatchTiger(tiger *Tiger, cage Cage, food string) {
	if tiger.Animal.Food == food && tiger.Animal.Size == cage.Size{
		fmt.Println(tiger.Animal.Name, "cought in a cage successfully!")
		tiger.SetStatus("free")
	} else {
		fmt.Println(tiger.Animal.Name, "catching to a cage failed :(")
		tiger.SetStatus("free")
	}
}

func (z Zookeeper) CatchCow(cow *Cow, cage Cage, food string) {
	if cow.Animal.Food == food && cow.Animal.Size == cage.Size{
		fmt.Println(cow.Animal.Name, "cought in a cage successfully!")
		cow.SetStatus("free")
	} else {
		fmt.Println(cow.Animal.Name, "catching to a cage failed :(")
		cow.SetStatus("free")
	}
}

func (z Zookeeper) CatchMonkey(monkey *Monkey, cage Cage, food string) {
	if monkey.Animal.Food == food && monkey.Animal.Size == cage.Size{
		fmt.Println(monkey.Animal.Name, "cought in a cage successfully!")
		monkey.SetStatus("free")
	} else {
		fmt.Println(monkey.Animal.Name, "catching to a cage failed :(")
		monkey.SetStatus("free")
	}
}

func (z Zookeeper) CatchTurtle(turtle *Turtle, cage Cage, food string) {
	if turtle.Animal.Food == food && turtle.Animal.Size == cage.Size{
		fmt.Println(turtle.Animal.Name, "cought in a cage successfully!")
		turtle.SetStatus("free")
	} else {
		fmt.Println(turtle.Animal.Name, "catching to a cage failed :(")
		turtle.SetStatus("free")
	}
}


func main() {
	zookeeper := Zookeeper{"John"}

	cage_eagle := Cage{10}
	cage_tiger := Cage{5}
	cage_cow := Cage{7}
	cage_monkey := Cage{4}
	cage_turtle := Cage{2}

	big_eagle := &Eagle{Animal{"Eagle", 10, "meat", "free"}}
	white_tiger := &Tiger{Animal{"Tiger", 5,"meat", "free"}}
	grey_cow := &Cow{Animal{"Cow", 5, "grass", "free"}}
	little_monkey := &Monkey{Animal{"Monkey", 4, "banana", "free"}}
	sea_turtle := &Turtle{Animal{"Turtle", 2, "fish", "free"}}

	zookeeper.CatchEagle(big_eagle,cage_eagle,"meat")
	zookeeper.CatchTiger(white_tiger,cage_tiger,"fish")
	zookeeper.CatchTiger(white_tiger,cage_tiger,"meat")
	zookeeper.CatchCow(grey_cow,cage_turtle,"fish")
	zookeeper.CatchCow(grey_cow,cage_cow,"grass")
	zookeeper.CatchMonkey(little_monkey,cage_monkey,"banana")
	zookeeper.CatchTurtle(sea_turtle,cage_turtle,"fish")
}