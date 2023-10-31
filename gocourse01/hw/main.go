package main

import (
	"fmt"
	"strings"
	"github.com/fatih/color"
	"strconv"
)

type (
	title string
	body string
	educationPeriod string
	calcYears func ()
	generateTitle func (text string)
	generateBody func (text string)
)

func (t title) formatTitle() {
    whilte := color.New(color.FgWhite)
    boldWhite := whilte.Add(color.Bold)
    boldWhite.Println(t)
}

func (ep educationPeriod) calcYears() int {
	stringSlice := strings.Split(string(ep), "-")
	start, _ := strconv.Atoi(stringSlice[0])
	end, _ := strconv.Atoi(stringSlice[1])
	return end - start
}

func printYears (years int) {
	fmt.Printf("%d years\n", years)
}

func generateChapter (title generateTitle, body generateBody)

const (
	school_period educationPeriod = "2008-2010"
	school_body = `
	Went to scool. It was very nice! Better than kindegrarden`
)


func main() {
	printYears(school_period.calcYears())
	// calcYears(2008,2010)
    whilte := color.New(color.FgWhite)
    boldWhite := whilte.Add(color.Bold)
    boldWhite.Println("This will print text in bold white.")
}
