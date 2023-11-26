package main

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

type (
	title           string
	body            string
	educationPeriod string
	calcYears       func()
	generateTitle   func(text title)
	generateBody    func(text body)
)

func formatTitle(text title) {
	white := color.New(color.FgWhite)
	boldWhite := white.Add(color.Bold)
	boldWhite.Println(text)
}

func formatBody(text body) {
	blue := color.New(color.FgBlue)
	blue.Println(text)
}

func (b body) Border() {
	green := color.New(color.FgGreen)
	for i := 0; i <= len([]rune(b)); i++ {
		green.Print("-")
	}
	fmt.Println()
}

func (ep educationPeriod) calcYears() int {
	stringSlice := strings.Split(string(ep), "-")
	start, _ := strconv.Atoi(stringSlice[0])
	end, _ := strconv.Atoi(stringSlice[1])
	return end - start
}

func printYears(years int) {
	red := color.New(color.FgHiRed)
	red.Printf("Education lasted: %d years\n", years)
}

func generateChapter(title_text title, body_text body, years int, title generateTitle, body generateBody) {
	body_text.Border()
	title(title_text)
	body(body_text)
	printYears(years)
	body_text.Border()
}

const (
	school_period educationPeriod = "1995-2005"
	uni_period    educationPeriod = "2005-2010"
	school_title  title           = "School time! 🤪"
	uni_title     title           = "University time! ✍"
	school_body   body            = `
Went to scool. It was very nice! Better than kindegrarden`
	uni_body body = `
Went to University. It was better than school!`
)

func main() {
	generateChapter(school_title, school_body, school_period.calcYears(), formatTitle, formatBody)
	generateChapter(uni_title, uni_body, uni_period.calcYears(), formatTitle, formatBody)
}
