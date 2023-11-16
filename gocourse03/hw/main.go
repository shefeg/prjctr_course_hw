package main

import (
	"fmt"
	"os"
)

type Human struct {
	Name string
	NameMemory bool
	GlobalMemory bool
	Bag bool
}

type Cave struct {
	Entry string
	Dark bool
}

type Backpack struct {
	Size int
	Matches *Matches
	Flashlight *Flashlight
	Knife *Knife
}

type Matches struct {
	Item string
	Exist bool
}

type Flashlight struct {
	Item string
	Exist bool
}

type Knife struct {
	Item string
	Exist bool
}

type Road struct {
	Destination string
}

type Body struct {
	IsDead bool
	Type string
}

type Camp struct {
	HasPeople bool
}

type Bug struct {
	IsBig bool
}

func (b Bug) Bite() {
	if b.IsBig {
		fmt.Printf("Big bug bites her victim and it passes out!")
	} else {
		fmt.Printf("Small bug bites her victim and it just hearts!")
	}
} 

func (h *Human) FindBag(backpack Backpack) {
	switch h.Bag {
	case true:
		fmt.Printf("%s finds bag.\n", h.Name)
		fmt.Printf("%s's bag contains:\n", h.Name)
		if backpack.Matches.Exist == true {
			fmt.Printf("- Matches\n") 
		} else if backpack.Matches.Exist == false {
				fmt.Printf("- No Matches\n") 
		}
		if backpack.Flashlight.Exist == true {
			fmt.Printf("- Flashlight\n") 
		} else if backpack.Flashlight.Exist == false {
			fmt.Printf("- No Flashlight\n") 
		}
		if backpack.Knife.Exist == true {
			fmt.Printf("- Knife\n") 
		} else if backpack.Knife.Exist == false {
			fmt.Printf("- No Knife\n") 
		}
	case false:
		fmt.Printf("%s bag is lost.\n", h.Name)
	}
}

func (h Human) Choice(choice string) {
	fmt.Printf(choice)
}

var a int

func main() {
	men := &Human{
		Name: "Steven", 
		NameMemory: false,
		Bag: false,
	}

	cave := Cave{"big cage entrance", true}

	bag := Backpack{
		Size: 10,
		Matches: &Matches{"Matches", false},
		Flashlight: &Flashlight{"Flashlight", false},
		Knife: &Knife{"Knife", false},
	}

	fmt.Printf("%s wokes up near %s.\n", men.Name, cave.Entry)

	fmt.Printf("Does %s remember his name? 1 or 2?\n", men.Name)
	fmt.Scan(&a)
	switch a {
	case 1:
		men.NameMemory = true
		fmt.Printf("%s remembers his name.\n", men.Name)
	case 2:
		men.NameMemory = false
		fmt.Printf("%s doesn't remember his name.\n", men.Name)
	}
	fmt.Println()

	fmt.Printf("Does %s find a bag? 1 or 2?\n", men.Name)
	fmt.Scan(&a)
	switch a {
	case 1:
		men.Bag = true
		bag.Matches.Exist = true
		bag.Flashlight.Exist = true
		bag.Knife.Exist = true
		men.FindBag(bag)
	case 2:
		men.Bag = false
		men.FindBag(bag)
	}
	fmt.Println()
	
	fmt.Printf("Is it dark in the cage? 1 or 2?\n")
	fmt.Scan(&a)
	switch a {
	case 1:
		cave.Dark = false
		fmt.Printf("%s stays in cage and waits for help.\n", men.Name)
		os.Exit(0)
	case 2:
		cave.Dark = true
		fmt.Printf("Goes from cage to the forest.\n")
	}
	fmt.Println()
	
	road := &Road{}
	fmt.Printf("%s uses the road that leads to? 1 or 2?\n", men.Name)
	fmt.Scan(&a)
	switch a {
	case 1:
		road.Destination = "forest"
		fmt.Printf("Road leads to the %s.\n", road.Destination)
	case 2:
		road.Destination = "field"
		fmt.Printf("Road leads to the %s.\n", road.Destination)
	}
	fmt.Println()
	
	body := &Body{}
	fmt.Printf("%s sees body in the %s. What is this body? 1 or 2?\n", men.Name, road.Destination)
	fmt.Scan(&a)
	switch a {
	case 1:
		body.IsDead = true
		body.Type = "animal"
		if body.IsDead {
			fmt.Printf("Body belongs to dead %s.\n", body.Type)
		}
	case 2:
		body.Type = "human"
		if body.IsDead == false {
			fmt.Printf("Body belongs to alive %s.\n", body.Type)
		}
	}
	fmt.Println()

	fmt.Printf("What should %s do to the body? Leave the body or come closer and look? 1 or 2?\n", men.Name)
	fmt.Scan(&a)
	switch a {
	case 1:
		fmt.Printf("%s just leaves %s body and goes along.\n", men.Name, body.Type)
	case 2:
		if body.IsDead == false {
			fmt.Printf("%s comes closer to the body. Alive %s kills %s.\n",  men.Name, body.Type, men.Name)
			os.Exit(0)
		} else {
			fmt.Printf("%s comes closer to the body. %s body smells bad.\n", men.Name, body.Type)
		}
	}
	fmt.Println()

	camp := &Camp{}
	if camp.HasPeople {
		fmt.Printf("After some time %s comes to camp full of people.\n", men.Name)
	} else {
		fmt.Printf("After some time %s comes to camp with no people.\n", men.Name)
	}
	fmt.Println()

	var proceed_journey bool
	fmt.Printf("What should %s do in the camp? Rest of proceed journey? 1 or 2?\n", men.Name)
	fmt.Scan(&a)
	switch a {
	case 1:
		if proceed_journey == false {
			fmt.Printf("%s decides to rest.\n", men.Name)
			fmt.Printf("In the nearest tent %s finds safe with two numbers and tries to open it.\n", men.Name)
		}
	case 2:
		proceed_journey = true
		if proceed_journey {
			fmt.Printf("%s decides to proceed journey.\n", men.Name)
			fmt.Printf("After some time %s finds his home.\n", men.Name)
			os.Exit(0)
		}
	}
	
	fmt.Println()
	fmt.Printf("Can %s find the code and open safe? 1 or 2?\n", men.Name)
	fmt.Scan(&a)
	switch a {
	case 1:
		fmt.Printf("%s opens safe.\n", men.Name)
		bug := Bug{true}
		bug.Bite()
	case 2:
		fmt.Printf("%s can't find the code.\n", men.Name)
		os.Exit(0)
	}

}