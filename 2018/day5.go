package main

import (
	"fmt"
	"io/ioutil"
)

const (
	rDiff = 'a' - 'A'
)

func neg(r byte) byte {
	if r >= 'a' {
		return r - rDiff
	}

	return r + rDiff
}

func reduce(polymer []byte) []byte {
	dest := make([]byte, len(polymer))
	for i, c := range polymer[1:] {
		prev := polymer[i]
		if prev != neg(c) {
			dest[i] = prev
			continue
		}

		copy(dest[i:], polymer[i+2:])
		dest = dest[:len(polymer)-2]
		break
	}
	return dest
}

func reduceAll(polymer []byte) []byte {
	for {
		out := reduce(polymer)
		switch {
		case len(out) == len(polymer):
			return polymer
		case len(out) == 0:
			return out
		}
		polymer = out
	}
}

func part1(polymer []byte) {
	out := reduceAll(polymer)
	fmt.Println(len(out))
}

func remove(c byte, polymer []byte) []byte {
	out := make([]byte, len(polymer))
	i := 0
	for _, b := range polymer {
		if b == c || b == neg(c) {
			continue
		}
		out[i] = b
		i++
	}

	return out[:i]
}

func part2(polymer []byte) {
	// Using goroutines reduces runtime from 12sec to 5
	out := make(chan int)
	for c := byte('a'); c <= 'z'; c++ {
		go func(c byte) {
			r := remove(c, polymer)
			out <- len(reduceAll(r))
		}(c)
	}

	bestSize := -1
	for c := byte('a'); c <= 'z'; c++ {
		size := <-out
		if bestSize == -1 || size < bestSize {
			bestSize = size
		}
	}
	fmt.Println(bestSize)
}

func main() {
	polymer, err := ioutil.ReadFile("day-5.txt")
	if err != nil {
		panic(err)
	}
	polymer = polymer[:len(polymer)-1] // trim newline

	// polymer = []byte("dabAcCaCBAcCcaDA")
	part1(polymer)

	//polymer = []byte("dabAcCaCBAcCcaDA")
	part2(polymer)
}
