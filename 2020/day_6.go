package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var data = `
abc

a
b
c

ab
ac

a
a
a
a

b
`

func sumAnyone(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	total := 0
	m := make(map[rune]int)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if len(line) == 0 {
			if len(m) > 0 {
				total += len(m)
				m = make(map[rune]int)
			}
		}
		for _, r := range line {
			m[r] = 1
		}
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return total + len(m), nil
}

func maxCount(m map[rune]int, size int) int {
	total := 0
	for _, count := range m {
		if count == size {
			total++
		}
	}
	return total
}

func sumEveryone(r io.Reader) (int, error) {
	s := bufio.NewScanner(r)
	total := 0
	size := 0
	m := make(map[rune]int)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		size += 1
		if len(line) == 0 {
			if len(m) > 0 {
				total += maxCount(m, size-1)
			}
			m = make(map[rune]int)
			size = 0
		}
		for _, r := range line {
			m[r]++
		}
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return total + maxCount(m, size), nil
}

func main() {
	var r io.Reader
	if false {
		r = strings.NewReader(data)
	} else {
		file, err := os.Open("day-6.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		r = file
	}
	//fmt.Println(sumAnyone(r))
	fmt.Println(sumEveryone(r))
}
