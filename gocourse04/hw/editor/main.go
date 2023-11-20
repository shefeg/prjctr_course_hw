package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scanner(description string) []string {
	var text []string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(description)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		text = append(text, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return text
}

func findLine(saved_text []string, search_text []string) {
	var result_key []int
	var result_value []string
	for k, v := range saved_text {
		for _, m := range search_text {
			if strings.Contains(v, m) {
				result_key = append(result_key, k)
				result_value = append(result_value, m)
			}
		}
	}
	for k, v := range saved_text {
		fmt.Printf("\nSaved text:\n%d %s\n", k, v)
	}
	for k, v := range search_text {
		fmt.Printf("\nSearched text:\n%d %s\n", k, v)
	}
	fmt.Printf("\nResults:\n")
	for _, v := range result_key {
		fmt.Printf("line '%s' contains '%s'\n", saved_text[v], result_value[v])
	}
}

func main() {
	findLine(
		scanner("Enter your text to save to editor:\n*(you can finish input with empty line)\n"),
		scanner("Enter your search text:\n*(you can finish input with empty line)\n"),
	)

}
