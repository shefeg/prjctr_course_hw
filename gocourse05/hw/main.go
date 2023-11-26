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

func printField(init bool, field map[Key]string, size int) {
	if init {
		fmt.Println("Initing field")
	}
	for i := 1; i <= size; i++ {
		s := strconv.Itoa(i)
		if !init {
			fmt.Print(s)
		}
		for j := 1; j <= size; j++ {
			if init {
				field[Key{x: i, y: j}] = "-"
				s += field[Key{x: i, y: j}]
			} else {
				fmt.Print(field[Key{x: i, y: j}])
			}
		}
		if init {
			fmt.Println(s)
		} else {
			fmt.Println()
		}
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
	var xerr error
	var yerr error
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
			x, xerr = strconv.Atoi(result[0])
			y, yerr = strconv.Atoi(result[1])
		}
		if xerr != nil || yerr != nil || move != move || x < 1 || y < 1 || x > fieldSize || y > fieldSize {
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

func checkResult(move string, field map[Key]string, size int) bool {
	fmt.Println()
	for i := 1; i <= size; i++ {
		h_line := 0
		for j := 1; j <= size; j++ {
			if field[Key{x: i, y: j}] == move {
				h_line++
			}
			if h_line == size {
				return true
			}
		}
	}
	for i := 1; i <= size; i++ {
		v_line := 0
		for j := 1; j <= size; j++ {
			if field[Key{x: j, y: i}] == move {
				v_line++
			}
			if v_line == size {
				return true
			}
		}
	}
	var bd_line int
	for i := 1; i <= size; i++ {
		if field[Key{x: i, y: i}] == move  {
			bd_line++
		}
		if bd_line == size {
			return true
		}
	}
	var fd_line int
	for i := 1; i <= size; i++ {
		if field[Key{x: i, y: size + 1 - i}] == move {
			fd_line++
		}
		if fd_line == size {
			return true
		}
	}
	return false
}

const (
	fieldSize = 3
)

func main() {
	field := make(map[Key]string)
	printField(true, field, fieldSize)
	fmt.Println()
	for i := 1; i <= fieldSize * fieldSize; i++ {
		if i%2 != 0 {
			player := 1
			fmt.Println("Player", player)
			for !updateField("x", moveValidation, field, fieldSize) {
				fmt.Println("Make correct move!")
			}
			printField(false, field, fieldSize)
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
			printField(false, field, fieldSize)
			if checkResult("o", field, fieldSize) {
				fmt.Printf("Player %d won!\n", player)
				os.Exit(0)
			}
		}
	}
	fmt.Println("It's a draw!")
}
