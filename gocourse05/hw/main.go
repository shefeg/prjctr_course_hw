package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Key struct {
	x int
	y int
}

func initField(field map[Key]string, size int) {
	fmt.Println("Initing field")
	for i := 1; i <= size; i++ {
		s := strconv.Itoa(i)
		for j := 1; j <= size; j++ {
			field[Key{x: i, y: j}] = "-"
			s += field[Key{x: i, y: j}]
		}
		fmt.Println(s)
	}
	fmt.Print(" ")
	for i := 1; i <= size; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}

func moveValidation(move string, fieldSize int) (int, int) {
	var s []string
	var x int
	var y int
	var err error
	intro := "Enter move ('%s') and coordinates from 1 to %d splitting with space. Example: 1 1\n"
	scanner := bufio.NewScanner(os.Stdin)
LoopStart:
	for {
		fmt.Printf(intro, move, fieldSize)
		fmt.Println()
		scanner.Scan()
		line := scanner.Text()
		s = append(s, line)
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		result := strings.Split(s[0], " ")
		if len(result) < 2 {
			break LoopStart
		} else {
			x, err = strconv.Atoi(result[0])
			y, err = strconv.Atoi(result[1])
		}
		if err != nil || move != move || x < 1 || y < 1 || x > fieldSize || y > fieldSize {
			fmt.Printf(intro, fieldSize)
		} else {
			break
		}
	}
	fmt.Println()
	return x, y
}

func updateField(move string, answer func(string, int) (int, int), field map[Key]string, size int) bool {
	x, y := answer(move, size)
	currentKey := field[Key{x: x, y: y}]
	if currentKey == "-" {
		field[Key{x: x, y: y}] = move
	} else {
		fmt.Println("Key", field[Key{x: x, y: y}], "already exists!")
		return false
	}
	return true
}
func printResult(field map[Key]string, size int) {
	for i := 1; i <= size; i++ {
		s := strconv.Itoa(i)
		fmt.Print(s)
		for j := 1; j <= size; j++ {
			fmt.Print(field[Key{x: i, y: j}])
		}
		fmt.Println()
	}
	fmt.Print(" ")
	for i := 1; i <= size; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}
func checkResult(move string, field map[Key]string, size int) bool {
	fmt.Println()
	var line int
	for i := 1; i <= size; i++ {
		line = 0
		for j := 1; j <= size; j++ {
			if field[Key{x: i, y: j}] == move || field[Key{x: j, y: i}] == move {
				line++
			}
		}
		if line == size {
			return true
		}
	}
	// hacky solution need advice to make better
	if field[Key{x: 1, y: 1}] == move && field[Key{x: 3, y: 3}] == move && field[Key{x: 2, y: 2}] == move {
		return true
	} else if field[Key{x: 1, y: 3}] == move && field[Key{x: 3, y: 1}] == move && field[Key{x: 2, y: 2}] == move {
		return true
	}

	return false
}

const (
	fieldSize = 3
)

func main() {
	field := make(map[Key]string)
	initField(field, fieldSize)
	fmt.Println()
	for i := 1; i <= fieldSize*3; i++ {
		if i%2 != 0 {
			player := 1
			fmt.Println("Player", player)
			for !updateField("x", moveValidation, field, fieldSize) {
				fmt.Println("Make correct move!")
			}
			printResult(field, fieldSize)
			if checkResult("x", field, fieldSize) {
				fmt.Printf("Player %d won!\n", player)
				os.Exit(0)
			}
		} else {
			player := 2
			fmt.Println("Player", player)
			for !updateField("o", moveValidation, field, fieldSize) {
				fmt.Println("Make correct move!")
			}
			printResult(field, fieldSize)
			if checkResult("o", field, fieldSize) {
				fmt.Printf("Player %d won!\n", player)
				os.Exit(0)
			}
		}
	}
	fmt.Println("It's a draw!")
}
