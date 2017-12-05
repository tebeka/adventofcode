package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isValid(passphrase string) bool {
	seen := make(map[string]bool)

	for _, word := range strings.Fields(passphrase) {
		if seen[word] {
			return false
		}
		seen[word] = true
	}

	return true
}

func main() {
	file, err := os.Open("input-4.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	numValid := 0
	for scanner.Scan() {
		if isValid(scanner.Text()) {
			numValid++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(numValid)
}
