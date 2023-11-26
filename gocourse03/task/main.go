package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func numberValidation() int {
	var s string
	var i int
	fmt.Println("Enter number: ")
	for {
		_, err := fmt.Scan(&s)
		i, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("Enter a valid number")
		} else {
			break
		}
	}
	return i
}

func main() {
	rand_number := rand.Intn(10)
	fmt.Printf("Random number: %d\n", rand_number)
	var win bool
LoopStart:
	for i, j := 1, 3; i <= 3; {
		fmt.Printf("You have %d tries out of %d\n", i, j)
		console_number := numberValidation()
		switch {
		case console_number < rand_number:
			fmt.Println("Your number is more.\n")
		case console_number == rand_number:
			fmt.Println("Bingo!\n")
			win = true
			break LoopStart
		case console_number > rand_number:
			fmt.Println("Your number is less.\n")
		}
		i++
	}

	if win {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost!")
	}

}
