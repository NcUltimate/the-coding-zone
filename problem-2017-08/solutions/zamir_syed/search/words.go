package main

import (
	"bufio"
	"os"
	"strings"
)

var words map[string]bool

// Valid ...
func Valid(word string) bool {
	return words[word]
}

// LoadWords ...
func LoadWords(filename string) {

	words = make(map[string]bool)
	
	// Open File
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// Scan Lines
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		for _, w := range tokens {
			upper := strings.ToUpper(w)
			words[upper] = true
		}
	}
}
