package main

import (
    "fmt"
    "os"
    "reflect"
    "runtime"
    "strconv"
    "strings"
)

func GetFunctionName(i interface{}) string {
    fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
    _, name, _ := strings.Cut(fullName, ".")
    return name
}

type Human struct {
    Name         string
    NameMemory   bool
    GlobalMemory bool
    Bag          bool
}

type Cave struct {
    Entry string
    Dark  bool
}

type Backpack struct {
    Size       int
    Matches    *Matches
    Flashlight *Flashlight
    Knife      *Knife
}

type Matches struct {
    Item  string
    Exist bool
}

type Flashlight struct {
    Item  string
    Exist bool
}

type Knife struct {
    Item  string
    Exist bool
}

type Road struct {
    Destination string
}

type Body struct {
    IsDead bool
    Type   string
}

type Camp struct {
    HasPeople bool
}

type Bug struct {
    IsBig bool
}

func (b *Bug) Bite() {
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

func answerValidation() int {
    var s string
    var i int
    fmt.Println("Enter answer number: 1 or 2")
    for {
        fmt.Scan(&s)
        i, _ = strconv.Atoi(s)
        if i == 1 || i == 2 {
            break
        } else {
            fmt.Println("Enter a valid number: 1 or 2")
        }
    }
    return i
}

func rememberNameChoice(remember bool, men *Human) {
    men.NameMemory = remember
    if men.NameMemory {
        fmt.Printf("%s remembers his name.\n", men.Name)
    } else {
        fmt.Printf("%s doesn't remember his name.\n", men.Name)
    }
}

func findBagChoice(bagFind bool, men *Human, bag *Backpack) {
    men.Bag = bagFind
    if men.Bag {
        bag.Matches.Exist = true
        bag.Flashlight.Exist = true
        bag.Knife.Exist = true
        men.FindBag(*bag)
    } else {
        men.Bag = false
        men.FindBag(*bag)
    }
}

func caveChoice(caveDark bool, men *Human, cave *Cave) {
    cave.Dark = caveDark
    if cave.Dark {
        fmt.Printf("%s goes from cave to the forest.\n", men.Name)
    } else {
        fmt.Printf("%s stays in cage and waits for help.\n", men.Name)
        os.Exit(0)
    }
}

func roadChoice(location string, men *Human, road *Road) {
    road.Destination = location
    fmt.Printf("%s goes from cave to the %s.\n", men.Name, location)
}

func bodyChoice(isDead bool, body *Body) {
    if isDead {
        body.IsDead = true
        body.Type = "animal"
        fmt.Printf("Body belongs to dead %s.\n", body.Type)
    } else {
        body.IsDead = false
        body.Type = "human"
        fmt.Printf("Body belongs to alive %s.\n", body.Type)
    }
}

func comeTobodyChoice(come bool, men *Human, body *Body) {
    if come {
        if !body.IsDead {
            fmt.Printf("%s comes closer to the body. Alive %s kills %s.\n", men.Name, body.Type, men.Name)
            os.Exit(0)
        } else {
            fmt.Printf("%s comes closer to the body. %s body smells bad.\n", men.Name, body.Type)
        }
    } else {
        fmt.Printf("%s just leaves %s body and goes along.\n", men.Name, body.Type)
    }
}

func campPeopleChoice(camp *Camp, men *Human) {
    if camp.HasPeople {
        fmt.Printf("%s finds that camp is full of people.\n", men.Name)
    } else {
        fmt.Printf("%s finds that camp is empty.\n", men.Name)
    }
}

func campDecisionChoice(rest bool, men *Human) {
    if rest {
        fmt.Printf("%s decides to rest.\n", men.Name)
        fmt.Printf("In the nearest tent %s finds safe with two numbers and tries to open it.\n", men.Name)
    } else {
        fmt.Printf("%s decides to proceed journey.\n", men.Name)
        fmt.Printf("After some time %s finds his home.\n", men.Name)
        os.Exit(0)
    }
}

func openSafeChoice(open bool, men *Human, bug *Bug) {
    if open {
        fmt.Printf("%s opens safe.\n", men.Name)
        bug := Bug{true}
        bug.Bite()
    } else {
        fmt.Printf("%s can't find the code.\n", men.Name)
        os.Exit(0)
    }
}

func main() {
    men := &Human{
        Name:       "Steven",
        NameMemory: false,
        Bag:        false,
    }
    cave := &Cave{"big cage entrance", true}
    bag := &Backpack{
        Size:       10,
        Matches:    &Matches{"Matches", false},
        Flashlight: &Flashlight{"Flashlight", false},
        Knife:      &Knife{"Knife", false},
    }
    road := &Road{}
    body := &Body{}
    camp := &Camp{}
    bug := &Bug{}

    scenarioMapSlice := []map[string]interface{}{
        {
            "question": "Does %s remember his name?\n1: Yes\n2: No\n",
            "choice":   rememberNameChoice,
        },
        {
            "question": "Does %s find a bag?\n1: Yes\n2: No\n",
            "choice":   findBagChoice,
        },
        {
            "question": "Is it dark in the cave for %s?\n1: Yes\n2: No\n",
            "choice":   caveChoice,
        },
        {
            "question": "%s uses the road that leads to?\n1: forest\n2: field\n",
            "choice":   roadChoice,
        },
        {
            "question": "%s sees body in the %s. What is this body?\n1: animal\n2: human\n",
            "choice":   bodyChoice,
        },
        {
            "question": "What should %s do to the body?\n1: Leave the body\n2: Come closer and look\n",
            "choice":   comeTobodyChoice,
        },
        {
            "question": "After some time %s comes to camp. How many people in the camp?\n1: Camp is full of people\n2: Camp is empty\n",
            "choice":   campPeopleChoice,
        },
        {
            "question": "What should %s do in the camp?\n1: Have a rest\n2: Proceed journey\n",
            "choice":   campDecisionChoice,
        },
        {
            "question": "Can %s find the code and open safe?\n1: Finds the code.\n2: Can't find code.\n",
            "choice":   openSafeChoice,
        },
    }

    fmt.Printf("%s wokes up near %s.\n", men.Name, cave.Entry)

    for _, v := range scenarioMapSlice {
        switch GetFunctionName(v["choice"]) {
        case "bodyChoice":
            fmt.Printf(fmt.Sprint(v["question"]), men.Name, road.Destination)
        default:
            fmt.Printf(fmt.Sprint(v["question"]), men.Name)
        }
        a := answerValidation()
        switch a {
        case 1:
            switch GetFunctionName(v["choice"]) {
            case "rememberNameChoice":
                v["choice"].(func(remember bool, men *Human))(true, men)
            case "findBagChoice":
                v["choice"].(func(bagFind bool, men *Human, bag *Backpack))(true, men, bag)
            case "caveChoice":
                v["choice"].(func(caveDark bool, men *Human, cave *Cave))(true, men, cave)
            case "roadChoice":
                v["choice"].(func(location string, men *Human, road *Road))("forest", men, road)
            case "bodyChoice":
                v["choice"].(func(isDead bool, body *Body))(true, body)
            case "comeTobodyChoice":
                v["choice"].(func(come bool, men *Human, body *Body))(false, men, body)
            case "campPeopleChoice":
                v["choice"].(func(camp *Camp, men *Human))(camp, men)
            case "campDecisionChoice":
                v["choice"].(func(rest bool, men *Human))(true, men)
            case "openSafeChoice":
                v["choice"].(func(open bool, men *Human, bug *Bug))(true, men, bug)
            }
        case 2:
            switch GetFunctionName(v["choice"]) {
            case "rememberNameChoice":
                v["choice"].(func(remember bool, men *Human))(false, men)
            case "findBagChoice":
                v["choice"].(func(bagFind bool, men *Human, bag *Backpack))(false, men, bag)
            case "caveChoice":
                v["choice"].(func(caveDark bool, men *Human, cave *Cave))(false, men, cave)
            case "roadChoice":
                v["choice"].(func(location string, men *Human, road *Road))("field", men, road)
            case "bodyChoice":
                v["choice"].(func(isDead bool, body *Body))(false, body)
            case "comeTobodyChoice":
                v["choice"].(func(come bool, men *Human, body *Body))(true, men, body)
            case "campPeopleChoice":
                v["choice"].(func(camp *Camp, men *Human))(camp, men)
            case "campDecisionChoice":
                v["choice"].(func(rest bool, men *Human))(false, men)
            case "openSafeChoice":
                v["choice"].(func(open bool, men *Human, bug *Bug))(false, men, bug)
            }
        default:
            panic("unsupported")
        }
        fmt.Println()
    }
}
