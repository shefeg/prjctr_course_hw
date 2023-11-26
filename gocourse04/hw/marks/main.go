package main

import (
	"fmt"
	"os"
)

func calcAvgMark(marks []float64) float64 {
	var sum float64
	if len(marks) == 0 {
		fmt.Println("No marks found. Insert marks into slice.")
		os.Exit(0)
	}
	for _, v := range marks {
		sum += v
	}
	return sum / float64(len(marks))
}

func main() {
	marks := []float64{3, 5, 2, 3.5, 4.3}
	fmt.Println("Marks: ", marks)
	fmt.Println("Avg mark:", calcAvgMark(marks))
}
