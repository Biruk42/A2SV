package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(WordFrequency("This is Task 2"))
}

func WordFrequency(text string) map[string]int {
	text = strings.ToLower(text)
	re := regexp.MustCompile(`[a-z0-9]+`)
	words := re.FindAllString(text, -1)

	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}
	return frequency
}


