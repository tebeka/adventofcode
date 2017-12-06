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

func permutations(word string) chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		if len(word) == 1 {
			ch <- word
			return
		}
		for i, c := range word {
			w := word[:i] + word[i+1:]
			for subw := range permutations(w) {
				ch <- string(c) + subw
			}
		}
	}()
	return ch
}

func inWords(word string, words []string) bool {
	for _, w := range words {
		if word == w {
			return true
		}
	}
	return false
}

func isValid2(passphrase string) bool {
	words := strings.Fields(passphrase)
	for i, word := range words {
		other := make([]string, len(words)-1)
		copy(other, words[:i])
		copy(other[i:], words[i+1:])

		for perm := range permutations(word) {
			if inWords(perm, other) {
				return false
			}
		}
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
	numValid2 := 0
	for scanner.Scan() {
		if isValid(scanner.Text()) {
			numValid++
		}
		if isValid2(scanner.Text()) {
			numValid2++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(numValid)
	fmt.Println(numValid2)
}
